CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE certificates (
    id          TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
    code        TEXT NOT NULL UNIQUE,
    student_id  TEXT NOT NULL,
    course_id   TEXT NOT NULL,
    issued_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    grade       TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_certificates_code ON certificates(code);
CREATE INDEX idx_certificates_student ON certificates(student_id);
