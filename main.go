package main

import (
	"fmt"
	"time"
	"github.com/victor-mejia/transfer-service/util"
)

type transfer struct {
	fromCountry string
	fromPerson  string
	destCountry string
	destPerson  string
	date        time.Time
	currency    util.Currency
	amount      float64
}

func main() {

	t := transfer{"Colombia", "81720890", "USA", "52430430", globalTime(), util.COP, 10000000.10}
	fmt.Printf("Transfer details %v \n", t)
}

func globalTime() time.Time {
	return time.Now()
}
