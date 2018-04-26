package main

import (
	"log"
	"net/http"
	"sort"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Mileage (km)",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			AxisType:  chart.YAxisSecondary,
			Name:      "Price (€)",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: 10000,
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:  30,
				Left: 30,
			},
		},
		Width: 1024,
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeWidth: chart.Disabled,
					DotWidth:    5,
				},
				XValues: mileages,
				YValues: prices,
			},
			chart.ContinuousSeries{
				XValues: linearRegXs,
				YValues: linearRegYs,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

var (
	aX          float64
	aY          float64
	bX          float64
	bY          float64
	linearRegXs []float64
	linearRegYs []float64
)

func print() {
	sort.Slice(cars, func(i, j int) bool { return cars[i].mileage < cars[j].mileage })
	for _, v := range cars {
		mileages = append(mileages, float64(v.mileage))
		prices = append(prices, float64(v.price))
	}

	aX = 22899
	aY = theta1*aX + theta0
	linearRegXs = append(linearRegXs, aX)
	linearRegYs = append(linearRegYs, aY)
	bX = 240000
	bY = theta1*bX + theta0
	linearRegXs = append(linearRegXs, bX)
	linearRegYs = append(linearRegYs, bY)

	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
