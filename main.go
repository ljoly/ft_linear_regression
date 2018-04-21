package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
)

type Car struct {
	Mileage int `json:"mileage"`
	Price   int `json:"price"`
}

var cars []Car

func getData() {
	csvFile, _ := os.Open("./assets/data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		var (
			m int
			p int
		)
		fmt.Sscanf(line[0], "%d", &m)
		fmt.Sscanf(line[1], "%d", &p)
		if !(m == 0 && p == 0) {
			cars = append(cars, Car{
				Mileage: m,
				Price:   p,
			},
			)
		}
	}
}

var flagVisualizer *bool

// coordinates for visualizer
var (
	mileages []float64
	prices   []float64
)

func print() {
	sort.Slice(cars, func(i, j int) bool { return cars[i].Mileage < cars[j].Mileage })
	for _, v := range cars {
		mileages = append(mileages, float64(v.Mileage))
		prices = append(prices, float64(v.Price))
	}
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	flagVisualizer = flag.Bool("v", false, "Plot data in a graph")
	flag.Parse()
	getData()
	if *flagVisualizer {
		print()
	}
	// carsJSON, _ := json.Marshal(cars)
	// fmt.Println(string(carsJSON))
}
