CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE students (
                          id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          full_name     TEXT NOT NULL,
                          email         TEXT NOT NULL UNIQUE,
                          age           INTEGER NOT NULL DEFAULT 0,
                          performance   TEXT NOT NULL DEFAULT '',
                          photo_url     TEXT,
                          city          TEXT,
                          phone         TEXT,
                          format        TEXT NOT NULL DEFAULT 'Онлайн',
                          progress      INTEGER NOT NULL DEFAULT 0,
                          course_id     TEXT,
                          created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
                          updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_students_course_id ON students(course_id);
CREATE INDEX idx_students_created_at ON students(created_at);
