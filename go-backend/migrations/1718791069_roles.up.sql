-- create roles table
CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255)
);

-- add user role
INSERT INTO roles(name) VALUES('user');

-- add admin role   
INSERT INTO roles(name) VALUES('admin');

-- add super_admin role
INSERT INTO roles(name) VALUES('super_admin');