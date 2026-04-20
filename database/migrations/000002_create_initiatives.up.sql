CREATE TABLE initiatives (
    id TEXT PRIMARY KEY,
    kind TEXT NOT NULL CHECK (kind IN ('project', 'event'))
);
