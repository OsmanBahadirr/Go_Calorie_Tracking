package main

import (
	x "calorieTracking/business"
	"fmt"
)

func main() {

	fmt.Println(x.GetValues(x.GetUrl("muz")))
}
