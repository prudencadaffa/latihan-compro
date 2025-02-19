CREATE TABLE IF NOT EXISTS portofolio_testimonial (
    id SERIAL PRIMARY KEY,
    portofolio_section_id INT REFERENCES portofolio_section(id) ON DELETE CASCADE,
    client_name  varchar(150),
    thumbnail varchar(200),
    message text,
    role varchar(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_portofolio_testimonial_portofolio_section_id ON portofolio_testimonial(portofolio_section_id);