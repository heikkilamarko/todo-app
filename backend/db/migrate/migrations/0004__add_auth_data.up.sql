INSERT INTO roles (name) VALUES ('todo-user'), ('todo-viewer');

INSERT INTO permissions (name) VALUES ('todos.read'), ('todos.write');

INSERT INTO role_permissions (role_id, permission_id)
    SELECT r.id, p.id
    FROM roles r, permissions p
    WHERE r.name = 'todo-user' AND p.name IN ('todos.read', 'todos.write');

INSERT INTO role_permissions (role_id, permission_id)
    SELECT r.id, p.id
    FROM roles r, permissions p
    WHERE r.name = 'todo-viewer' AND p.name IN ('todos.read');
