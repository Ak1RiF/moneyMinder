CREATE TABLE records(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    amount FLOAT NOT NULL,
    description TEXT DEFAULT ""
);

CREATE TABLE goals(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    amount FLOAT NOT NULL,
    total_contributed FLOAT DEFAULT 0,
    date_create VATCHAR(255) NOT NULL,
    date_completion VATCHAR(255) NOT NULL
);