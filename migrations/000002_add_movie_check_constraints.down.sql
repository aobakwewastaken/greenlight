ALTER TABLE movie
DROP CONSTRAINT IF EXISTS movie_runtime_check;

ALTER TABLE movie
DROP CONSTRAINT IF EXISTS movie_year_check;

ALTER TABLE movie
DROP CONSTRAINT IF EXISTS genres_length_check;