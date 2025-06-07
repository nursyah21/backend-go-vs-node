-- name: ListMusics :many
SELECT * FROM musics
ORDER BY title;

-- name: CreateMusic :exec
INSERT INTO musics (
  userid, title, artist, link
) VALUES (
  $1, $2, $3, $4
);

-- name: UpdateMusic :exec
UPDATE musics 
  set title = $3,
  artist = $4,
  link = $5
WHERE
  id = $1 AND userid = $2;

-- name: DeleteMusic :exec
DELETE FROM musics
WHERE id = $1 AND userid = $2;