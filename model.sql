create table schools(
    id serial not null primary key,
    name character varying(32)

);

create table classes(
    id serial not null primary key,
    name character varying(32),
    schools_id int not null references schools(id)
);

create table pupils(
    id serial not null primary key,
    name character varying(64),
    classes_id int not null references classes(id)
);


INSERT INTO schools(name) values ('Ulugbek'), ('Valixanovich');

INSERT INTO classes(name, schools_id) values ('7-sinf',1), ('7-sinf',2);

INSERT INTO pupils(name, classes_id) values ('Ismoilov Feruz',1), ('Ismoilov Bexruz',1), ('Ismoilov Farrux',2),('Shomurodov Shahzod',2);