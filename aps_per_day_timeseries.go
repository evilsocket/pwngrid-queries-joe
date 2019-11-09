package main

import (
	"github.com/evilsocket/islazy/log"
	"github.com/evilsocket/joe/models"
	"github.com/wcharczuk/go-chart"
	"time"
)

func View(res *models.Results) models.Chart {
	ts := chart.TimeSeries{
		XValues: make([]time.Time, res.NumRows),
		YValues: make([]float64, res.NumRows),
	}

	layout := "2006-01-02"
	for i, row := range res.Rows {
		day := row["day"].(string)
		count := float64(row["reported"].(int64))

		if t, err := time.Parse(layout, day); err == nil {
			ts.XValues[i] = t
			ts.YValues[i] = count
		} else {
			log.Error("can't parse date '%s' as date layout %s", day, layout)
		}
	}

	return chart.Chart{
		Title: "Reported APs per Day",
		TitleStyle: chart.Style{
			Hidden: false,
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeDateValueFormatter,
		},
		Series: []chart.Series{ts},
	}
}
