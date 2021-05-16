CREATE TABLE foot_prints(
  visitor_uid VARCHAR(255) NOT NULL,
	user_uid VARCHAR(255) NOT NULL,
  unread BOOLEAN NOT NULL DEFAULT '1',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `index_foot_prints_on_visitor_uid` (`visitor_uid`),
  KEY `index_foot_prints_on_user_uid` (`user_uid`),
  UNIQUE INDEX (`visitor_uid`, `user_uid`),
  CONSTRAINT `index_foot_prints_on_visitor_uid` FOREIGN KEY (`visitor_uid`) REFERENCES `users` (`uid`),
  CONSTRAINT `index_foot_prints_on_user_uid` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
