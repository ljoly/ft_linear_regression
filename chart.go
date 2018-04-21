package main

import (
	"net/http"

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
				XValues: mileages,
				YValues: prices,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}
