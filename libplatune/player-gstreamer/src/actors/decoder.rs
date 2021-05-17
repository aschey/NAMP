use crate::util::clocktime_to_seconds;
use servo_media_audio::decoder::DecoderExit;
use std::{collections::VecDeque, iter::FromIterator};

use std::{
    convert::TryInto,
    fs::File,
    io::Read,
    sync::{
        mpsc::{Receiver, SyncSender},
        Arc, Mutex,
    },
};

//use crate::context::CONTEXT;
use act_zero::*;
use bus::Bus;
use futures::future::join;
use gstreamer::{
    glib::filename_to_uri, prelude::ObjectExt, ClockTime, ElementExt, ElementExtManual,
    ElementFactory, State,
};
use log::{error, info, warn};
use postage::{mpsc, prelude::Stream, sink::Sink};
use servo_media_audio::{
    context::{AudioContext, RealTimeAudioContextOptions},
    decoder::AudioDecoderCallbacks,
};

use super::{gstreamer_context::GStreamerContext, player::Player, song_queue::SongQueue};

pub struct Decoder {
    context_addr: Addr<GStreamerContext>,
    player_addr: Addr<Player>,
    prev_duration: Option<f64>,
}

impl Actor for Decoder {}

// impl Drop for Decoder {
//     fn drop(&mut self) {
//         info!("dropping");
//     }
// }

