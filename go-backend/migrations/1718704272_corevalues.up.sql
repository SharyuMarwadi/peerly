CREATE TABLE core_values (
    id SERIAL PRIMARY KEY,
    org_id INTEGER NOT NULL,
    text VARCHAR(225) NOT NULL,
    thumbnail_url VARCHAR(225),
    description VARCHAR(225),
    parent_id INTEGER,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp,
    soft_delete BOOLEAN DEFAULT false,
    soft_delete_by BIGINT,
    created_by BIGINT,
    CONSTRAINT unique_org_core UNIQUE (org_id, id),
    CONSTRAINT fk_org_id FOREIGN KEY (org_id) REFERENCES organizations(id),
    CONSTRAINT fk_parent_id FOREIGN KEY (parent_id) REFERENCES core_values(id),
    CONSTRAINT fk_soft_delete_by FOREIGN KEY (soft_delete_by) REFERENCES users(id),
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id)
);