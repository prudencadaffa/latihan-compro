CREATE TABLE IF NOT EXISTS portofolio_sections (
    id SERIAL PRIMARY KEY,
    name varchar(150),
    tagline text,
    thumbnail text null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);