package main

import (
	"io/ioutil"
	"log"
	"github.com/golang/freetype/truetype"
)

func getZHFont(fontFile string) *truetype.Font {
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return nil
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return nil
	}
	return font
}
