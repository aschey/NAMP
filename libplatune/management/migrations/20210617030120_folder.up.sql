CREATE TABLE IF NOT EXISTS folder (
    folder_id INTEGER PRIMARY KEY NOT NULL,
    folder_path TEXT NOT NULL UNIQUE
)