CREATE TABLE foot_prints(
  -- idは不要なのでコメントアウトしとく。消すのめんどくさいから
  -- id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  visitor_uid VARCHAR(255) NOT NULL,
	user_uid VARCHAR(255) NOT NULL,
  unread BOOLEAN NOT NULL DEFAULT '1',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `index_foot_prints_on_visitor_uid` (`visitor_uid`),
  KEY `index_foot_prints_on_user_uid` (`user_uid`),
  UNIQUE INDEX (`visitor_uid`, `user_uid`),
  CONSTRAINT `index_foot_prints_on_visitor_uid` FOREIGN KEY (`visitor_uid`) REFERENCES `users` (`uid`),
  CONSTRAINT `index_foot_prints_on_user_uid` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
