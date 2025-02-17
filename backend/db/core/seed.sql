-- insert admin role
INSERT INTO role (name, permissions)
VALUES ('admin', '{"manage_restaurant": true, "manage_employees": true, "view_reports": true}')
ON CONFLICT (name) DO NOTHING; -- Ensures 'admin' role exists