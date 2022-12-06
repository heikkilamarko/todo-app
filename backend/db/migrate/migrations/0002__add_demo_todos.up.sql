INSERT INTO todos (name, description, created_at, updated_at)
VALUES
    ('Todo 1', 'This is the first todo.', CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    ('Todo 2', 'This is the second todo.', CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC');
