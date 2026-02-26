SELECT id, name, description
FROM funa_item
WHERE name LIKE $1
ORDER BY id DESC
LIMIT $2
OFFSET $3
