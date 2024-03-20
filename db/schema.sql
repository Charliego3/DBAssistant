CREATE TABLE IF NOT EXISTS connection(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type TEXT NOT NULL,
    name TEXT NOT NULL,
    host TEXT NOT NULL,
    user TEXT NOT NULL,
    pass TEXT NOT NULL,
    database TEXT NOT NULL,
    parent INTEGER NOT NULL
);
