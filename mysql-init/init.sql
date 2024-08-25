DROP DATABASE IF EXISTS gobooks;
--
CREATE DATABASE gobooks;
--
USE gobooks;
CREATE TABLE books (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  genre VARCHAR(255) NOT NULL
);
--
INSERT INTO books (title, author, genre)
VALUES ("Nausea", "Jean-Paul Sartre", "Philosophy"),
  ("The Stranger", "Albert Camus", "Philosophy"),
  (
    "Thus Spoke Zarathustra",
    "Friedrich Nietzsche",
    "Philosophy"
  );