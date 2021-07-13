CREATE TABLE events
(
    event_id   serial PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    city       VARCHAR(100) NOT NULL,
    state      VARCHAR(3)   NOT NULL,
    photo_url  VARCHAR(255) NOT NULL,
    datetime   TIMESTAMP    NOT NULL,
    created_on TIMESTAMP    NOT NULL
);