-- Write your migrate up statements here
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS authors (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name VARCHAR(50) NOT NULL,
  bio TEXT
);

---- create above / drop below ----

DROP TABLE IF EXISTS authors;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
