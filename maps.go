package main

import "fmt"

type vertex struct {
	Lat, Long float64
}

var m map[string]vertex

func maps() {
	m = make(map[string]vertex)
	m["Bell Labs"] = vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Printf("map element: %#v\n", m["Bell Labs"])
	fmt.Printf("map: %#v\n", m)

	m = map[string]vertex{
		"Bell Labs": vertex{
			40.68433, -74.39967,
		},
		"Google": vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Printf("literal map: %v\n", m)

	m = map[string]vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Printf("literal map with no type name: %v\n", m)
}
