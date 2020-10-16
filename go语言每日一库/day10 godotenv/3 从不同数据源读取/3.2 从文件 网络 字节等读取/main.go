package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

/*
只要实现了io.Reader接口，就能作为数据源。可以从文件（os.File），网络（net.Conn），bytes.Buffer等多种来源读取：
*/
func main() {
	file, _ := os.OpenFile(".env", os.O_RDONLY, 0666)
	myEnv, err := godotenv.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])

	buf := &bytes.Buffer{}
	buf.WriteString("name: awesome web @buffer")
	buf.Write([]byte{'\n'})
	buf.WriteString("version: 0.0.1")
	myEnv, err = godotenv.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])
}
