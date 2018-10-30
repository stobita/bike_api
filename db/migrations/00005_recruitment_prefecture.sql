-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE recruitment_prefecture (
  id int unsigned NOT NULL AUTO_INCREMENT,
  recruitment_id int NOT NULL,
  prefecture_id int NOT NULL,
  created_at datetime,
  updated_at datetime,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `unique` (recruitment_id, prefecture_id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE recruitment_prefecture;
