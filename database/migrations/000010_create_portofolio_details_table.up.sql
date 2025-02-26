CREATE TABLE IF NOT EXISTS portofolio_details (
    id SERIAL PRIMARY KEY,
    portofolio_section_id INT REFERENCES portofolio_sections(id) ON DELETE CASCADE,
    category varchar(150),
    client_name text,
    project_date timestamp,
    project_url text NULL,
    title varchar(200),
    description text,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_portofolio_details_portofolio_section_id ON portofolio_details(portofolio_section_id);