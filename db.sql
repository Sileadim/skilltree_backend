CREATE DATABASE skilltree_backend CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE skilltree_backend;
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE on skilltree_backend.* TO 'web'@'localhost';
ALTER USER 'web'@'localhost' INDENTIFIED BY 'password'; 
CREATE TABLE users (
	    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	    name VARCHAR(255) NOT NULL,
	    email VARCHAR(255) NOT NULL,
	    hashed_password CHAR(60) NOT NULL,
	    created DATETIME NOT NULL,
	    active BOOLEAN NOT NULL DEFAULT TRUE
);
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
CREATE TABLE trees (
       	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
        title VARCHAR(100) NOT NULL,
       	uuid VARCHAR(100) NOT NULL,
       	content TEXT NOT NULL,
       	created DATETIME NOT NULL); 
