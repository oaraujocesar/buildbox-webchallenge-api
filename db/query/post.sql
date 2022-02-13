-- name: CreatePost :one
INSERT INTO posts (image_url, name, message)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPost :one
SELECT *
FROM posts
WHERE id = $1
LIMIT 1;

-- name: ListPosts :many
SELECT *
FROM posts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;