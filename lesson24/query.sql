create table student (
    id serial not null primary key,
    name varchar(25) not null default '',
    age smallint not null default 0,
    email varchar(30) not null unique,
    course varchar(20) not null default '',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at bigint default 0
);

-- alter table book add column student_id int references student(id);

-- insert into student (email) values('email1@mail.com'),('email1@mail.com');

create table book (
    id serial not null primary key,
    name varchar(32) not null,
    author_id int not null references author (id),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at bigint default 0
)

create table author(
    id serial primary key,
    name varchar(32) not null,
    age smallint not null default 0,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at bigint default 0
)

insert into author (name, age) values('Akrom Malik', 35);

insert into book(name, author_id) values('Qo''rqma', 1),('Ona', 1);

select b.id, b.name, author_id, a.name from book as b
join author as a on true;

select a.id, a.name, b.name from author a left join book b on b.author_id=a.id;