-- name: GetPermission :one
SELECT * FROM permissions
WHERE id = sqlc.arg(id) LIMIT 1;

-- name: ListPermission :many
SELECT * FROM permissions
ORDER BY title;

-- name: CreatePermission :one
INSERT INTO permissions (
  title
) VALUES (
  sqlc.arg(id)
)
RETURNING *;

-- name: UpdatePermission :exec
UPDATE permissions
  set title = sqlc.arg(title)
WHERE id = sqlc.arg(id);

-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = sqlc.arg(id);