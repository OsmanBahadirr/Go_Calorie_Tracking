package business

import (
	entity "calorieTracking/Entities"
	"fmt"
	"net/http"

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

	doc.Find(".tr").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".td").Text()
		fmt.Println(title)

	})

	return nutrient
}
