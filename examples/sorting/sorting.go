package main

import (
	"encoding/json"
	"fmt"

	"github.com/nexusriot/gopou/sorting"
)

type Album struct {
	Artist string
	Name   string
	Year   int
}

func main() {
	albums := []*Album{
		{"Nightwish", "Oceanborn", 1998},
		{"Nightwish", "Wishmaster", 2000},
		{"Turmion Kätilöt", "Technodiktator", 2013},
		{"Turmion Kätilöt", "Diskovibrator", 2015},
		{"Turmion Kätilöt", "Global Warning", 2020},
		{"Amaranthe", "Manifest", 2020},
	}

	criteria := []*sorting.Criteria{
		{"Artist", true},
		{"Year", true},
	}
	sorted, _ := sorting.Sorted(criteria, albums)

	marshalled, err := json.Marshal(sorted)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshalled))
}
