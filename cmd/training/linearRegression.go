package main

import (
	"fmt"
	"math"
)

var (
	theta0    float64
	theta1    float64
	tmpTheta0 float64
	tmpTheta1 float64
	sum0      float64
	sum1      float64
)

func getEstimatePrice(mileage float64) float64 {
	return theta0 + (theta1 * mileage)
}

func linearRegression() {
	tmpTheta0 = 1
	tmpTheta1 = 1

	learningRate := 0.3
	len := len(carsNormalized)

	for math.Abs(tmpTheta0) > 0.0001 && math.Abs(tmpTheta1) > 0.0001 {
		sum0 = 0
		sum1 = 0
		theta0 -= tmpTheta0
		theta1 -= tmpTheta1
		for _, v := range carsNormalized {
			sum0 += getEstimatePrice(v.mileage) - v.price
			sum1 += (getEstimatePrice(v.mileage) - v.price) * v.mileage
		}
		tmpTheta0 = learningRate / float64(len) * sum0
		tmpTheta1 = learningRate / float64(len) * sum1
	}
	// denormalizeThetas()
	fmt.Println(theta0, theta1)
}
