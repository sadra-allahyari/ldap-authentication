-- name: CreateUser :execresult
INSERT INTO users (
    dn
) VALUES (
    ?
);

-- name: GetAPIKey :one
SELECT api_key FROM users WHERE dn = ?;

-- name: DeleteAPIKey :exec
DELETE FROM users WHERE api_key = ?;

-- name: GetUserIDByAPIKey :one
SELECT id FROM users WHERE api_key = ?;