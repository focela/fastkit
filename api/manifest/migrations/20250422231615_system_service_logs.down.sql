-- Migration: 20250401235900_create_system_service_logs.down.sql
-- Description: Rollback migration for the system_service_logs table creation
-- This script removes the system_service_logs table and all its dependencies

-- Drop the trigger first to avoid dependency errors
DROP TRIGGER IF EXISTS trigger_update_system_service_logs_timestamp ON system_service_logs;

-- Drop the timestamp update function
DROP FUNCTION IF EXISTS update_system_service_logs_timestamp();

-- Drop the system_service_logs table and all its data
DROP TABLE IF EXISTS system_service_logs;
