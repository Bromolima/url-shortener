CREATE TABLE urls (
    id CHAR(36) PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(10) NOT NULL UNIQUE
);