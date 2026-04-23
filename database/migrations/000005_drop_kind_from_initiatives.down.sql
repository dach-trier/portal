ALTER TABLE initiatives
ADD COLUMN kind TEXT NOT NULL
CHECK (kind in ('project', 'event'));
