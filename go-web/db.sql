create table products (
	id serial primary key,
	name varchar,
	description varchar,
	price decimal,
	quantity integer
);

insert into products (name, description, price, quantity) values ('T-Shirt', 'Funny Science t-shirt', 19.90, 10);
insert into products (name, description, price, quantity) values ('Laptop Dell', 'Very slim and fast process', 1299.90, 2);
insert into products (name, description, price, quantity) values ('Headphone', 'Excelent sound - bass improviment', 39.90, 1);

