package wangpiao

import (
	"container/list"
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"strings"
	//"errcode"
	"client/common"
	model "client/wangpiao/model"

)

func GeMovie() *list.List {
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

func InserMovieList(l *list.List) {
	common.InsertList(l, "insert into movie values ", func(any interface{}) string {
		movie := any.(*model.Movie)
		return "(0,'" + movie.Name + "','" + movie.Jpg + "')"
	})
}

func getHasNames(rows *sql.Rows) []string {
	lhasnames := list.New().Init()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		lhasnames.PushBack(name)
	}
	return common.StringListToArray(lhasnames)
}

func getNoHasNames(oldl *list.List, hasnames []string) []string {
	var linserMovienames = list.New().Init()
	for e := oldl.Front(); e != nil; e = e.Next() {
		movie := e.Value.(*model.Movie)
		isinsert := true
		for j := 0; j < len(hasnames); j++ {
			if movie.Name == hasnames[j] {
				isinsert = false
				break
			}
		}
		if isinsert == true {
			linserMovienames.PushBack(movie.Name)
		}
	}
	return common.StringListToArray(linserMovienames)
}

func geMovieByNames(oldl *list.List, names []string) *list.List {
	var linserMovie = list.New().Init()
	for i := 0; i < len(names); i++ {
		movie := geMovieByName(oldl, names[i])
		if movie != nil {
			linserMovie.PushBack(movie)
		}
	}
	fmt.Println(linserMovie.Len())
	return linserMovie
}
func geMovieByName(oldl *list.List, name string) *model.Movie {
	for e := oldl.Front(); e != nil; e = e.Next() {
		movie := e.Value.(*model.Movie)
		if movie.Name == name {
			return movie
		}
	}
	return nil
}

func SaveMovie(l *list.List) {
	db, err := sql.Open("mysql", "root:1CUI@/piaofang")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	names := make([]string, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*model.Movie).Name)
		names[i] = "\"" + e.Value.(*model.Movie).Name + "\""
		i++
	}
	namestrs := strings.Join(names, ",")
	stmtOut, err := db.Prepare("SELECT name FROM movie WHERE name in (" + namestrs + ")")
	fmt.Print("SELECT name FROM movie WHERE name in (" + namestrs + ")")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	rows, err := stmtOut.Query()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	hasnames := getHasNames(rows)
	fmt.Println("has count" + strconv.Itoa(len(hasnames)))
	nohasnames := getNoHasNames(l, hasnames)
	fmt.Println("no has count" + strconv.Itoa(len(nohasnames)))
	inserMovies := geMovieByNames(l, nohasnames)
	fmt.Printf(strconv.Itoa(inserMovies.Len()))
	InserMovieList(inserMovies)
}
