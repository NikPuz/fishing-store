-- +goose Up

SET TIME ZONE 'UTC';

--
-- таблица `product`
--

CREATE TABLE product (
                           id int NOT NULL PRIMARY KEY,
                           name varchar(64) NOT NULL,
                           price numeric CHECK (price > 0),
                           category_id int NOT NULL,
                           manufacturer_id int
                       );

CREATE INDEX product_category_id_index ON product (category_id);
CREATE INDEX product_manufacturer_id_index ON product (manufacturer_id);

--
-- таблица `category`
--

CREATE TABLE category (
                            id int NOT NULL PRIMARY KEY,
                            name varchar(64) NOT NULL
);

--
-- таблица `manufacturer`
--

CREATE TABLE manufacturer (
                            id int NOT NULL PRIMARY KEY,
                            name varchar(64) NOT NULL
);