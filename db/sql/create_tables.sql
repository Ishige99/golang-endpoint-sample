-- テーブル作成
CREATE TABLE article (
     id BIGINT(20) NOT NULL AUTO_INCREMENT,
     title VARCHAR(255) NOT NULL,
     description VARCHAR(255) NOT NULL,
     content TEXT NOT NULL,
     PRIMARY KEY (id)
);