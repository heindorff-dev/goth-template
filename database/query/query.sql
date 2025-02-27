-- name: GetApiErrorLog :one
SELECT * FROM api_error_log WHERE error_id = ? LIMIT 1;

-- name: ListApiErrorLogs :many
SELECT * FROM api_error_log ORDER BY id;

-- name: CreateApiErrorLog :execresult
INSERT INTO api_error_log (
    error_message,
    timestamp
) VALUES (
    ?, ?
);

-- name: DeleteApiErrorLog :exec
DELETE FROM api_error_log WHERE error_id = ?;