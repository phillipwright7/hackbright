CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(25),
    password VARCHAR(50)
);

INSERT INTO users(
    username,
    password
) VALUES (
    'johnappleseed',
    'mypassword123'
);