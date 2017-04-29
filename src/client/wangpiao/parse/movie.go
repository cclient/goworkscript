package wangpiao

import (
	"container/list"
	//"database/sql"
	//"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	model "client/wangpiao/model"

	//"strconv"
	//"strings"
	//"errcode"
)


func GetMovie() *list.List {
	var lmovies = list.New().Init()
	resp, err := http.Get("http://www.wangpiao.com/Movie/movies/")
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		panic(err)
	}
	htmlbody := doc.Find("body")
	maindiv := htmlbody.Find("div[id=filmon]")
	divs := maindiv.Find("div[class=mt20\\ pb20\\ bodbCCC]")
	divs.Each(func(num int, s *goquery.Selection) {
		img := s.Find("div[class=movie_bg\\ movie_pic\\ pr]").Children().First().Children()
		movie := new(model.Movie)
		movie.Name = img.AttrOr("title", "")
		movie.Jpg = img.AttrOr("src", "")
		lmovies.PushBack(movie)
	})
	return lmovies
}
