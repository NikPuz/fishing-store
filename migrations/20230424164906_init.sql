-- +goose Up

SET TIME ZONE '+03';
--
-- таблица `products`
--

CREATE TABLE products (
                           id SERIAL PRIMARY KEY,
                           name varchar(64) NOT NULL,
                           price int NOT NULL CHECK (price > 0),
                           description varchar(128) NOT NULL,
                           stock int NOT NULL,
                           category_id int NOT NULL,
                           manufacturer_id int NOT NULL
                       );

CREATE INDEX product_category_id_index ON products (category_id);
CREATE INDEX product_manufacturer_id_index ON products (manufacturer_id);

--
-- таблица `categories`
--

CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name varchar(64) NOT NULL
);

--
-- таблица `manufacturers`
--

CREATE TABLE manufacturers (
                            id SERIAL PRIMARY KEY,
                            name varchar(64) NOT NULL
);

--
-- таблица `sales`
--

CREATE TABLE sales (
                          id SERIAL PRIMARY KEY,
                          sum int NOT NULL CHECK (sum > 0),
                          cashier_id int NOT NULL,
                          date timestamptz default current_timestamp
);

CREATE INDEX sales_cashier_id_index ON sales (cashier_id);
CREATE INDEX sales_date_index ON sales (date);

--
-- таблица `sales_items`
--

CREATE TABLE sales_items (
                             id SERIAL PRIMARY KEY,
                             sale_id int NOT NULL,
                             product_id int,
                             unit_price int NOT NULL CHECK (unit_price > 0),
                             count int NOT NULL

);

CREATE INDEX sales_items_product_id_index ON sales_items (product_id);
CREATE INDEX sales_items_sale_id_index ON sales_items (sale_id);

--
-- таблица `supplies`
--

CREATE TABLE supplies (
                             id SERIAL PRIMARY KEY,
                             product_id int NOT NULL,
                             unit_price int NOT NULL CHECK (unit_price > 0),
                             count int NOT NULL,
                             date timestamptz default current_timestamp
);

CREATE INDEX supplies_product_id_index ON supplies (product_id);
CREATE INDEX supplies_date_index ON supplies (date);