package webScrapping

import (
	"fmt"
	"github.com/gocolly/colly"
)

type product struct {
	Name    string `json:"name"`
	ImgUrl  string `json:"url"`
	Sku     string `json:"sku"`
	Barcode string `json:"barcode"`
	Price   string `json:"price"`
}

func Scraper() {
	c := colly.NewCollector(
		colly.AllowedDomains("flo.com.tr"))
	c.OnHTML("div[class=js-product-vertical product-list__item listing__col-product]", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText("div.class=data-product-sku"))
	})

	c.Visit("https://www.flo.com.tr/ayakkabi?cinsiyet=erkek")
}
