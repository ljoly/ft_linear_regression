package main

import "fmt"

var (
	theta0    float64
	theta1    float64
	tmpTheta0 float64
	tmpTheta1 float64
)

func getEstimatePrice(mileage float64) float64 {
	return theta0 + (theta1 * mileage)
}

func linearRegression() {
	learningRate := 0.1
	len := len(carsNormalized)

	for _, v := range carsNormalized {
		theta0 = theta0 - tmpTheta0
		theta1 = theta1 - tmpTheta1

		tmpTheta0 = learningRate / float64(len) * (getEstimatePrice(v.mileage) - v.price)
		tmpTheta1 = learningRate / float64(len) * (getEstimatePrice(v.mileage) - v.price) * v.mileage
	}
	fmt.Println(theta0, theta1)
}
