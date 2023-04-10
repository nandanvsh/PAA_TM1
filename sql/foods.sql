create table if not exists foods (
    id serial primary key,
    user_id int references users(id),
    name varchar(50) not null,
    price int not null,
    description varchar(255) not null
);