package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var totalId []int
var whitelistId []int

var N int = 1

func init() {

	var s string
	flag.StringVar(&s, "n", "800", "")
	flag.Parse()

	N, _ := strconv.ParseInt(s, 10, 64)

	for i := 0; i < int(N); i++ {
		tmp := rand.Intn(math.MaxInt32)
		totalId = append(totalId, tmp)
	}

	for i := 0; i < int(N); i++ {
		tmp := rand.Intn(math.MaxInt32)
		whitelistId = append(whitelistId, tmp)
	}
}

func test1() {
	now := time.Now()

	for _, v1 := range totalId {

		match := false
		for _, v2 := range whitelistId {
			if v1 == v2 {
				match = true
				break
			}
		}

		if !match {
			// do logic
		}
	}

	gap := time.Since(now)
	fmt.Println(gap)
}

func test2() {
	now := time.Now()

	m := make(map[int]int)
	for _, v2 := range whitelistId {
		m[v2] = 0
	}

	for _, v1 := range totalId {
		if _, ok := m[v1]; ok {
			continue
		}

		// do logic
	}

	gap := time.Since(now)
	fmt.Println(gap)
}

func main() {
	test1()
	test2()
}
