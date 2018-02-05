package main

import (
	"fmt"
	"time"
)

type transfer struct {
	fromCountry string
	fromPerson string
	destCountry string
	destPerson string
	date       time.Time
	currency   Currency
	amount     float64
}

func main() {

	t := transfer{"Colombia","81720890","USA","52430430", globalTime(), COP, 10000000.10}
	fmt.Printf("Transfer details %v \n", t)
}

func globalTime() time.Time {
	return time.Now()
}
