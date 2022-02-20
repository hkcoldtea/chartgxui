package main

import (
	"image"
	"math/rand"
	"os"

	chart "github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

const (
	colorMultiplier = 256
	fontFile = "/usr/share/fonts/truetype/arphic/uming.ttc"
)

func chartPNG(fname string, bWritePNG bool) (image.Image, error) {
	var zhfont = getZHFont(fontFile)
	xv, yv := GetData(fname)

	priceSeries := chart.TimeSeries{
		Name: "恒生指數",
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: xv,
		YValues: yv,
	}

	smaSeries := chart.SMASeries{
		Name: "恒生指數 - SMA",
		Style: chart.Style{
			StrokeColor:     drawing.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: priceSeries,
	}

	bbSeries := &chart.BollingerBandsSeries{
		Name: "恒生指數 - Bol. Bands",
		Style: chart.Style{
			StrokeColor: drawing.ColorFromHex("efefef"),
			FillColor: drawing.Color{
				R: uint8(rand.Intn(colorMultiplier)),
				G: uint8(rand.Intn(colorMultiplier)),
				B: uint8(rand.Intn(colorMultiplier)),
				A: uint8(colorMultiplier - 1), // 透明度
			}.WithAlpha(64),
		},
		InnerSeries: priceSeries,
	}

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top: 20,
				Right: 60,
			},
		},
		XAxis: chart.XAxis{
			TickPosition: chart.TickPositionBetweenTicks,
			TickStyle: chart.Style{
				TextRotationDegrees: 45.0,
			},
		},
		YAxis: chart.YAxis{
			Name: "恒生指數",
			NameStyle: chart.Style{
				TextRotationDegrees: 90.0,
				FontColor: chart.ColorCyan,
				FontSize:  14,
				Font:      zhfont,
			},
		},
		Series: []chart.Series{
			bbSeries,
			priceSeries,
			smaSeries,
		},
	}

	style := chart.Style{
		FontColor:    drawing.ColorBlack,
		FontSize:     12,
		Font:         zhfont,
	}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph, style),
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
