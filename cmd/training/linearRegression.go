package main

var (
	theta0 float64
	theta1 float64
	sum0   float64
	sum1   float64
	length float64
)

const (
	learningRate = 0.3
	iterations   = 1000
)

func getEstimatePrice(mileage float64) float64 {
	return theta0 + (theta1 * mileage)
}

func linearRegression() {
	length = float64(len(carsNormalized))

	theta0 = 1
	theta1 = 1
	for i := 0; i < iterations; i++ {
		sum0 = 0
		sum1 = 0
		for _, v := range carsNormalized {
			sum0 += getEstimatePrice(v.mileage) - v.price
			sum1 += (getEstimatePrice(v.mileage) - v.price) * v.mileage
		}
		theta0 -= learningRate / length * sum0
		theta1 -= learningRate / length * sum1
	}
	denormalizeThetas()
}
