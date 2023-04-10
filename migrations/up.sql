CREATE TABLE users
(
    id        SERIAL PRIMARY KEY,
    username  VARCHAR(50)  NOT NULL UNIQUE,
    email     VARCHAR(100) NOT NULL UNIQUE,
    name      VARCHAR(50)  NOT NULL,
    last_name VARCHAR(50)  NOT NULL,
    password  VARCHAR(255) NOT NULL
);

CREATE TABLE books
(
    name   VARCHAR(255) PRIMARY KEY,
    author VARCHAR(255) NOT NULL
);

CREATE TABLE book_user
(
    book_name VARCHAR(255),
    user_id   INT,
    returned  BOOLEAN,
    taken_date VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (book_name) REFERENCES books(name) ON DELETE CASCADE
);