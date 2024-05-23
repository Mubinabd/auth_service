CREATE TYPE gender_type AS ENUM ('m', 'f');


CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    phone_number VARCHAR(30) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    gender gender_type NOT NULL,

    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS resumes(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    position VARCHAR NOT NULL,
    experience INT NOT NULL,
    description TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),

    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS companies(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    location VARCHAR NOT NULL,
    workers INT DEFAULT 0,

    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS recruiters(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    birthday TIMESTAMP,
    gender gender_type NOT NULL,
    company_id UUID NOT NULL REFERENCES companies(id),

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS vacancies(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    position VARCHAR NOT NULL,
    min_exp INT NOT NULL,
    company_id UUID NOT NULL REFERENCES companies(id),
    description TEXT DEFAULT NULL,

    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS interviews(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    vacancy_id UUID REFERENCES vacancies(id),
    recruiter_id UUID REFERENCES recruiters(id),
    interview_date TIMESTAMP NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);


insert into vacancies (name, position, min_exp, company_id, description) values ('We need frontend Developers', 'Frontend Developers', 6, 'Min exp 6 years and ee hire frontend developers In the company')

insert into interviews (user_id, vacancy_id, recruiter_id, interview_date) values ('467107d1-29a9-4725-a85e-6d976709d5cb', '89b2a7d4-a1f1-44a3-813c-0769878feb24', '', '2024-06-14');

