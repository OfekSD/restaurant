CREATE TABLE Orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    dishes varchar(255)[] not null check (public.check_foreign_key_array(dishes, 'public', 'foods', 'name')),
    orderer varchar(255),
    order_time timestamp DEFAULT CURRENT_TIMESTAMP,
    delivered boolean DEFAULT FALSE
)

