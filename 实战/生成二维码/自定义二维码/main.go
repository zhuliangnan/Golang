package main

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
)

func main() {
	qr, err := qrcode.New("http://www.codesuger.com/", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{50, 205, 50, 255}
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "./zidingy.png")
	}

}
