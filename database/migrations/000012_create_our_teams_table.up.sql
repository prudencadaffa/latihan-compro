CREATE TABLE IF NOT EXISTS our_teams (
    id SERIAL PRIMARY KEY,
    name  varchar(150),
    path_photo text,
    role varchar(100),
    tagline text,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);