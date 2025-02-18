CREATE TABLE IF NOT EXISTS hero_sections (
    id SERIAL PRIMARY KEY,
    heading varchar(150),
    sub_heading varchar(150),
    path_video text NULL,
    path_banner text NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);