package main

func denormalizeThetas() {
	theta0 = theta0*(maxMileage-minMileage) + minMileage
	theta1 = theta1*(maxMileage-minMileage) + minMileage
}

func normalizeData() {
	for _, v := range cars {
		carsNormalized = append(carsNormalized, Car{
			mileage: (v.mileage - minMileage) / (maxMileage - minMileage),
			price:   (v.price - minPrice) / (maxPrice - minPrice),
		})
	}
	// fmt.Println(carsNormalized)
}
