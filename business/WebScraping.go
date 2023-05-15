package business

import (
	entity "calorieTracking/Entities"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetUrl(n string) string {
	URL := "https://www.diyetkolik.com/kac-kalori/" + n

	return URL
}

func GetValues(URL string) entity.Nutrient {

	var nutrient entity.Nutrient

	res, err := http.Get(URL)

	if err != nil || res.StatusCode != 200 {
		fmt.Println(err)
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {

	})

	titleCal := doc.Find("span.kkBigNumber.nut_kcal_count2").Each(func(i int, s *goquery.Selection) {

	})

	nutrientCarb := strings.Split(title.Find(".lbl_carb100").Text(), ".")[0]
	nutrientFiber := strings.Split(title.Find(".lbl_lif100").Text(), ".")[0]
	nutrientPro := strings.Split(title.Find(".lbl_prot100").Text(), ".")[0]
	nutrientFat := strings.Split(title.Find(".lbl_fat100").Text(), ".")[0]

	nutrient.Cal, _ = strconv.Atoi(titleCal.Text())
	nutrient.Carb, _ = strconv.Atoi(nutrientCarb)
	nutrient.Fiber, _ = strconv.Atoi(nutrientFiber)
	nutrient.Pro, _ = strconv.Atoi(nutrientPro)
	nutrient.Fat, _ = strconv.Atoi(nutrientFat)

	return nutrient
}
