-- +goose Up

SET TIME ZONE 'UTC';

--
-- таблица `product`
--

CREATE TABLE product (
                           id SERIAL PRIMARY KEY,
                           name varchar(64) NOT NULL,
                           price numeric NOT NULL CHECK (price > 0),
                           stock int NOT NULL,
                           category_id int NOT NULL,
                           manufacturer_id int
                       );

CREATE INDEX product_category_id_index ON product (category_id);
CREATE INDEX product_manufacturer_id_index ON product (manufacturer_id);

--
-- таблица `category`
--

CREATE TABLE category (
                            id SERIAL PRIMARY KEY,
                            name varchar(64) NOT NULL
);

--
-- таблица `manufacturer`
--

CREATE TABLE manufacturer (
                            id SERIAL PRIMARY KEY,
                            name varchar(64) NOT NULL
);