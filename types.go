package main

import (
	"time"
)

type CoffeePerson struct {
	Id       uint64
	Name     string
	Lat, Lng float64
}

type CoffeeToken struct {
	Id     uint64
	Person CoffeePerson
}

type CoffeeRequest struct {
	Id       uint64
	Host     CoffeePerson
	Date     time.Time
	Label    string
	Lat, Lng float64
}
