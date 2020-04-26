CREATE TABLE IF NOT EXISTS import_temp (
    import_id INTEGER PRIMARY KEY NOT NULL,
    import_song_path TEXT NOT NULL UNIQUE,
    import_artist TEXT NOT NULL,
    import_album_artist TEXT NOT NULL,
    import_title TEXT NOT NULL,
    import_album TEXT NOT NULL,
    import_file_size INTEGER NOT NULL,
    import_track_number INTEGER NOT NULL,
    import_disc_number INTEGER NOT NULL
)