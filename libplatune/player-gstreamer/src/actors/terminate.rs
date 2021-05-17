use std::thread::JoinHandle;

use super::queue_loop::QueueLoop;
use act_zero::*;
use log::info;
use postage::{mpsc::Receiver, prelude::Stream};
use servo_media_audio::decoder::DecoderExit;

pub struct Terminate {
    receiver: Receiver<Addr<QueueLoop>>,
}

impl Actor for Terminate {}

impl Terminate {
    pub fn new(receiver: Receiver<Addr<QueueLoop>>) -> Self {
        Self { receiver }
    }
    pub async fn run(&mut self) {
        while let Some(handle) = self.receiver.recv().await {
            info!("awaiting termination");
            handle.termination().await;
            info!("termination complete");
        }
    }
}
