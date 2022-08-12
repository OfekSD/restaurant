CREATE TABLE Orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    dishes varchar(255)[],
    orderer varchar(255),
    order_time timestamp DEFAULT CURRENT_TIMESTAMP,
    delivered boolean DEFAULT FALSE
)

