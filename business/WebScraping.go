package business

import (
	entity "calorieTracking/Entities"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetUrl(n string) string {
	URL := "https://www.diyetkolik.com/kac-kalori/" + n

	return URL
}

func GetValues(URL string) entity.NutrientError {

	var nutrient entity.NutrientError

	res, err := http.Get(URL)

	if err != nil {
		fmt.Println(err)
		nutrient.Err = err
		return nutrient
	}

	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		nutrient.Err = err
		return nutrient
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		fmt.Println(err)
		nutrient.Err = err
		return nutrient
	}

	title := doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {

	})

	nutrientCarb := strings.Split(title.Find(".lbl_carb100").Text(), ".")[0]
	nutrientFiber := strings.Split(title.Find(".lbl_lif100").Text(), ".")[0]
	nutrientPro := strings.Split(title.Find(".lbl_prot100").Text(), ".")[0]
	nutrientFat := strings.Split(title.Find(".lbl_fat100").Text(), ".")[0]

	nutrient.Nutrient.Carb, _ = strconv.Atoi(nutrientCarb)
	nutrient.Nutrient.Fiber, _ = strconv.Atoi(nutrientFiber)
	nutrient.Nutrient.Pro, _ = strconv.Atoi(nutrientPro)
	nutrient.Nutrient.Fat, _ = strconv.Atoi(nutrientFat)

	return nutrient
}
