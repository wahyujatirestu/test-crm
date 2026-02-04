CREATE TABLE membership (
    membership_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    updated_date TIMESTAMP,
    updated_by VARCHAR(100)
);

CREATE TABLE contact (
    contact_id SERIAL PRIMARY KEY,
    membership_id INT NOT NULL REFERENCES membership(membership_id) ON DELETE CASCADE,
    contact_type VARCHAR(20) NOT NULL CHECK (contact_type IN ('email','phone')),
    contact_value VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100) NOT NULL,
    updated_date TIMESTAMP,
    updated_by VARCHAR(100)
);
