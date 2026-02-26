INSERT INTO funa_item (name, description)
VALUES ($1, $2)
RETURNING id;
