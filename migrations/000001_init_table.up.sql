
ALTER DATABASE postgres SET timezone TO 'Asia/bangkok';
CREATE TABLE
    IF NOT EXISTS attendance (
        id SERIAL NOT NULL PRIMARY KEY,
        user_id TEXT NOT NULL,
        check_in TIMESTAMPTZ,
        check_out TIMESTAMPTZ,
        work_type TEXT NOT NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS breaks (
        id SERIAL NOT NULL PRIMARY KEY,
        user_id TEXT NOT NULL,
        break_in TIMESTAMPTZ,
        break_out TIMESTAMPTZ,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE OR REPLACE FUNCTION UPDATE_UPDATED_AT_COLUMN()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER UPDATE_ATTENDANCE_MODTIME 
    BEFORE UPDATE ON attendance
    FOR EACH ROW
    EXECUTE FUNCTION UPDATE_UPDATED_AT_COLUMN();

CREATE TRIGGER UPDATE_BREAKS_MODTIME 
    BEFORE UPDATE ON breaks
    FOR EACH ROW
    EXECUTE FUNCTION UPDATE_UPDATED_AT_COLUMN();