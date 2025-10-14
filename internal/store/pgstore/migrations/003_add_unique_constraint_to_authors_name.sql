-- Write your migrate up statements here

ALTER TABLE authors
ADD CONSTRAINT authors_name_unique UNIQUE (name);


---- create above / drop below ----

ALTER TABLE authors
DROP CONSTRAINT IF EXISTS authors_name_unique;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
