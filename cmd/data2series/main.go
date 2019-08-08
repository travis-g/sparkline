package main

import (
	"fmt"
)

func main() {
	input := map[string]interface{}{
		"data": []interface{}{"31", 100, "50m"},
	}

	series := Decode(input)
	fmt.Printf("%#v", *series)
}
