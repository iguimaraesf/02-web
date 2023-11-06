package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)
	delete(m, "c")
	fmt.Println(m["c"])
	_, exists := m["c"]
	fmt.Println(exists)

	var x = map[string]int{"x": 5, "y": 10}
	fmt.Println(x)

	if value, exists := m["c"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("não existe")
	}
}
