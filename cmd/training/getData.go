package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

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
	}
}
