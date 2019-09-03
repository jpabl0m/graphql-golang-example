-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE post_comments (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  post_id INT UNSIGNED NOT NULL,
  content text NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id),
  CONSTRAINT fk_post_comments_user_id
    FOREIGN KEY (user_id)
    REFERENCES users(id),
  CONSTRAINT fk_post_comments_post_id
    FOREIGN KEY (post_id)
    REFERENCES posts(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS post_comments;
