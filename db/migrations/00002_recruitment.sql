-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE recruitment (
  id int NOT NULL AUTO_INCREMENT,
  user_id int NOT NULL,
  title text NOT NULL,
  content text NOT NULL,
  image text,
  min_engine_capacity int,
  max_engine_capacity int,
  min_age int,
  max_age int,
  created_at datetime,
  updated_at datetime,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE recruitment;
