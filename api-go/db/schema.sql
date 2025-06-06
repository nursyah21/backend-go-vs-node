CREATE TABLE musics (
  id     BIGSERIAL PRIMARY KEY,
  title  text      NOT NULL,
  artist text      NOT NULL,
  link   text      NOT NULL
);