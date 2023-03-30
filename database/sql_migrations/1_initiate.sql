-- +migrate Up
-- +StatementBegin

CREATE TABLE account (
  id SERIAL PRIMARY KEY,
  full_name VARCHAR(256),
  username VARCHAR(256) UNIQUE NOT NULL CHECK (username <> ''),
  password VARCHAR(256) NOT NULL CHECK (password <> ''),
  email VARCHAR(256) UNIQUE NOT NULL CHECK (email <> ''),
  role VARCHAR(10) NOT NULL CHECK (role <> ''),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  class_id INT NOT NULL
);

CREATE TABLE class (
  id SERIAL PRIMARY KEY,
  name VARCHAR(10) UNIQUE NOT NULL CHECK (name <> ''),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  teacher_id INT NOT NULL CHECK (teacher_id <> 0)
);

CREATE TABLE subject (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) UNIQUE NOT NULL CHECK (name <> ''),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE question (
  id SERIAL PRIMARY KEY,
  title VARCHAR(50) NOT NULL CHECK (title <> ''),
  question VARCHAR(1000) NOT NULL CHECK (question <> ''),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  user_role VARCHAR(10) NOT NULL CHECK (user_role <> ''),
  class_id INT NOT NULL CHECK (class_id <> 0),
  user_id INT NOT NULL CHECK (user_id <> 0),
  subject_id INT NOT NULL CHECK (subject_id <> 0)
);

CREATE TABLE answer (
  id SERIAL PRIMARY KEY,
  answer VARCHAR(100) NOT NULL CHECK (answer <> ''),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  user_role VARCHAR(10) NOT NULL CHECK (user_role <> ''),
  question_id INT NOT NULL CHECK (question_id <> 0),
  user_id INT NOT NULL CHECK (user_id <> 0)
);

-- +migrate StatementEnd