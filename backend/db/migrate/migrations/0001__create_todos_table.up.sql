CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);
