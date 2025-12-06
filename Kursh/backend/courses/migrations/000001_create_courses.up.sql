DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_class WHERE relname = 'courses_id_seq') THEN
        CREATE SEQUENCE courses_id_seq;
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS courses (
    id               TEXT PRIMARY KEY DEFAULT nextval('courses_id_seq')::text,
    title            TEXT NOT NULL,
    price            NUMERIC(10,2) NOT NULL,
    description      TEXT NOT NULL,
    full_description TEXT NOT NULL,
    level            TEXT NOT NULL,
    duration_hours   INTEGER NOT NULL,
    photo_url        TEXT,
    materials_url    TEXT,
    prerequisites    TEXT,
    target_audience  TEXT,
    instructor       TEXT,
    rating           NUMERIC(3,1),
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);
