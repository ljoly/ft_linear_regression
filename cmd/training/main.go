package main

import (
	"flag"
)

type Car struct {
	mileage float64
	price   float64
}

var (
	cars           []Car
	carsNormalized []Car
	mileages       []float64
	prices         []float64
)

var flagVisualizer *bool

func main() {
	flagVisualizer = flag.Bool("v", false, "Plot data in a graph")
	flag.Parse()
	getData()
	normalizeData()
	linearRegression()
	saveThetas()
	if *flagVisualizer {
		print()
	}
}
