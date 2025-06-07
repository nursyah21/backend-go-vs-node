CREATE TABLE musics (
  id     BIGSERIAL PRIMARY KEY,
  title  text      NOT NULL,
  artist text      NOT NULL,
  link   text      NOT NULL
);

CREATE TABLE users (
  id     BIGSERIAL PRIMARY KEY,
  username  text      NOT NULL,
  password text      NOT NULL
);