impl Decoder {
    pub fn new(context_addr: Addr<GStreamerContext>, player_addr: Addr<Player>) -> Decoder {
        Decoder {
            context_addr,
            player_addr,
            prev_duration: None,
        }
    }
    pub async fn decode(
        &mut self,
        song: String,
        start_millis: Option<u64>,
        shutdown_sender: SyncSender<DecoderExit>,
        shutdown_receiver: Arc<Mutex<Receiver<DecoderExit>>>,
    ) -> ActorResult<DecoderExit> {
        // let decoded_audio: Arc<Mutex<Vec<Vec<f32>>>> = Arc::new(Mutex::new(Vec::new()));
        // let decoded_audio_ = decoded_audio.clone();
        // let decoded_audio__ = decoded_audio.clone();
        //let (mut sender, mut receiver) = mpsc::channel(32);
        info!("Start decode");

        let uri = {
            if let Ok(parsed_uri) = filename_to_uri(song.clone(), None) {
                parsed_uri.to_string()
            } else {
                song
            }
        };
        let (data_sender, mut data_receiver) = mpsc::channel::<(Box<[f32]>, u32)>(100000);
        let mut_sender = Mutex::new(data_sender);
        let (mut chan_sender, mut chan_receiver) = mpsc::channel(32);
        let uri_ = uri.clone();
        let callbacks = AudioDecoderCallbacks::new()
            .eos(move || {
                info!("EOS");
                // if let Err(err) = sender.try_send(()) {
                //     warn!("Error sending EOS: {}", err);
                // }
            })
            .error(|e| {
                error!("Error decoding audio {:?}", e);
            })
            .progress(move |buffer, mut channel| {
                if channel == 0 {
                    channel = 1;
                }
                let buf = (*buffer).as_ref().try_into().unwrap();
                if let Err(e) = mut_sender.lock().unwrap().try_send((buf, channel)) {
                    //error!("{}", e);
                }
            })
            .ready(move |channels| {
                info!("Decoding {}, channels: {}", uri_, channels);
                chan_sender.try_send(channels as usize).unwrap();
            })
            .build();

        let decode_bus = Arc::new(Mutex::new(Bus::new(16)));

        let handle = call!(self.context_addr.decode_audio_data(
            uri.clone(),
            start_millis,
            decode_bus.clone(),
            shutdown_sender,
            shutdown_receiver,
            callbacks
        ))
        .await
        .unwrap();

        let buffer_node_id = call!(self.context_addr.create_buffer()).await.unwrap();

        let decode_bus_ = decode_bus.clone();
        send!(self.context_addr.subscribe_need_data(
            buffer_node_id,
            Box::new(move || {
                decode_bus_.lock().unwrap().broadcast(());
            })
        ));

        info!("about to load");
        send!(self
            .player_addr
            .load(buffer_node_id, uri.clone(), self.prev_duration, 0.));

        // context
        //     .lock()
        //     .unwrap()
        //     .decode_audio_data(bytes.to_vec(), callbacks);

        // let (_, duration) = join(receiver.recv(), self.get_duration(&filename)).await;
        //info!("Finished decoding");

        let RealTimeAudioContextOptions {
            sample_rate,
            latency_hint: _,
        } = RealTimeAudioContextOptions::default();
        let sample_rate = sample_rate as f64;

        let chans = chan_receiver.recv().await.unwrap_or_default();

        let mut decoded_audio = vec![<VecDeque<f32>>::new(); chans];

        let mut start_silence = 0;
        let mut start_done = false;
        let mut end_silence = 0;
        //drop(context);
        //let context_mut_ = context_mut.clone();
        //thread::spawn(move || {
        let mut set = false;
        let mut got_duration = false;
        //let mut duration_millis = 0;
        let mut pushed = 0;
        let mut zeroes: i32 = 0;
        // if start_millis.is_some() {
        //     return Produces::ok(DecoderExit::Eos);
        // }

        while let Some((data, channel)) = data_receiver.recv().await {
            //println!();
            //println!("{}", start_millis);
            // if !got_duration {
            //     if let Ok(dur) = duration_receiver.try_recv() {
            //         duration_millis = dur;
            //         got_duration = true;
            //         println!("Duration {}", duration_millis);
            //     }
            // }
            if chans == 0 || data.len() == 0 {
                continue;
            }
            decoded_audio[(channel - 1) as usize]
                .append(&mut VecDeque::from_iter((*data).to_owned()));
            let mut able_to_push = decoded_audio.iter().map(|a| a.len()).min().unwrap();

            pushed += able_to_push;

            if able_to_push == 0 {
                continue;
            }
            for chan in &mut decoded_audio {
                chan.make_contiguous();
            }
            if !start_done {
                let new_silence = decoded_audio
                    .iter()
                    .map(|a| {
                        a.iter()
                            .try_fold(0, |acc, n| if *n == 0. { Ok(acc + 1) } else { Err(acc) })
                            .unwrap_or_else(|e| e)
                    })
                    .min()
                    .unwrap();

                start_silence += new_silence;
                decoded_audio.iter_mut().for_each(|a| {
                    a.drain(..new_silence);
                });

                if new_silence == able_to_push {
                    continue;
                }

                if new_silence < able_to_push {
                    able_to_push -= new_silence;

                    info!("start silence {}", start_silence);
                    start_done = true;
                    info!("loaded");
                }
            } else {
                let new_silence = decoded_audio
                    .iter()
                    .map(|a| {
                        a.as_slices().0[..able_to_push]
                            .iter()
                            .rev()
                            .try_fold(0, |acc, n| if *n == 0. { Ok(acc + 1) } else { Err(acc) })
                            .unwrap_or_else(|e| e)
                    })
                    .min()
                    .unwrap();
                if end_silence > 0 && new_silence < able_to_push {
                    end_silence = 0;
                } else {
                    end_silence += new_silence;
                }

                //println!("end silence {}", end_silence);
            }
            let with_zero = decoded_audio
                .iter_mut()
                .map(|a| a.drain(..able_to_push).collect::<VecDeque<_>>())
                .collect::<Vec<_>>();
            // if pushed == 2413568 {
            //     println!("{:?}", with_zero);
            // }

            let mut to_push = vec![VecDeque::<f32>::new(); chans];
            // if start_millis < 1000 || (duration_millis > 0 && start_millis > duration_millis - 1000) {
            for i in 0..able_to_push {
                if with_zero.iter().all(|f| f[i] == 0.) {
                    zeroes += 1;
                    //println!("zero {}", start_millis);
                    continue;
                }
                // for j in 0..chans {
                //     to_push[j].push_back(with_zero[j][i]);
                // }
            }
            //}
            to_push = with_zero;
            // let mut to_push = vec![VecDeque::<f32>::new(); chans];

            if !set {
                call!(self
                    .context_addr
                    .set_buffer(buffer_node_id, to_push, sample_rate as f32))
                .await
                .unwrap();

                set = true;
            } else {
                call!(self.context_addr.push_buffer(buffer_node_id, to_push))
                    .await
                    .unwrap();
            }
        }
        info!("Waiting for thread");
        let res = handle.join().unwrap();
        match res {
            DecoderExit::Eos => {
                self.prev_duration = Some((pushed - end_silence) as f64 / sample_rate);
            }
            _ => {
                self.prev_duration = None;
            }
        }

        info!("decode done");
        return Produces::ok(res);
        //call!(self.queue_addr.load_next());
        //let data = decoded_audio.lock().unwrap();
        // let l = &data[0];
        // let r = if data.len() > 1 { &data[1] } else { &data[0] };
        //println!("{:?}", data);

        // let start_gap = self.find_start_gap(l, r, sample_rate);
        // let end_gap = self.find_end_gap(l, r, sample_rate);
        // info!("start gap = {}", start_gap);
        // info!("end gap = {}", end_gap);

        // Produces::ok(FileInfo {
        //     data: data.to_vec(),
        //     start_gap: 0.,
        //     end_gap: 0.,
        //     duration: 0.,
        //     sample_rate,
        // })
    }

