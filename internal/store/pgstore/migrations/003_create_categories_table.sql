-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    entries INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

---- create above / drop below ----
DROPE TABLE IF EXISTS categories;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
