package wangpiao

import (
	_ "github.com/go-sql-driver/mysql"
	"container/list"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	//"strconv"
	model "client/wangpiao/model"
	//"fmt"
)



func GetCity() *list.List {
	lcitys := list.New().Init()
	resp, err := http.Get("http://www.wangpiao.com/")
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		panic(err)
	}
	cityatags := doc.Find("body").Find("div[class=tab-content]").Find("li").Find("a")
	cityatags.Each(func(num int, s *goquery.Selection) {
		city := new(model.City)
		city.Name = s.Text()
		//fmt.Println(city.Name)
		//id, _ := strconv.Atoi(s.AttrOr("cityid", ""))
		city.SiteId = s.AttrOr("cityid", "")
		lcitys.PushBack(city)
	})
	return lcitys
}

