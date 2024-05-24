CREATE TYPE roles AS ENUM ('admin', 'user');
CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(32),
    lastname VARCHAR(32),
    phone VARCHAR(13) UNIQUE,
    email VARCHAR(64) UNIQUE,
    age INT,
    role roles,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TYPE tasktype AS ENUM ('todo', 'done', 'habit');

CREATE TABLE task(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(64),
    type tasktype,
    user_id UUID REFERENCES users,
    description TEXT,
    deadline TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

