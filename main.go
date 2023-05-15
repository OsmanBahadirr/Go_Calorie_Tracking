package main

import (
	business "calorieTracking/business"
	"fmt"
)

func main() {

	x := business.CalculateValues(business.GetValues(business.GetUrl("muz")), 30)
	fmt.Println(x)
}
