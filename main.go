package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Car struct {
	Mileage int `json:"mileage"`
	Price   int `json:"price"`
}

func main() {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var cars []Car
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("COUCOU")
			log.Fatal(error)
		}
		var (
			m int
			p int
		)
		fmt.Sscanf(line[0], "%d", &m)
		fmt.Sscanf(line[1], "%d", &p)
		cars = append(cars, Car{
			Mileage: m,
			Price:   p,
		},
		)
	}
	carsJSON, _ := json.Marshal(cars)
	fmt.Println(string(carsJSON))
}
