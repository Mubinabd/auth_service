create table products
(
    id   uuid primary key,
    name varchar
);

create table orders
(
    id             uuid primary key,
    name           varchar,
    is_paid        bool default false,
    address        varchar,
    delivered_time timestamp
);

create table selected_product
(
    id         uuid primary key,
    order_id   uuid references orders (id),
    product_id uuid references products (id)
);

-- Insert data into the `product` table
INSERT INTO products (id, name) VALUES
   ('550e8400-e29b-41d4-a716-446655440000', 'Widget A'),
   ('550e8400-e29b-41d4-a716-446655440001', 'Widget B'),
   ('550e8400-e29b-41d4-a716-446655440002', 'Widget C');

-- Insert data into the `orders` table
INSERT INTO orders (id, name, is_paid, address, delivered_time) VALUES
    ('740e8400-e29b-11d4-a716-446655440000', 'Order 1', true, '123 Main St', '2023-10-01 10:30:00'),
    ('740e8400-e29b-11d4-a716-446655440001', 'Order 2', false, '456 Elm St', '2023-10-02 14:00:00'),
    ('740e8400-e29b-11d4-a716-446655440002', 'Order 3', true, '789 Oak St', '2023-10-03 16:45:00');

-- Insert data into the `selected_product` table
INSERT INTO selected_product (id, order_id, product_id) VALUES
    ('350e8400-e29b-21d4-a716-446655440000', '740e8400-e29b-11d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440000'),
    ('350e8400-e29b-21d4-a716-446655440001', '740e8400-e29b-11d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440001'),
    ('350e8400-e29b-21d4-a716-446655440002', '740e8400-e29b-11d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440002'),
    ('350e8400-e29b-21d4-a716-446655440003', '740e8400-e29b-11d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440000');
