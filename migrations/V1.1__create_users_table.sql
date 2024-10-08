CREATE TYPE role_enum AS ENUM ('STUDENT', 'TEACHER', 'COORDINATOR', 'ADMIN');

CREATE TABLE IF NOT EXISTS pensatta_user (
    id SERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    role role_enum NOT NULL,
    institution_id BIGINT NOT NULL,
    list_number INTEGER NOT NULL,
    date_joined TIMESTAMP WITH TIME ZONE NOT NULL,
    last_login TIMESTAMP WITH TIME ZONE
);
