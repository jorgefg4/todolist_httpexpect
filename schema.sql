-- schema.sql
drop table if exists tasks;

create table tasks (
   id serial not null primary key,
   name varchar not null,
   check_valid boolean not null
);