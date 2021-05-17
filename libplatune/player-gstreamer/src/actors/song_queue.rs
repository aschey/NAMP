use crate::actors::queue_loop::QueueLoop;
use act_zero::{runtimes::default::spawn_actor, *};
use async_trait::async_trait;
use postage::prelude::Sink;
use servo_media_audio::decoder::DecoderExit;
use std::{
    sync::mpsc::{sync_channel, Receiver, SyncSender},
    time::Duration,
};
use std::{
    sync::{Arc, Mutex},
    thread,
};

use log::info;
use postage::{broadcast, mpsc::Sender};

use crate::{libplayer::PlayerEvent, util::get_filename_from_path};

use super::{
    decoder::Decoder,
    gstreamer_context::GStreamerContext,
    player::{Player, SenderExt},
};
pub struct SongQueue {
    songs: Vec<String>,
    position: usize,
    player_addr: Addr<Player>,
    context_addr: Addr<GStreamerContext>,
    terminate_tx: Sender<Addr<QueueLoop>>,
    self_addr: Option<Addr<Self>>,
    loop_addr: Option<Addr<QueueLoop>>,
    event_tx: broadcast::Sender<PlayerEvent>,
    shutdown_tx: SyncSender<DecoderExit>,
    shutdown_rx: Arc<Mutex<Receiver<DecoderExit>>>,
}
#[async_trait]
impl Actor for SongQueue {
    async fn started(&mut self, addr: Addr<Self>) -> ActorResult<()> {
        self.self_addr = Some(addr);
        Produces::ok(())
    }
}

impl Drop for SongQueue {
    fn drop(&mut self) {
        println!("drop");
    }
}

impl SongQueue {
    pub fn new(
        player_addr: Addr<Player>,
        context_addr: Addr<GStreamerContext>,
        terminate_tx: Sender<Addr<QueueLoop>>,
        event_tx: broadcast::Sender<PlayerEvent>,
    ) -> SongQueue {
        let (shutdown_tx, shutdown_rx) = sync_channel(32);
        let shutdown_rx = Arc::new(Mutex::new(shutdown_rx));
        SongQueue {
            songs: vec![],
            position: 0,
            player_addr,
            context_addr,
            self_addr: None,
            terminate_tx,
            event_tx,
            shutdown_tx,
            shutdown_rx,
            loop_addr: None,
        }
    }

    pub async fn on_ended(&mut self) {
        info!("Starting ended");
        let should_load_next = call!(self.player_addr.should_load_next()).await.unwrap();
        if should_load_next {
            info!("Queue event: ended. Loading next.");
            //self.load_next().await;

            if self.position == self.songs.len() {
                self.event_tx.publish(PlayerEvent::QueueEnded);
            }
        } else {
            info!("Queue event: ended. Not loading next.");
            self.position = 0;
        }
    }

    pub async fn next(&mut self) {
        if self.position == self.songs.len() - 1 {
            info!("Cannot load next. Already at end of queue");
            return;
        }
        self.position += 1;
        //self.start(None).await;

        self.event_tx.publish(PlayerEvent::Next)
    }

    pub async fn previous(&mut self) {
        if self.position == 0 {
            info!("Cannot load previous. Already at beginning of queue");
            return;
        }
        self.position -= 1;
        //self.start(None).await;

        self.event_tx.publish(PlayerEvent::Previous);
    }

    // pub async fn load_next(&mut self) {
    //     self.position += 1;
    //     call!(self.player_addr.on_ended()).await.unwrap();
    //     self.load_if_exists(self.position + 1, None).await;
    // }

