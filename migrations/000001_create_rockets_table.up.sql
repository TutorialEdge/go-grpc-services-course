-- TODO - need to figure out sql syntax for
-- creating tables
CREATE TABLE IF NOT EXISTS rockets(
    id serial NOT NULL PRIMARY KEY,
    type varchar (50),
    name varchar (50)
);

