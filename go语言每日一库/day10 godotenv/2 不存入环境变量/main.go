package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	var myEnv map[string]string
	//默认读取根目录下.env
	//return []string{".env"}
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])
}
