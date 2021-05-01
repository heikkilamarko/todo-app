SELECT
    id,
    name,
    description,
    created_at,
    updated_at
FROM
    todos.todos
ORDER BY
    created_at DESC
LIMIT $1 OFFSET $2
