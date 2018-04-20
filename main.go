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

	chart "github.com/wcharczuk/go-chart"
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
		cars = append(cars, Car{
			Mileage: m,
			Price:   p,
		},
		)
	}
}

func drawChart(res http.ResponseWriter, req *http.Request) {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func drawChartWide(res http.ResponseWriter, req *http.Request) {
	graph := chart.Chart{
		Width: 1920, //this overrides the default.
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	flagVisualizer := flag.Bool("v", false, "Plot data in a graph")
	flag.Parse()
	getData()
	if *flagVisualizer {
		http.HandleFunc("/", drawChart)
		http.HandleFunc("/wide", drawChartWide)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
	// carsJSON, _ := json.Marshal(cars)
	// fmt.Println(string(carsJSON))
}
