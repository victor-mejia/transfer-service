package main

type Currency string

const (
	USD Currency = "USD"
	COP Currency = "COP"
	GBP Currency = "GBP"
)


type cop struct {}
var cop = cop{}
func COP() *Currency {
	return cop
}



