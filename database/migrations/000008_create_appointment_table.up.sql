CREATE TABLE IF NOT EXISTS appointment (
    id SERIAL PRIMARY KEY,
    service_id INT REFERENCES service_section(id) ON DELETE CASCADE,
    name varchar(150),
    phone_number varchar(15),
    email varchar(150),
    brief text,
    meet_at timestamp NOT NULL,
    budget DECIMAL(10,1) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_appointment_service_id ON appointment(service_id);