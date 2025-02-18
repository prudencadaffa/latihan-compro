CREATE TABLE IF NOT EXISTS service_details (
    id SERIAL PRIMARY KEY,
    service_id INT REFERENCES service_sections(id) ON DELETE CASCADE,
    path_image text NOT NULL,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    path_pdf text NULL,
    path_docx text NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_service_details_service_id ON service_details(service_id);