package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"os"
	"path/filepath"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/themes/light"
)

var (
	BUILD              string
	fname              string
	bWritePNG          bool
	DefaultScaleFactor float64
	FlagTheme          string
)

func appMain(driver gxui.Driver) {
	source, err := chartPNG(fname, bWritePNG)
	if err != nil {
		return
	}
	debug("fname=%s\n", fname)

	theme := CreateTheme(driver)
	img := theme.CreateImage()

	mx := source.Bounds().Max
	wTitle := fmt.Sprintf("CHART: [%s]", fname)
	window := theme.CreateWindow(mx.X, mx.Y, wTitle)
	window.SetScale(float32(DefaultScaleFactor))

	// Copy the image to a RGBA format before handing to a gxui.Texture
	rgba := image.NewRGBA(source.Bounds())
	draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
	texture := driver.CreateTexture(rgba, 1)
	img.SetTexture(texture)
	img.OnClick(func(gxui.MouseEvent) { window.Close() })

	window.AddChild(img)
	window.OnClose(driver.Terminate)
}

// CreateTheme creates and returns the theme specified on the command line.
// The default theme is dark.
func CreateTheme(driver gxui.Driver) gxui.Theme {
	if FlagTheme == "light" {
		return light.CreateTheme(driver)
	}
	return dark.CreateTheme(driver)
}

func main() {
	flag.BoolVar(&bWritePNG, "w", false, "write output.png")
	flag.StringVar(&fname, "d", "", "filename of json")
	flag.StringVar(&FlagTheme, "theme", "dark", "Theme to use {dark|light}.")
	flag.Float64Var(&DefaultScaleFactor, "scaling", 1.0, "Adjusts the scaling of UI rendering")
	flag.Usage = func() {
		w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily
		if len(BUILD) > 0 {
			fmt.Fprintf(w, "Build: %s\n", BUILD)
		}
		fmt.Fprintf(w, "Usage of %s: [datafile]\n\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		fmt.Fprintf(w, "\n")
	}
	flag.Parse()
	if bWritePNG {
		chartPNG(fname, bWritePNG)
		return
	}
	if len(fname) == 0 {
		if flag.NArg() == 0 {
			flag.Usage()
			return
		}
		fname = flag.Args()[0]
	}
	gl.StartDriver(appMain)
}
