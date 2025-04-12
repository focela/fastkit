-- Migration: 20250401235844_create_admin_user.down.sql
-- Description: Rollback migration for the admin_user table creation
-- This script removes the admin_user table and all its dependencies

-- Drop the trigger first to avoid dependency errors
DROP TRIGGER IF EXISTS trigger_update_admin_user_timestamp ON admin_user;

-- Drop the timestamp update function
DROP FUNCTION IF EXISTS update_admin_user_timestamp();

-- Drop the admin_user table and all its data
DROP TABLE IF EXISTS admin_user;
