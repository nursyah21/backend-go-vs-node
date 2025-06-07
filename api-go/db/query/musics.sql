-- name: ListMusics :many
SELECT * FROM musics
ORDER BY title;

-- name: CreateMusic :exec
INSERT INTO musics (
  title, artist, link
) VALUES (
  $1, $2, $3
);

-- name: UpdateMusic :exec
UPDATE musics 
  set title = $2,
  artist = $3,
  link = $4
WHERE
  id = $1;

-- name: DeleteMusic :exec
DELETE FROM musics
WHERE id = $1;