CREATE TABLE pokedex (
    id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    name string varchar(255) NOT NULL,
    description string varchar(255) NOT NULL,
    category string varchar(255) NOT NULL,
    type string varchar(255) NOT NULL,
    abilities varchar(255) NOT NULL
);

INSERT INTO pokedex ("name","description", "category", "type" ,"abilities") VALUES
("Bulbasaur","There is a plant seed...", "seed", '{"type1":"grass", "type2":"poison"}', '{"ability1":"overgrow"}');