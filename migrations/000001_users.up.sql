create table users
(
    user_id    serial PRIMARY KEY,
    username   varchar(255) not null,
    password   text not null,
    created_at timestamp   NOT NULL,
    updated_at  timestamp  NULL,
    deleted_at timestamp NULL
);