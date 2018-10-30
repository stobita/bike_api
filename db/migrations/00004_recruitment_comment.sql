-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE recruitment_comment (
  id int NOT NULL AUTO_INCREMENT,
  recruitment_id int NOT NULL,
  user_id int NOT NULL,
  content text NOT NULL,
  created_at datetime,
  updated_at datetime,
  PRIMARY KEY (id),
  INDEX idx_user_id (user_id),
  INDEX idx_recruitment_id (recruitment_id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE recruitment_comment;
