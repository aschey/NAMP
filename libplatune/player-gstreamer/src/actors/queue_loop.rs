use super::{
    decoder::Decoder, gstreamer_context::GStreamerContext, player::Player, song_queue::SongQueue,
};
use act_zero::*;
use async_trait::async_trait;
use log::info;
use runtimes::default::spawn_actor;
use servo_media_audio::decoder::DecoderExit;
use std::sync::{
    mpsc::{Receiver, SyncSender},
    Arc, Mutex,
};

pub struct QueueLoop {
    queue_addr: Addr<SongQueue>,
    player_addr: Addr<Player>,
    context_addr: Addr<GStreamerContext>,
    tx: SyncSender<DecoderExit>,
    rx: Arc<Mutex<Receiver<DecoderExit>>>,
}

#[async_trait]
impl Actor for QueueLoop {
    async fn error(&mut self, error: ActorError) -> bool {
        info!("{}", error);
        true
    }

    async fn started(&mut self, _addr: Addr<Self>) -> ActorResult<()>
    where
        Self: Sized,
    {
        Produces::ok(())
    }
}

impl QueueLoop {
    pub fn new(
        queue_addr: Addr<SongQueue>,
        player_addr: Addr<Player>,
        context_addr: Addr<GStreamerContext>,
        tx: SyncSender<DecoderExit>,
        rx: Arc<Mutex<Receiver<DecoderExit>>>,
    ) -> Self {
        Self {
            queue_addr,
            player_addr,
            context_addr,
            tx,
            rx,
        }
    }
    pub async fn run(&mut self, start_millis: Option<u64>) -> ActorResult<()> {
        let decoder_addr = spawn_actor(Decoder::new(
            self.context_addr.clone(),
            self.player_addr.clone(),
        ));

        info!("start loop");

        //if start_millis.is_none() {
        call!(self.player_addr.reset()).await.unwrap();
        call!(self.player_addr.ensure_resumed()).await.unwrap();
        //}

        while let Some(song) = call!(self.queue_addr.current()).await.unwrap() {
            info!("Next loop {}", song);

            let res = call!(decoder_addr.decode(
                song.to_owned(),
                start_millis,
                self.tx.clone(),
                self.rx.clone()
            ))
            .await
            .unwrap();

            info!("Finish loop");
            match res {
                DecoderExit::Eos => {
                    info!("Got EOS");

                    //call!(self.queue_addr.decode_end()).await.unwrap();
                }
                _ => {
                    return Err("".into());
                    //return;
                }
            }
        }
        info!("Loop done");
        Err("".into())
    }
}
