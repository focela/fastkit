-- Enable UUID functionality
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create administrator users table
CREATE TABLE admin_user
(
    -- Primary identification
    id                   UUID PRIMARY KEY      DEFAULT uuid_generate_v4() COMMENT 'Administrator ID',

    -- Organizational relationships
    dept_id              UUID                  DEFAULT NULL COMMENT 'Department ID',
    role_id              UUID                  DEFAULT NULL COMMENT 'Role ID',
    parent_id            UUID                  DEFAULT NULL COMMENT 'Direct parent administrator ID',

    -- Authentication information
    username             VARCHAR(32)  NOT NULL UNIQUE COMMENT 'Username',
    password_hash        VARCHAR(100) NOT NULL COMMENT 'Password (hashed)',
    salt                 CHAR(16)     NOT NULL COMMENT 'Password salt',
    password_reset_token VARCHAR(150)          DEFAULT NULL COMMENT 'Token for resetting password',

    -- Profile information
    full_name            VARCHAR(64)  NOT NULL DEFAULT '' COMMENT 'Full name',
    avatar               VARCHAR(255)          DEFAULT NULL COMMENT 'User avatar URL',
    gender               SMALLINT              DEFAULT 1 COMMENT 'Gender: 1=Male, 2=Female, 0=Unknown',
    email                VARCHAR(100)          DEFAULT NULL COMMENT 'Email address',
    mobile               VARCHAR(20)           DEFAULT NULL COMMENT 'Mobile phone number',
    birthday             DATE                  DEFAULT NULL COMMENT 'Birthday',

    -- Location information
    city_id              UUID                  DEFAULT NULL COMMENT 'City ID',
    address              VARCHAR(255)          DEFAULT NULL COMMENT 'Contact address',

    -- Account metrics
    integral             NUMERIC(10, 2)        DEFAULT 0 COMMENT 'User points',
    balance              NUMERIC(10, 2)        DEFAULT 0 COMMENT 'User balance',

    -- Hierarchy information
    depth                SMALLINT              DEFAULT 1 COMMENT 'Level in the hierarchy tree',
    node_path            TEXT         NOT NULL DEFAULT '' COMMENT 'Full path in the hierarchy tree, e.g. 1/3/7',

    -- System features
    invite_code          VARCHAR(12) UNIQUE    DEFAULT NULL COMMENT 'Invitation code',
    cash_config          JSONB                 DEFAULT NULL COMMENT 'Cash withdrawal configuration (as JSON)',

    -- Activity tracking
    last_active_at       TIMESTAMP             DEFAULT NULL COMMENT 'Last active time',
    is_active            BOOLEAN               DEFAULT TRUE COMMENT 'User status: active or disabled',

    -- Metadata
    remark               TEXT                  DEFAULT NULL COMMENT 'Additional remarks',
    created_at           TIMESTAMP             DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation time',
    updated_at           TIMESTAMP             DEFAULT CURRENT_TIMESTAMP COMMENT 'Record last update time'
);

-- Create indexes for performance optimization
CREATE INDEX idx_admin_user_email ON admin_user (email);
CREATE INDEX idx_admin_user_mobile ON admin_user (mobile);
CREATE INDEX idx_admin_user_parent_id ON admin_user (parent_id);
CREATE INDEX idx_admin_user_path ON admin_user (node_path);

-- Add additional indexes for common query patterns
CREATE INDEX idx_admin_user_role_id ON admin_user (role_id);
CREATE INDEX idx_admin_user_dept_id ON admin_user (dept_id);
CREATE INDEX idx_admin_user_is_active ON admin_user (is_active);

-- Create a function to automatically update the updated_at timestamp
CREATE
OR REPLACE FUNCTION update_admin_user_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at
= CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$
LANGUAGE plpgsql;

-- Create a trigger to apply the timestamp update function
CREATE TRIGGER trigger_update_admin_user_timestamp
    BEFORE UPDATE
    ON admin_user
    FOR EACH ROW
    EXECUTE FUNCTION update_admin_user_timestamp();