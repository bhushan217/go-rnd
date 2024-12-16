-- name: GetPost :one
SELECT * FROM posts
WHERE id = sqlc.arg(id) LIMIT 1;

-- name: ListPost :many
SELECT * FROM posts
ORDER BY updated_at DESC;

-- name: CreatePost :one
INSERT INTO posts (
  title,
  body,
  user_id,
  status
) VALUES (
  sqlc.arg(id),
  sqlc.arg(body),
  sqlc.arg(user_id),
  sqlc.arg(status)
)
RETURNING *;

-- name: UpdatePost :exec
UPDATE posts
  set title = sqlc.arg(title),
  version = version + 1,
  updated_at = CURRENT_TIMESTAMP,
  updated_by = sqlc.arg(updated_by)
WHERE id = sqlc.arg(id) and version = sqlc.arg(version);

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = sqlc.arg(id) and version = sqlc.arg(version);