CREATE TABLE IF NOT EXISTS song (
    song_id INTEGER PRIMARY KEY NOT NULL,
    song_path TEXT NOT NULL UNIQUE,
    modified_date INTEGER NOT NULL,
    last_scanned_date INTEGER NOT NULL,
    artist_id INTEGER NOT NULL,
    song_title TEXT NOT NULL,
    album_id INTEGER NOT NULL,
    track_number INTEGER NOT NULL,
    play_count INTEGER NOT NULL DEFAULT 0,
    disc_number INTEGER NOT NULL,
    song_year INTEGER NOT NULL,
    song_month INTEGER NOT NULL,
    song_day INTEGER NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT 0,
    duration INTEGER NOT NULL,
    sample_rate INTEGER NOT NULL,
    bit_rate INTEGER NOT NULL,
    album_art_path TEXT NULL,
    fingerprint TEXT NOT NULL,
    FOREIGN KEY(artist_id) REFERENCES artist(artist_id),
    FOREIGN KEY(album_id) REFERENCES album(album_id)
)