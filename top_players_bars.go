package main

import (
	"github.com/evilsocket/joe/models"
	"github.com/wcharczuk/go-chart"
)

func View(res *models.Results) models.Chart {
	ch := chart.BarChart{
		Title: "Top Players by APs",
		TitleStyle: chart.Style{
			Hidden: false,
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Bars: make([]chart.Value, res.NumRows),
	}

	for i, row := range res.Rows {
		ch.Bars[i] = chart.Value{
			Label: row["name"].(string),
			Value: float64(row["networks"].(int64)),
		}
	}

	return ch
}
