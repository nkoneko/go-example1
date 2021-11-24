GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';

CREATE DATABASE app;
GRANT ALL PRIVILEGES ON app.* TO 'appuser'@'%';
USE app;

CREATE TABLE books (id INT AUTO_INCREMENT, title VARCHAR(256), author VARCHAR(128), PRIMARY KEY(id));
INSERT INTO books (title, author) VALUES ('title1', 'author1'), ('title2', 'author2');
