use strum_macros::Display;
#[derive(Debug, Clone)]
pub enum Command {
    SetQueue(Vec<String>),
    AddToQueue(String),
    Seek(u64),
    SetVolume(f32),
    Pause,
    Resume,
    Start,
    Stop,
    Ended,
    Next,
    Previous,
    Shutdown,
}

#[derive(Clone, Debug, Display)]
pub enum PlayerEvent {
    StartQueue(Vec<String>),
    QueueUpdated(Vec<String>),
    Stop,
    Pause,
    Resume,
    Ended,
    Next,
    Previous,
    SetVolume(f32),
    Seek(u64),
    QueueEnded,
}
