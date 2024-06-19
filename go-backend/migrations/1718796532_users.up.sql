-- CREATE TABLE IF NOT EXISTS users (
--   id BIGSERIAL NOT NULL PRIMARY KEY,
--   org_id BIGINT, 
--   name VARCHAR(255) NOT NULL,
--   email VARCHAR(255) NOT NULL,
--   password VARCHAR(255),
--   display_name VARCHAR(30) DEFAULT NULL,
--   profile_image_url TEXT DEFAULT NULL,
--   soft_delete BOOLEAN DEFAULT FALSE,
--   role_id BIGINT, 
--   hi5_quota_balance INT DEFAULT 0,
--   soft_delete_by BIGINT DEFAULT NULL,
--   soft_delete_on TIMESTAMP DEFAULT NULL,
--   created_at TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'UTC'),
--   CONSTRAINT fk_org_id FOREIGN KEY (org_id) REFERENCES organizations(id),
--   CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(id),
--   CONSTRAINT fk_user_id FOREIGN KEY (soft_delete_by) REFERENCES users(id)
-- );

ALTER TABLE users ADD CONSTRAINT fk_org_id FOREIGN KEY (org_id) REFERENCES organizations(id);

