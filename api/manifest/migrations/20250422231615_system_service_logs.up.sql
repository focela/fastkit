-- Enable UUID functionality
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create system service logs table
CREATE TABLE system_service_logs
(
    -- Primary identification
    id                   UUID PRIMARY KEY      DEFAULT gen_random_uuid(),

    -- Trace and severity
    trace_id             VARCHAR(50)           DEFAULT NULL,
    log_level            VARCHAR(32)           DEFAULT NULL,
    message              TEXT                  DEFAULT NULL,
    stack_trace          JSONB                 DEFAULT NULL,
    caller_line          VARCHAR(255) NOT NULL,

    -- Timing
    triggered_at_unix_ns      BIGINT                DEFAULT NULL,

    -- Activity tracking
    is_active            BOOLEAN               NOT NULL DEFAULT TRUE,

    -- Metadata
    created_at           TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP             DEFAULT CURRENT_TIMESTAMP
);

-- Add comments to columns
COMMENT ON COLUMN system_service_logs.id IS 'Log ID';
COMMENT ON COLUMN system_service_logs.trace_id IS 'Trace ID for request tracking';
COMMENT ON COLUMN system_service_logs.log_level IS 'Severity level of the log (e.g. INFO, ERROR)';
COMMENT ON COLUMN system_service_logs.message IS 'Log message content';
COMMENT ON COLUMN system_service_logs.stack_trace IS 'Stack trace in JSON format';
COMMENT ON COLUMN system_service_logs.caller_line IS 'Source code location where the log was triggered';
COMMENT ON COLUMN system_service_logs.triggered_at_unix_ns IS 'Time the event was triggered (in nanoseconds)';
COMMENT ON COLUMN system_service_logs.is_active IS 'Log entry status (active or inactive)';
COMMENT ON COLUMN system_service_logs.created_at IS 'Log creation timestamp';
COMMENT ON COLUMN system_service_logs.updated_at IS 'Last updated timestamp';

-- Create indexes for performance optimization
CREATE INDEX idx_system_service_logs_trace_id ON system_service_logs (trace_id);
CREATE INDEX idx_system_service_logs_log_level ON system_service_logs (log_level);
CREATE INDEX idx_system_service_logs_is_active ON system_service_logs (is_active);

-- Create a function to automatically update the updated_at timestamp
CREATE OR REPLACE FUNCTION update_system_service_logs_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to apply the timestamp update function
CREATE TRIGGER trigger_update_system_service_logs_timestamp
    BEFORE UPDATE
    ON system_service_logs
    FOR EACH ROW
    EXECUTE FUNCTION update_system_service_logs_timestamp();
