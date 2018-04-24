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
		Width: 860,
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

func print() {
	sort.Slice(cars, func(i, j int) bool { return cars[i].mileage < cars[j].mileage })
	for _, v := range cars {
		mileages = append(mileages, float64(v.mileage))
		prices = append(prices, float64(v.price))
	}

	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
