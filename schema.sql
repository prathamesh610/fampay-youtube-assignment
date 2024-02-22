CREATE SCHEMA fampay;

CREATE TABLE fampay.search_query (
    id INTEGER PRIMARY KEY,
    query VARCHAR(255) NOT NULL
);

CREATE TABLE fampay.videos (
    id SERIAL PRIMARY KEY UNIQUE,
    search_query VARCHAR(255) NOT NULL,
    video_title VARCHAR(255) NOT NULL,
    description TEXT,
    publishing_date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    video_id VARCHAR(20) NOT NULL
);


CREATE TABLE fampay.thumbnail_urls (
    id SERIAL PRIMARY KEY,
    video_id VARCHAR(20),
    thumbnail_url VARCHAR(255)
);
