mod actors;
mod context;
mod util;
#[cfg(all(feature = "runtime-tokio", feature = "runtime-async-std"))]
compile_error!("features 'runtime-tokio' and 'runtime-async-std' are mutually exclusive");

pub mod libplayer {
    use crate::actors::{
        analyser::Analyser,
        decoder::Decoder,
        gstreamer_context::GStreamerContext,
        player::Player,
        queue_loop::QueueLoop,
        request_handler::{Command, RequestHandler},
        song_queue::SongQueue,
        terminate::Terminate,
    };

    use act_zero::{call, runtimes::default::spawn_actor, Addr};
    use log::info;
    pub use postage::*;
    use servo_media::BackendInit;
    use strum_macros::Display;

    use gstreamer::glib::{self, MainLoop};
    pub use postage::{sink::Sink, stream::Stream};
    //use postage::{broadcast::Sender, mpsc, sink::Sink};
    use std::{
        fmt,
        thread::{self, JoinHandle},
    };

    pub struct PlatunePlayer {
        //glib_main_loop: MainLoop,
        //glib_handle: Option<JoinHandle<()>>,
        cmd_sender: mpsc::Sender<Command>,
    }

    impl PlatunePlayer {
        pub fn create() -> (PlatunePlayer, broadcast::Receiver<PlayerEvent>) {
            PlatunePlayer::create_backend::<servo_media_auto::Backend>()
        }

        pub fn create_dummy() -> (PlatunePlayer, broadcast::Receiver<PlayerEvent>) {
            PlatunePlayer::create_backend::<servo_media_auto::DummyBackend>()
        }

        fn create_backend<T: BackendInit>() -> (PlatunePlayer, broadcast::Receiver<PlayerEvent>) {
            let (event_tx, event_rx) = broadcast::channel(32);
            let (tx, rx) = mpsc::channel(32);
            let (terminate_tx, terminate_rx) = mpsc::channel::<Addr<QueueLoop>>(1000);
            let (analysis_tx, analysis_rx) = mpsc::channel(32);
            let context_addr = spawn_actor(GStreamerContext::new::<T>());

            let player_addr = spawn_actor(Player::new(
                //decoder_addr,
                context_addr.clone(),
                event_tx.clone(),
                analysis_tx,
                tx.clone(),
            ));

            //let decoder_addr = spawn_actor(Decoder::new(context_addr.clone(), player_addr.clone()));
            let terminate_addr = spawn_actor(Terminate::new(terminate_rx));
            let queue_addr = spawn_actor(SongQueue::new(
                player_addr.clone(),
                context_addr.clone(),
                terminate_tx,
                event_tx,
            ));

            let handler_addr =
                spawn_actor(RequestHandler::new(rx, queue_addr.clone(), player_addr));
            let analyser_addr = spawn_actor(Analyser::new(analysis_rx));

            let handler_task = async move {
                call!(handler_addr.run()).await.unwrap();
            };

            let analyser_task = async move {
                call!(analyser_addr.run()).await.unwrap();
            };

            let terminate_task = async move {
                call!(terminate_addr.run()).await.unwrap();
            };

            #[cfg(feature = "runtime-tokio")]
            {
                tokio::spawn(handler_task);
                tokio::spawn(analyser_task);
                tokio::spawn(terminate_task);
            }

            #[cfg(feature = "runtime-async-std")]
            {
                async_std::task::spawn(handler_task);
                async_std::task::spawn(analyser_task);
                async_std::task::spawn(terminate_task);
            }

            // let main_loop = glib::MainLoop::new(None, false);
            // let main_loop_ = main_loop.clone();
            // let glib_handle = thread::spawn(move || {
            //     main_loop_.run();
            // });
            (
                PlatunePlayer {
                    // glib_main_loop: main_loop,
                    // glib_handle: Some(glib_handle),
                    cmd_sender: tx,
                },
                event_rx,
            )
        }

        pub fn set_queue(&mut self, queue: Vec<String>) {
            self.cmd_sender.try_send(Command::SetQueue(queue)).unwrap();
        }

        pub fn seek(&mut self, millis: u64) {
            self.cmd_sender.try_send(Command::Seek(millis)).unwrap();
        }

        pub fn stop(&mut self) {
            self.cmd_sender.try_send(Command::Stop).unwrap();
        }

        pub fn set_volume(&mut self, volume: f32) {
            self.cmd_sender
                .try_send(Command::SetVolume(volume))
                .unwrap();
        }

        pub fn pause(&mut self) {
            self.cmd_sender.try_send(Command::Pause).unwrap();
        }

        pub fn resume(&mut self) {
            self.cmd_sender.try_send(Command::Resume).unwrap();
        }

        pub fn next(&mut self) {
            self.cmd_sender.try_send(Command::Next).unwrap();
        }

        pub fn previous(&mut self) {
            self.cmd_sender.try_send(Command::Previous).unwrap();
        }

        pub fn join(&mut self) {
            self.cmd_sender.try_send(Command::Shutdown).unwrap();
            // self.glib_main_loop.quit();
            // self.glib_handle.take().unwrap().join().unwrap();
        }
    }

    #[derive(Clone, Debug, Display)]
    pub enum PlayerEvent {
        StartQueue { queue: Vec<String> },
        Stop,
        Pause,
        Resume,
        Ended,
        Next,
        Previous,
        SetVolume { volume: f32 },
        Seek { millis: u64 },
        QueueEnded,
    }
}
