package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	Color string
}

func main() {
	var car Car
	data := []byte(`{"Name":"Corcel","Year":1970,"Color":"Green"}`)
	json.Unmarshal(data, &car)
	fmt.Println(car)
}
