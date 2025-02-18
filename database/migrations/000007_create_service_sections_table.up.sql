CREATE TABLE IF NOT EXISTS service_sections (
    id SERIAL PRIMARY KEY,
    path_icon text,
    name varchar(150),
    tagline text null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);