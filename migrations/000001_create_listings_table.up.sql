CREATE TABLE URL_DB(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    url TEXT NOT NULL,
    short TEXT NOT NULL,
    expiry INTERVAL NOT NULL,
    rate_limit INT NOT NULL,
    rate_limit_rest INTERVAL NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)