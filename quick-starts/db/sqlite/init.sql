CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(50)
);

CREATE TABLE tasks (
    id INTEGER PRIMARY KEY,
    title VARCHAR(50),
    due_date TIMESTAMP,
    completed BOOLEAN DEFAULT false
);
