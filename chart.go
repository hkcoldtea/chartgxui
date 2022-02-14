package main

import (
	"image"
	"os"

	chart "github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func chartPNG(fname string, bWritePNG bool) (image.Image, error) {
	xv, yv := GetData(fname)

	priceSeries := chart.TimeSeries{
		Name: "SPY",
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: xv,
		YValues: yv,
	}

	smaSeries := chart.SMASeries{
		Name: "SPY - SMA",
		Style: chart.Style{
			StrokeColor:     drawing.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: priceSeries,
	}

	bbSeries := &chart.BollingerBandsSeries{
		Name: "SPY - Bol. Bands",
		Style: chart.Style{
			StrokeColor: drawing.ColorFromHex("efefef"),
			FillColor:   drawing.ColorFromHex("efefef").WithAlpha(64),
		},
		InnerSeries: priceSeries,
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			TickPosition: chart.TickPositionBetweenTicks,
			TickStyle: chart.Style{
				TextRotationDegrees: 45.0,
			},
		},
/*
		YAxis: chart.YAxis{
		},
*/
		Series: []chart.Series{
			bbSeries,
			priceSeries,
			smaSeries,
		},
	}

	if bWritePNG {
		f, _ := os.Create("output.png")
		defer f.Close()
		graph.Render(chart.PNG, f)
		return nil, nil
	}

	collector := &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	return collector.Image()
}
