-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, to_account_id,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;
