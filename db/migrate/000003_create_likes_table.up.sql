CREATE TABLE likes(
  sent_user_uid VARCHAR(255) NOT NULL,
	recieved_user_uid VARCHAR(255) NOT NULL,
  consented BOOLEAN NOT NULL DEFAULT '0',
  unread BOOLEAN NOT NULL DEFAULT '1',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `index_likes_on_sent_user_uid` (`sent_user_uid`),
  KEY `index_likes_on_recieved_user_uid` (`recieved_user_uid`),
  UNIQUE INDEX (`sent_user_uid`, `recieved_user_uid`),
  CONSTRAINT `index_likes_on_sent_user_uid` FOREIGN KEY (`sent_user_uid`) REFERENCES `users` (`uid`),
  CONSTRAINT `index_likes_on_recieved_user_uid` FOREIGN KEY (`recieved_user_uid`) REFERENCES `users` (`uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
