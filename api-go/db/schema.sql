CREATE TABLE musics (
  id     BIGSERIAL PRIMARY KEY,
  userid   BIGSERIAL NOT NULL REFERENCES users(id),
  title  text      NOT NULL,
  artist text      NOT NULL,
  link   text      NOT NULL
);

CREATE TABLE users (
  id     BIGSERIAL PRIMARY KEY,
  username  text      NOT NULL UNIQUE,
  password text      NOT NULL
);