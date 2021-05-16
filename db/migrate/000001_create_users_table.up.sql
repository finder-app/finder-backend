-- CREATE TABLE IF NOT EXISTS users(
CREATE TABLE users(
    uid VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,

    -- 前はis_maleで性別を判別してたけど、genderに変更してる
    is_male BOOLEAN NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY(uid),
    UNIQUE KEY `index_users_on_email` (`email`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
