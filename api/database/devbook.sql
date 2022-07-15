CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DELIMITER ;;
CREATE TRIGGER IF NOT EXISTS `users_uuid_trigger` 
BEFORE INSERT ON `users` FOR EACH ROW 
BEGIN
  IF new.id IS NULL THEN
    SET new.id = uuid();
  END IF;
END;;
DELIMITER ;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `followers`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` VARCHAR(255) NOT NULL PRIMARY KEY,
  `first_name` VARCHAR(50) NOT NULL,
  `last_name` VARCHAR(50) NOT NULL,
  `username` VARCHAR(50) NOT NULL UNIQUE,
  `password` VARCHAR(100) NOT NULL,
  `created_at` TIMESTAMP DEFAULT current_timestamp(),
  `updated_at` TIMESTAMP DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` TIMESTAMP NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS `followers`(
  `user_id` VARCHAR(255) NOT NULL,
  `follower_id` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT current_timestamp(),
  `updated_at` TIMESTAMP DEFAULT current_timestamp() ON UPDATE current_timestamp(),

  FOREIGN KEY (`user_id`)
  REFERENCES `users`(`id`)
  ON DELETE CASCADE,

  FOREIGN KEY (`follower_id`)
  REFERENCES `users`(`id`)
  ON DELETE CASCADE,

  PRIMARY KEY(`user_id`, `follower_id`)
);
