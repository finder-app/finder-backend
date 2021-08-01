CREATE TABLE IF NOT EXISTS messages(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  room_id BIGINT UNSIGNED NOT NULL,
  user_uid VARCHAR(255) NOT NULL,
  text TEXT NOT NULL,
  unread BOOLEAN NOT NULL DEFAULT '1',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,

  PRIMARY KEY (`id`),
  KEY `index_messages_on_room_id` (`room_id`),
  KEY `index_messages_on_user_uid` (`user_uid`),
  CONSTRAINT `index_messages_on_room_id` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`),
  CONSTRAINT `index_messages_on_user_uid` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
