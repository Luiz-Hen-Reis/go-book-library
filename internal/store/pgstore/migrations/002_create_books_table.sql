-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS books (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  title VARCHAR(80) NOT NULL,
  author_id UUID NOT NULL REFERENCES authors(id) ON DELETE CASCADE,
  published_year INT
);

---- create above / drop below ----

DROP TABLE IF EXISTS books;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
