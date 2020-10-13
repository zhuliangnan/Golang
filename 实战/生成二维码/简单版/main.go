package main

import "github.com/skip2/go-qrcode"

func main() {

	qrcode.WriteFile("http://www.codesuger.com/", qrcode.Medium, 256, "./blog_qrcode.png")

}
