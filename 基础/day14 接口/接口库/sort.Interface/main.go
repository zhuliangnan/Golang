package main

import (
	"fmt"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type Tracks []Track

func ParseDurationTime(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return ParseDurationTime("0s")
	} else {
		return d
	}
}

// {{{ implementation of Sort interface
// Len is the number of elements in the collection.
func (x Tracks) Len() int {
	return len(x)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (x Tracks) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

// Swap swaps the elements with indexes i and j.
func (x Tracks) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// end implementation of Sort interface }}}

var tracks = Tracks{
	{Title: "C#", Artist: "Delu", Album: "Reading", Year: 2017, Length: ParseDurationTime("3m38s")},
	{Title: "Go", Artist: "Anderson", Album: "Reading", Year: 2018, Length: ParseDurationTime("3m38s")},

	{Title: "Java Bible", Artist: "Js", Album: "Reading", Year: 2016, Length: ParseDurationTime("3m38s")}}

//main function
func main() {

	sort.Sort(tracks)

	for key, value := range tracks {
		fmt.Printf("%v:%v \n", key, value)
	}
}