    pub async fn set_queue(&mut self, queue: Vec<String>) {
        info!("Priming queue");
        self.position = 0;

        self.songs = queue.clone();

        self.event_tx.publish(PlayerEvent::StartQueue { queue });
        if let Some(self_addr) = &self.self_addr {
            let self_addr = self_addr.clone();

            let loop_addr = spawn_actor(QueueLoop::new(
                self_addr,
                self.player_addr.clone(),
                self.context_addr.clone(),
                self.shutdown_tx.clone(),
                self.shutdown_rx.clone(),
            ));
            send!(loop_addr.run(None));
            self.loop_addr = Some(loop_addr);
            // send!(self_addr.start(None));
        }
    }

    pub async fn seek(&mut self, millis: u64) {
        info!("start seek");
        //send!(self.player_addr.stop());
        call!(self.player_addr.set_volume_temp(-1.)).await.unwrap();
        self.shutdown_tx.send(DecoderExit::Cancelled).unwrap();
        let old_addr = self.loop_addr.take().unwrap();
        old_addr.termination().await;
        // thread::sleep(Duration::from_secs(5));
        info!("finish cancel");
        // let loop_addr = self.loop_addr.take().unwrap();
        // loop_addr.termination().await;
        // info!("ended");
        // drop(loop_addr);
        // info!("drop");

        //info!("ended.await");
        self.position = 0;

        //call!(self.player_addr.empty());
        // let loop_addr = self.loop_addr.as_ref().unwrap();
        // send!(loop_addr.run(Some(millis)));
        if let Some(self_addr) = &self.self_addr {
            let self_addr = self_addr.clone();

            let loop_addr = spawn_actor(QueueLoop::new(
                self_addr,
                self.player_addr.clone(),
                self.context_addr.clone(),
                self.shutdown_tx.clone(),
                self.shutdown_rx.clone(),
            ));
            send!(loop_addr.run(Some(millis)));
            //let loop_addr_old = self.loop_addr.take().unwrap();
            info!("got old");
            // tokio::spawn(async move {
            //     loop_addr_old.termination().await;
            //     info!("terminated");
            // });
            //loop_addr_old.termination().await;
            //info!("old termination");

            //self.terminate_tx.send(old_addr).await.unwrap();

            self.loop_addr = Some(loop_addr);
            //send!(self_addr.start(None));
        }

        //self.start(Some(millis)).await;
    }

    pub async fn decode_end(&mut self) {
        self.position += 1;
    }

    pub async fn current(&mut self) -> ActorResult<Option<String>> {
        let song = self.songs.get(self.position).map(Into::into);
        Produces::ok(song)
    }

    // async fn prime(&mut self, start_millis: Option<u64>) {
    //     call!(self.player_addr.reset()).await.unwrap();
    //     call!(self.player_addr.ensure_resumed()).await.unwrap();
    //     self.load_if_exists(self.position, start_millis).await;
    //     //self.load_if_exists(self.position + 1).await;
    // }

    // async fn start(&mut self, start_millis: Option<u64>) {
    //     call!(self.player_addr.reset()).await.unwrap();
    //     call!(self.player_addr.ensure_resumed()).await.unwrap();

    //     let decoder_addr = self.decoder_addr.clone();
    //     while let Some(song) = self.songs.get(self.position) {
    //         info!("start decode {}", song);
    //         let (tx, rx) = sync_channel(32);
    //         call!(decoder_addr.decode(song.to_owned(), start_millis, tx, rx))
    //             .await
    //             .unwrap();
    //         info!("done");
    //         self.position += 1;
    //     }
    // }

    // async fn load_if_exists(&mut self, position: usize, start_millis: Option<u64>) {
    //     if let Some(song) = self.songs.get(position) {
    //         if let Some(decoder_addr) = &self.decoder_addr {
    //             let (tx, rx) = sync_channel(32);
    //             let song = song.to_owned();
    //             send!(decoder_addr.decode(song.to_owned(), start_millis, tx, rx));
    //         }

    //         // call!(self.player_addr.load(song.to_owned(), None))
    //         //     .await
    //         //     .unwrap();
    //     }
    // }

    fn current_file_name(&self) -> String {
        get_filename_from_path(&self.songs[self.position])
    }
}
