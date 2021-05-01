DROP TABLE IF EXISTS todos.todos;

CREATE TABLE todos.todos (
    id serial,
    name text NOT NULL,
    description text,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT todos_pkey PRIMARY KEY (id)
);

