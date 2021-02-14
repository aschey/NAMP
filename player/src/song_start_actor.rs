use gstreamer::{Clock, ClockExtManual, ClockId, ClockTime, SystemClock};
use std::{thread, time::Duration};
use tokio::sync::mpsc::Sender;

use crate::{player_actor::PlayerCommand, song_queue_actor::QueueItem};

pub struct SongStartActor {
    pub clock_id: Option<ClockId>,
    pub clock: Clock,
    subscriber: Sender<PlayerCommand>,
    next_song: QueueItem,
}

impl SongStartActor {
    pub fn new(subscriber: Sender<PlayerCommand>) -> SongStartActor {
        SongStartActor {
            clock: SystemClock::obtain(),
            clock_id: None,
            subscriber,
            next_song: QueueItem {
                uri: "".to_owned(),
                position: 0,
            },
        }
    }

    pub fn recv_queue_item(&mut self, item: QueueItem) {
        self.next_song = item;
    }

    pub async fn handle(&mut self, nseconds: i64, player_id: usize) -> () {
        if let Some(shot) = &self.clock_id {
            shot.unschedule();
        }

        let clock_id = self
            .clock
            .new_single_shot_id(ClockTime::from_nseconds(nseconds as u64))
            .unwrap();

        let subscriber = self.subscriber.clone();
        println!("{:?}", player_id);

        thread::sleep(Duration::from_millis(50));
        subscriber
            .send(PlayerCommand::SetUri {
                id: player_id,
                item: self.next_song.clone(),
            })
            .await
            .unwrap();

        subscriber
            .send(PlayerCommand::SchedulePlay {
                id: player_id,
                clock_id,
            })
            .await
            .unwrap();
        //self.clock_id = Some(clock_id);
    }
}

#[derive(Debug)]
pub enum SongStartCommand {
    Schedule { nseconds: i64, player_id: usize },
    RecvItem { item: QueueItem },
}
