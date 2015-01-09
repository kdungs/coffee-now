CREATE TABLE coffeerequests (
  Id serial PRIMARY KEY,
  Host varchar(50) not null,
  Date timestamp with time zone not null,
  Lat double precision,
  Lng double precision
);
