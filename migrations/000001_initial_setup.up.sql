CREATE TABLE IF NOT EXISTS tasks(
    id uuid,
    assignee VARCHAR(64),
    title VARCHAR(64),
    summary VARCHAR(512),
    deadline TIMESTAMP NOT NULL,
    status VARCHAR(64),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);