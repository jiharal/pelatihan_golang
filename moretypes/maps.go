package moretypes

import "fmt"

type VertexMaps struct {
	Lat, Long float64
}

var m map[string]VertexMaps

func Maps() {
	m = make(map[string]VertexMaps)
	m["Bell Labs"] = VertexMaps{
		40.68433, -74.39967,
	}
	m["BSB"] = VertexMaps{
		23.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["BSB"])
	fmt.Println(m)
}
