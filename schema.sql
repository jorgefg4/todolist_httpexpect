// schema.sql
drop table if exists tasks;

create table tasks (
   id varchar not null primary key,
   name varchar not null,
   check bool not null
);



insert into tasks (name) values ('Prueba1');
insert into users () values ('Prueba2');
