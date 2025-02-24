CREATE DATABASE todolist;

CREATE TABLE lists (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       is_done BOOLEAN DEFAULT FALSE,
                       list_id INT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       CONSTRAINT fk_list FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE
);
