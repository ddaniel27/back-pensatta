CREATE TYPE language_enum AS ENUM ('es', 'en', 'pt');

CREATE TABLE IF NOT EXISTS pensatta_institution (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    email VARCHAR(128) NOT NULL,
    country VARCHAR(64) NOT NULL,
    province VARCHAR(64) NOT NULL,
    city VARCHAR(64) NOT NULL,
    code VARCHAR(6) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS pensatta_languages (
    id SERIAL PRIMARY KEY,
    institution_id BIGINT,
    value language_enum DEFAULT 'es' NOT NULL
);
