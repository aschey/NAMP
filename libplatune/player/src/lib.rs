#[cfg(all(feature = "runtime-tokio", feature = "runtime-async-std"))]
compile_error!("features 'runtime-tokio' and 'runtime-async-std' are mutually exclusive");

mod event_loop;
#[cfg(feature = "runtime-tokio")]
mod http_stream_reader;
mod player;
pub mod libplayer {
    use crate::event_loop::{ended_loop, start_loop, Command};
    use log::info;
    pub use postage::*;

    use strum_macros::Display;

    pub use postage::{sink::Sink, stream::Stream};
    //use postage::{broadcast::Sender, mpsc, sink::Sink};
    use std::{
        fmt,
        io::BufReader,
        thread::{self, JoinHandle},
        time::Duration,
    };

    pub struct PlatunePlayer {
        cmd_sender: std::sync::mpsc::Sender<Command>,
    }

    impl PlatunePlayer {
        pub fn new() -> (PlatunePlayer, broadcast::Receiver<PlayerEvent>) {
            let (event_tx, event_rx) = broadcast::channel(32);
            let (tx, rx) = std::sync::mpsc::channel();
            let tx_ = tx.clone();
            let (finish_tx, finish_rx) = std::sync::mpsc::channel();
            let finish_tx_ = finish_tx.clone();

            let main_loop = || start_loop(finish_rx, tx_, event_tx);
            let ended = || ended_loop(rx, finish_tx_);
            #[cfg(feature = "runtime-tokio")]
            {
                tokio::task::spawn_blocking(main_loop);
                tokio::task::spawn_blocking(ended);
            }

            #[cfg(not(feature = "runtime-tokio"))]
            {
                thread::spawn(main_loop);
                thread::spawn(ended);
            }

            (
                PlatunePlayer {
                    cmd_sender: finish_tx,
                },
                event_rx,
            )
        }

        pub fn set_queue(&mut self, queue: Vec<String>) {
            self.cmd_sender.send(Command::SetQueue(queue)).unwrap();
        }

        pub fn seek(&mut self, millis: u64) {
            self.cmd_sender.send(Command::Seek(millis)).unwrap();
        }

        pub fn start(&mut self) {
            self.cmd_sender.send(Command::Start).unwrap();
        }

        pub fn stop(&mut self) {
            self.cmd_sender.send(Command::Stop).unwrap();
        }

        pub fn set_volume(&mut self, volume: f32) {
            self.cmd_sender.send(Command::SetVolume(volume)).unwrap();
        }

        pub fn pause(&mut self) {
            self.cmd_sender.send(Command::Pause).unwrap();
        }

        pub fn resume(&mut self) {
            self.cmd_sender.send(Command::Resume).unwrap();
        }

        pub fn next(&mut self) {
            self.cmd_sender.send(Command::Next).unwrap();
        }

        pub fn previous(&mut self) {
            self.cmd_sender.send(Command::Previous).unwrap();
        }

        pub fn join(&mut self) {
            self.cmd_sender.send(Command::Shutdown).unwrap();
        }
    }

    impl Drop for PlatunePlayer {
        fn drop(&mut self) {
            self.cmd_sender.send(Command::Shutdown).unwrap();
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
