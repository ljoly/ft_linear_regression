package main

import (
	"math"
)

var (
	mean float64
	std  float64
)

func denormalizeThetas() {
	theta0 -= theta1 * mean / std
	theta1 /= std
}

func getStdDeviation() {
	for _, v := range cars {
		abs := math.Abs(v.mileage - mean)
		std += abs * abs
	}
	std /= float64(len(cars))
	std = math.Sqrt(std)
}

func getMean() {
	for _, v := range cars {
		mean += v.mileage
	}
	mean /= float64(len(cars))
}

func normalizeData() {
	getMean()
	getStdDeviation()
	for _, v := range cars {
		carsNormalized = append(carsNormalized, Car{
			mileage: (v.mileage - mean) / std,
			price:   v.price,
		})
	}
}
