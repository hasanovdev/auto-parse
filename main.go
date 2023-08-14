package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/google/uuid"
)

type item struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Year      string `json:"year"`
	BodyStyle string `json:"body_style"`
	ImgUrl    string `json:"img_url"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("autopremiumgroup.ru"),
	)

	var items []item

	c.OnHTML("div.auto-list__results a.auto-list__results__i", func(h *colly.HTMLElement) {
		item := item{
			Id:        uuid.NewString(),
			Title:     h.ChildText("span.auto-list__results__i__title"),
			Year:      h.ChildText("span.auto-list__results__i__year"),
			BodyStyle: h.ChildText("span.auto-list__results__i__body_style"),
			ImgUrl:    "https://autopremiumgroup.ru" + h.ChildAttr("img", "src"),
		}

		items = append(items, item)
	})

	c.Visit("https://autopremiumgroup.ru/katalog-avtomobilej/")

	content, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("automobiles.json", content, 0644)
}
