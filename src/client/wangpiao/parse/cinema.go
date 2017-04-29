package wangpiao

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"strconv"
	"container/list"
	model "client/wangpiao/model"
)

func GetCinemaByCitySiteId(cityid int) []model.SiteCinema {
	resp, err := http.Get("http://dataservices.wangpiao.com/Portal/ajaxcms/ajaxjson_cinemalist.aspx?CityID" + strconv.Itoa(cityid))
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	body, err := ioutil.ReadAll(resp.Body)
	body = body[1 : len(body)-1]
	var cinemas []model.SiteCinema
	err = json.Unmarshal(body, &cinemas)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Printf("%+v", cinemas)
	return cinemas
}

func GetCinemaByCity(city *model.City) *list.List{
	lcinemas := list.New().Init()
	cityid, _ := strconv.Atoi(city.SiteId)
	cinemas := GetCinemaByCitySiteId(cityid)
	for i := 0; i < len(cinemas); i++ {
		cinema := new(model.Cinema)
		cinema.Name = cinemas[i].ICinemaName
		cinema.SiteId = cinemas[i].CinemaIndex
		lcinemas.PushBack(cinema)
	}
	return lcinemas
}

func GetCinemaByCitys(lcitys *list.List) *list.List{
	lallcinemas := list.New().Init()
	for e := lcitys.Front(); e != nil; e = e.Next() {
		city := e.Value.(*model.City)
		lcinemas:=GetCinemaByCity(city)
		lallcinemas.PushBackList(lcinemas)
	}
	return lallcinemas
}

func GetCinemaByCitysMulti(lcitys *list.List) *list.List{
	lallcinemas := list.New().Init()
	for e := lcitys.Front(); e != nil; e = e.Next() {
		city := e.Value.(*model.City)
		lcinemas:=GetCinemaByCity(city)
		lallcinemas.PushBackList(lcinemas)
	}
	return lallcinemas
}
