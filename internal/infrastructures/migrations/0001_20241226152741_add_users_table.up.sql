CREATE TABLE IF NOT EXISTS Users (
    user_id SERIAL PRIMARY KEY,
    user_firstname VARCHAR(255),
    user_lastname VARCHAR(255), 
    user_email VARCHAR(255),
    user_password VARCHAR(255)
);