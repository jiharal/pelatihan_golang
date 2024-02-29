package moretypes

import "fmt"

func MapLiteralsContinued() {
	var m = map[string]VertexMaps{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
		"BSB":       VertexMaps{10.10212, -12.90230},
	}
	fmt.Println(m)
}
