-- +goose Up
CREATE TABLE users (
    id serial not null unique,
    username varchar(128) not null
);

CREATE TABLE chats_users (
    chat_id int not null,
    user_id int not null
);

CREATE TABLE chats (
    id serial not null unique,
    name varchar(128) not null
);

CREATE TABLE messages (
    id serial not null unique,
    user_id int not null,
    chat_id int not null,
    message text not null,
    sending_time timestamp not null
);

CREATE INDEX messages_cid_time_ids ON messages(chat_id, sending_time);

-- +goose Down
DROP INDEX messages_cid_time_ids;
DROP TABLE messages;
DROP TABLE chats_users;
DROP TABLE chats;
DROP TABLE users;