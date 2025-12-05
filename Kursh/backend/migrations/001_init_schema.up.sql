-- 001_init_schema.up.sql

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE students (
                          id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          full_name   TEXT NOT NULL,
                          email       TEXT NOT NULL UNIQUE,
                          created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE courses (
                         id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         title           TEXT        NOT NULL,
                         price           NUMERIC(10,2) NOT NULL,
                         description     TEXT        NOT NULL,
                         full_description TEXT       NOT NULL,
                         level           TEXT        NOT NULL,
                         duration_hours  INTEGER     NOT NULL,
                         photo_url       TEXT,
                         materials_url   TEXT,
                         prerequisites   TEXT,
                         target_audience TEXT,
                         instructor      TEXT,
                         rating          NUMERIC(3,1),
                         created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE enrollments (
                             id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             student_id      UUID NOT NULL REFERENCES students(id) ON DELETE CASCADE,
                             course_id       UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
                             status          TEXT NOT NULL, -- in_progress | completed | dropped
                             enrolled_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
                             completed_at    TIMESTAMPTZ,
                             UNIQUE(student_id, course_id)
);

CREATE TABLE certificate_templates (
                                       id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                       course_id       UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
                                       name            TEXT NOT NULL,
                                       validity_days   INTEGER,
                                       created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE certificates (
                              id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                              code        TEXT NOT NULL UNIQUE,
                              student_id  UUID NOT NULL REFERENCES students(id) ON DELETE CASCADE,
                              course_id   UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
                              template_id UUID REFERENCES certificate_templates(id),
                              issued_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
                              expires_at  TIMESTAMPTZ,
                              status      TEXT NOT NULL, -- valid | expired | revoked
                              grade       TEXT,
                              created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_certificates_code ON certificates(code);
