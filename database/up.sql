DROP TABLE IF EXISTS students;

CREATE TABLE students (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INTEGER NOT NULL
);

DROP TABLE IF EXISTS tests;

CREATE TABLE tests(
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS questions;

CREATE TABLE questions(
  id VARCHAR(32) PRIMARY KEY,
  test_id VARCHAR(32) NOT NULL,
  question VARCHAR(255) NOT NULL,
  answer VARCHAR(255) NOT NULL,
  FOREIGN KEY (test_id) REFERENCES tests(id)
);

DROP TABLE IF EXISTS attempts;

CREATE TABLE attempts(
  id SERIAL PRIMARY KEY,
  test_id VARCHAR(32) NOT NULL,
  student_id VARCHAR(32) NOT NULL,
  score INT NOT NULL,
  FOREIGN KEY (student_id) REFERENCES students(id),
  FOREIGN KEY (test_id) REFERENCES tests(id)
);

DROP TABLE IF EXISTS answers;

CREATE TABLE answers(
  id SERIAL PRIMARY KEY,
  answer VARCHAR(255) NOT NULL,
  question_id VARCHAR(32) NOT NULL,
  attempts_id INT NOT NULL,
  FOREIGN KEY (question_id) REFERENCES questions(id),
  FOREIGN KEY (attempts_id) REFERENCES attempts(id)
);

DROP TABLE IF EXISTS enrollments;

CREATE TABLE enrollments(
  student_id VARCHAR(32) NOT NULL,
  test_id VARCHAR(32) NOT NULL,
  FOREIGN KEY (student_id) REFERENCES students(id),
  FOREIGN KEY (test_id) REFERENCES tests(id)
);