-- name: GetRole :one
SELECT * FROM roles
WHERE id = sqlc.arg(id) LIMIT 1;

-- name: ListRole :many
SELECT * FROM roles
ORDER BY title;

-- name: CreateRole :one
INSERT INTO roles (
  title
) VALUES (
  sqlc.arg(id)
)
RETURNING *;

-- name: UpdateRole :exec
UPDATE roles
  set title = sqlc.arg(title)
WHERE id = sqlc.arg(id) and version = sqlc.arg(version);

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = sqlc.arg(id) and version = sqlc.arg(version);