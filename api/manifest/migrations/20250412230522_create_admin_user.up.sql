-- Enable UUID functionality
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create administrator users table
CREATE TABLE admin_user
(
    -- Primary identification
    id                   UUID PRIMARY KEY      DEFAULT gen_random_uuid(),

    -- Organizational relationships
    dept_id              UUID                  DEFAULT NULL,
    role_id              UUID                  DEFAULT NULL,
    parent_id            UUID                  DEFAULT NULL,

    -- Authentication information
    username             VARCHAR(32)  NOT NULL UNIQUE,
    password_hash        VARCHAR(100) NOT NULL,
    salt                 CHAR(16)     NOT NULL,
    password_reset_token VARCHAR(150)          DEFAULT NULL,

    -- Profile information
    full_name            VARCHAR(64)  NOT NULL DEFAULT '',
    avatar               VARCHAR(255)          DEFAULT NULL,
    gender               SMALLINT              DEFAULT 1,
    email                VARCHAR(100)          DEFAULT NULL,
    mobile               VARCHAR(20)           DEFAULT NULL,
    birthday             DATE                  DEFAULT NULL,

    -- Location information
    city_id              UUID                  DEFAULT NULL,
    address              VARCHAR(255)          DEFAULT NULL,

    -- Account metrics
    integral             NUMERIC(10, 2)        DEFAULT 0,
    balance              NUMERIC(10, 2)        DEFAULT 0,

    -- Hierarchy information
    depth                SMALLINT              DEFAULT 1,
    node_path            TEXT         NOT NULL DEFAULT '',

    -- System features
    invite_code          VARCHAR(12) UNIQUE    DEFAULT NULL,
    cash_config          JSONB                 DEFAULT NULL,

    -- Activity tracking
    last_active_at       TIMESTAMP             DEFAULT NULL,
    is_active            BOOLEAN               DEFAULT TRUE,

    -- Metadata
    remark               TEXT                  DEFAULT NULL,
    created_at           TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP             DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to columns
COMMENT ON COLUMN admin_user.id IS 'Administrator ID';
COMMENT ON COLUMN admin_user.dept_id IS 'Department ID';
COMMENT ON COLUMN admin_user.role_id IS 'Role ID';
COMMENT ON COLUMN admin_user.parent_id IS 'Direct parent administrator ID';
COMMENT ON COLUMN admin_user.username IS 'Username';
COMMENT ON COLUMN admin_user.password_hash IS 'Password (hashed)';
COMMENT ON COLUMN admin_user.salt IS 'Password salt';
COMMENT ON COLUMN admin_user.password_reset_token IS 'Token for resetting password';
COMMENT ON COLUMN admin_user.full_name IS 'Full name';
COMMENT ON COLUMN admin_user.avatar IS 'User avatar URL';
COMMENT ON COLUMN admin_user.gender IS 'Gender: 1=Male, 2=Female, 0=Unknown';
COMMENT ON COLUMN admin_user.email IS 'Email address';
COMMENT ON COLUMN admin_user.mobile IS 'Mobile phone number';
COMMENT ON COLUMN admin_user.birthday IS 'Birthday';
COMMENT ON COLUMN admin_user.city_id IS 'City ID';
COMMENT ON COLUMN admin_user.address IS 'Contact address';
COMMENT ON COLUMN admin_user.integral IS 'User points';
COMMENT ON COLUMN admin_user.balance IS 'User balance';
COMMENT ON COLUMN admin_user.depth IS 'Level in the hierarchy tree';
COMMENT ON COLUMN admin_user.node_path IS 'Full path in the hierarchy tree, e.g. 1/3/7';
COMMENT ON COLUMN admin_user.invite_code IS 'Invitation code';
COMMENT ON COLUMN admin_user.cash_config IS 'Cash withdrawal configuration (as JSON)';
COMMENT ON COLUMN admin_user.last_active_at IS 'Last active time';
COMMENT ON COLUMN admin_user.is_active IS 'User status: active or disabled';
COMMENT ON COLUMN admin_user.remark IS 'Additional remarks';
COMMENT ON COLUMN admin_user.created_at IS 'Record creation time';
COMMENT ON COLUMN admin_user.updated_at IS 'Record last update time';

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
CREATE OR REPLACE FUNCTION update_admin_user_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to apply the timestamp update function
CREATE TRIGGER trigger_update_admin_user_timestamp
    BEFORE UPDATE
    ON admin_user
    FOR EACH ROW
    EXECUTE FUNCTION update_admin_user_timestamp();
