-- +migrate Up
-- +migrate StatementBegin

create table category (
    id              BIGINT PRIMARY KEY UNIQUE,
    name            varchar(256),
    created_at      TIMESTAMP,
    created_by      VARCHAR(256),
    modified_at     TIMESTAMP,
    modified_by     TIMESTAMP
);

create table books (
    id              BIGINT PRIMARY KEY UNIQUE,
    title           VARCHAR(256) NOT NULL,
    description     VARCHAR(256),
    image_url       VARCHAR(256),
    release_year    INT,
    price           INT,
    total_page      INT,
    thickness       VARCHAR(10),
    category_id     INT references category(id),
    created_at      TIMESTAMP,
    created_by      VARCHAR(256),
    modified_at     TIMESTAMP,
    modified_by     TIMESTAMP
);

create table users (
    id              BIGINT NOT NULL,
    username        varchar(256) NOT NULL UNIQUE,
    password        varchar(20) NOT NULL,
    created_at      TIMESTAMP,
    created_by      VARCHAR(256),
    modified_at     TIMESTAMP,
    modified_by     TIMESTAMP
);

-- +migrate StatementEnd

