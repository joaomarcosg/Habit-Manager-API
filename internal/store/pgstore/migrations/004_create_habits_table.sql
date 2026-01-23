-- Write your migrate up statements here

CREATE TYPE weekday AS ENUM (
    'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturnday', 'sunday'
);

CREATE TABLE IF NOT EXISTS habits (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    category VARCHAR(100) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    frequency weekday[] NOT NULL,
    start_date DATE,
    target_date DATE,
    priority SMALLINT NOT NULL CHECK (priority BETWEEN 0 AND 10),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

---- create above / drop below ----

DROP TABLE IF EXISTS habits;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
