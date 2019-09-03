-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE posts (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  title text NOT NULL,
  content text NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id),
  CONSTRAINT fk_posts_user_id
    FOREIGN KEY (user_id)
    REFERENCES users(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS posts;
