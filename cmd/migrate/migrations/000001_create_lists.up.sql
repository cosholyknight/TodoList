CREATE TABLE IF NOT EXISTS lists(
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL default NOW()
)