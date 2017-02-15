package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

const (
	//TABELOG_URL = "https://tabelog.com/"
	LOCALHOST = "http://8pockets.com/"
)

func main() {
	//
	// * ここで県を指定すると、その県の駅を全取得 *
	//
	area := "tokyo"

	prefPage, err := goquery.NewDocument(LOCALHOST + area + ".html")

	if err != nil {
		fmt.Print("url scarapping failed")
	}

	//Parse HTML By goquery module
	prefPage.Find(".modal__inner .link-list li.link a").Each(func(_ int, train *goquery.Selection) {
		//train.Attr("href")の例：https://tabelog.com/tokyo/R9/rstLst/
		trainUrl, _ := train.Attr("href")
		trainHtml, _ := goquery.NewDocument(trainUrl)

		trainHtml.Find("li.station label a").Each(func(_ int, station *goquery.Selection) {
			stationName := station.Text()
			stationUrl, _ := station.Attr("href")

			re := regexp.MustCompile("([a-z]*/A[0-9]*/A[0-9]*/R[0-9]*)")
			stationId := re.FindStringSubmatch(stationUrl)

			fmt.Println(stationName)
			fmt.Println(stationId[1])
		})
	})
}
