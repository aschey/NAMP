CREATE TABLE IF NOT EXISTS mount (
    mount_id INTEGER PRIMARY KEY NOT NULL,
    mount_name TEXT NOT NULL UNIQUE,
    mount_path TEXT NOT NULL
)