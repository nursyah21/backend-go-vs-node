-- name: CreateMusicsTable :exec
CREATE TABLE IF NOT EXISTS musics (
  id     BIGSERIAL PRIMARY KEY,
  title  TEXT NOT NULL,
  artist TEXT NOT NULL,
  link   TEXT NOT NULL
);