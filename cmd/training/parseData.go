package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func normalizeData() {
	for _, v := range cars {
		carsNormalized = append(carsNormalized, Car{
			mileage: (v.mileage - minMileage) / (maxMileage - minMileage),
			price:   (v.price - minPrice) / (maxPrice - minPrice),
		})
	}
}

func getMins() {
	minMileage = maxMileage
	minPrice = maxPrice
	for _, v := range cars {
		if v.mileage < minMileage {
			minMileage = v.mileage
		}
		if v.price < minPrice {
			minPrice = v.price
		}
	}
}

func getData() {
	csvFile, _ := os.Open("../../assets/data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		var (
			m float64
			p float64
		)
		fmt.Sscanf(line[0], "%f", &m)
		fmt.Sscanf(line[1], "%f", &p)
		if !(m == 0 && p == 0) {
			cars = append(cars, Car{
				mileage: m,
				price:   p,
			},
			)
		}
		if m > maxMileage {
			maxMileage = m
		}
		if p > maxPrice {
			maxPrice = p
		}
	}
	getMins()
}
