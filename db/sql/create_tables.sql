-- articleテーブルの作成
USE test_db;

CREATE TABLE `article` (
   `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
   `title` VARCHAR(255) NOT NULL,
   `description` VARCHAR(255) NOT NULL,
   `content` TEXT NOT NULL,
   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`)
);

-- テストデータの挿入

INSERT INTO `article` (`title`, `description`, `content`) VALUES
('Title 1', 'Description 1', 'Content 1'),
('Title 2', 'Description 2', 'Content 2'),
('Title 3', 'Description 3', 'Content 3'),
('Title 4', 'Description 4', 'Content 4'),
('Title 5', 'Description 5', 'Content 5'),
('Title 6', 'Description 6', 'Content 6'),
('Title 7', 'Description 7', 'Content 7'),
('Title 8', 'Description 8', 'Content 8'),
('Title 9', 'Description 9', 'Content 9'),
('Title 10', 'Description 10', 'Content 10'),
('Title 11', 'Description 11', 'Content 11'),
('Title 12', 'Description 12', 'Content 12'),
('Title 13', 'Description 13', 'Content 13'),
('Title 14', 'Description 14', 'Content 14'),
('Title 15', 'Description 15', 'Content 15'),
('Title 16', 'Description 16', 'Content 16'),
('Title 17', 'Description 17', 'Content 17'),
('Title 18', 'Description 18', 'Content 18'),
('Title 19', 'Description 19', 'Content 19'),
('Title 20', 'Description 20', 'Content 20');