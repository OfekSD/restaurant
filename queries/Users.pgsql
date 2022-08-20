CREATE TABLE Users(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username varchar(255) not NULL,
    password varchar(255) not NULL,
    is_admin bool DEFAULT FALSE
)