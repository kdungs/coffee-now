CREATE TABLE CoffeePeople (
  Id serial PRIMARY KEY,
  Name varchar(50) not null,
  Lat double precision,
  Lng double precision
);

CREATE TABLE CoffeeTokens (
  Id serial PRIMARY KEY,
  Person serial
);

CREATE TABLE CoffeeRequests (
  Id serial PRIMARY KEY,
  Host serial,
  Date timestamp with time zone not null,
  Label varchar(100),
  Lat double precision,
  Lng double precision
);
