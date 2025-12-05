-- 001_init_schema.down.sql

DROP INDEX IF EXISTS idx_certificates_code;

DROP TABLE IF EXISTS certificates;
DROP TABLE IF EXISTS certificate_templates;
DROP TABLE IF EXISTS enrollments;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS students;

-- расширение обычно не дропают в down, но можно при желании:
-- DROP EXTENSION IF EXISTS "uuid-ossp";
