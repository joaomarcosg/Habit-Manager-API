-- Write your migrate up statements here

ALTER TABLE habits
    ALTER COLUMN start_date TYPE TIMESTAMPTZ USING start_date::timestamp,
    ALTER COLUMN target_date TYPE TIMESTAMPTZ USING target_date::timestamp;

---- create above / drop below ----

ALTER TABLE habits
    ALTER COLUMN start_date  TYPE DATE USING start_date::date,
    ALTER COLUMN target_date TYPE DATE USING target_date::date;


-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
