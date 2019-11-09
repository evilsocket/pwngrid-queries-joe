package main

import (
	"github.com/evilsocket/islazy/log"
	"github.com/evilsocket/joe/models"
	"github.com/wcharczuk/go-chart"
	"time"
)

func View(res *models.Results) models.Chart {
	ts := chart.TimeSeries{
		XValues: make([]time.Time, 0),
		YValues: make([]float64, 0),
	}

	layout := "2006-01-02"
	for _, row := range res.Rows {
		if row["day"] != nil {
			day := row["day"].(string)
			count := float64(row["sent"].(int64))

			if t, err := time.Parse(layout, day); err == nil {
				ts.XValues = append(ts.XValues, t)
				ts.YValues = append(ts.YValues, count)
			} else {
				log.Error("can't parse date '%s' as date layout %s", day, layout)
			}
		}
	}

	return chart.Chart{
		Title: "Messages per Day",
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
