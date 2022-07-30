create table Foods(
    name varchar(255) PRIMARY KEY NOT NULL,
    description varchar(255)  NOT NULL,
    type varchar(255)  NOT NULL,
    price float  NOT NULL,
    is_vegan bool DEFAULT FALSE,
    is_vegetarian bool DEFAULT FALSE,
    is_spicy bool DEFAULT FALSE
)