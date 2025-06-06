-- name: ListMusics :many
SELECT * FROM musics
ORDER BY title;

-- name: CreateMusic :exec
INSERT INTO musics (
  title, artist, link
) VALUES (
  $1, $2, $3
);

-- name: DeleteMusic :exec
DELETE FROM musics
WHERE id = $1;