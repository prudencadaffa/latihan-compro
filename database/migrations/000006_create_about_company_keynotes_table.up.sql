CREATE TABLE IF NOT EXISTS about_company_keynotes (
    id SERIAL PRIMARY KEY,
    about_company_id INT REFERENCES about_companies(id) ON DELETE CASCADE,
    keypoint text,
    path_image text null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_about_company_keynotes_about_company_id ON about_company_keynotes(about_company_id);