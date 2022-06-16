-- schema.sql
drop table if exists tasks;

create table tasks (
   id varchar not null primary key,
   name varchar not null,
   check boolean not null
);



insert into tasks (id, name, check) values ('1', 'Prueba1', false);
insert into tasks (id, name, check) values ('2', 'Prueba2', true);
