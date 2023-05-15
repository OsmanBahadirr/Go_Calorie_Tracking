package business

import (
	entity "calorieTracking/Entities"
)

func CalculateValues(nutr entity.Nutrient, amount int) entity.Nutrient {
	var nutrient entity.Nutrient

	nutrient.Cal = nutr.Cal * amount / 100
	nutrient.Carb = nutr.Carb * amount / 100
	nutrient.Fiber = nutr.Fiber * amount / 100
	nutrient.Fat = nutr.Fat * amount / 100
	nutrient.Pro = nutr.Pro * amount / 100

	return nutrient
}