    // async fn get_duration(&self, filename: &str) -> f64 {
    //     info!("Reading duration");
    //     let fakesink = ElementFactory::make("fakesink", None).unwrap();
    //     let bin = ElementFactory::make("playbin", None).unwrap();

    //     bin.set_property("audio-sink", &fakesink).unwrap();
    //     let bus = bin.get_bus().unwrap();
    //     bus.add_signal_watch();
    //     let (sender, mut receiver) = mpsc::channel(32);
    //     let sender_mut = Mutex::new(sender);
    //     let bin_weak = bin.downgrade();
    //     let handler_id = bus
    //         .connect("message", false, move |_| {
    //             let bin = bin_weak.upgrade().unwrap();
    //             if let Some(duration) = bin.query_duration::<ClockTime>() {
    //                 if let Err(msg) = sender_mut.lock().unwrap().try_send(duration) {
    //                     warn!("{:?}", msg);
    //                 }
    //             }

    //             None
    //         })
    //         .unwrap();

    //     bin.set_property("uri", &filename_to_uri(filename, None).unwrap())
    //         .unwrap();

    //     bin.set_state(State::Playing).unwrap();
    //     let duration = receiver.recv().await.unwrap();
    //     let duration_seconds = clocktime_to_seconds(duration);
    //     info!("Got duration {:?}", duration_seconds);
    //     bus.disconnect(handler_id);
    //     bin.set_state(State::Null).unwrap();
    //     return duration_seconds;
    // }

    // fn find_start_gap(&self, l: &Vec<f32>, r: &Vec<f32>, sample_rate: f64) -> f64 {
    //     let duration = l.len();
    //     for i in 0..duration {
    //         if l[i] != 0. || r[i] != 0. {
    //             info!("start i {} dur {}", i, duration);
    //             return i as f64 / sample_rate;
    //         }
    //     }

    //     return duration as f64;
    // }

    // fn find_end_gap(&self, l: &Vec<f32>, r: &Vec<f32>, sample_rate: f64) -> f64 {
    //     let duration = l.len();
    //     for i in (0..duration).rev() {
    //         //info!("i={} r={} l={}", i, r[i], l[i]);
    //         if l[i] != 0. || r[i] != 0. {
    //             info!("i {} dur {}", i, duration);
    //             return (duration - (i + 1)) as f64 / sample_rate;
    //         }
    //     }

    //     return duration as f64;
    // }
}

#[derive(Debug)]
pub struct FileInfo {
    pub data: Vec<Vec<f32>>,
    pub start_gap: f64,
    pub end_gap: f64,
    pub duration: f64,
    pub sample_rate: f64,
}
