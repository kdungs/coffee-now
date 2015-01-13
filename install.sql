CREATE TABLE CoffeePeople (
  Id serial PRIMARY KEY,
  Name varchar(50) not null,
  Lat double precision,
  Lng double precision
);

CREATE TABLE CoffeeTokens (
  Id serial PRIMARY KEY,
  Person serial not null
);

CREATE TABLE CoffeeRequests (
  Id serial PRIMARY KEY,
  Host serial not null,
  Date timestamp with time zone not null,
  Label varchar(100) not null,
  Lat double precision,
  Lng double precision
);

CREATE TABLE CoffeeResponses {
  Id serial PRIMARY KEY,
  Request serial not null,
  Person serial not null
}
