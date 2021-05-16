CREATE TABLE rooms_users(
  room_id BIGINT UNSIGNED NOT NULL,
  user_uid VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `index_rooms_on_room_id` (`room_id`),
  KEY `index_rooms_on_user_uid` (`user_uid`),
  CONSTRAINT `index_rooms_on_room_id` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`),
  CONSTRAINT `index_rooms_on_user_uid` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
