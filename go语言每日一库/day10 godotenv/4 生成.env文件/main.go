package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	buf := &bytes.Buffer{}
	buf.WriteString("name = awesome web")
	buf.WriteByte('\n')
	buf.WriteString("version = 0.0.1")

	//生成map
	env, err := godotenv.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}

	//返回一个字符串
	content, err := godotenv.Marshal(env)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)

	/*	err = godotenv.Write(env, "go语言每日一库/day10 godotenv/4 生成.env文件/.env")
		if err != nil {
			log.Fatal(err)
		}*/
}
