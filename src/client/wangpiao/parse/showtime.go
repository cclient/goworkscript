package wangpiao

import (
	"client/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	//"container/list"
	model "client/wangpiao/model"

)


func GetShowTimeSingleCinema(cinemaid int, datestr string) []model.SiteShowTime {
	//http://dataservices1.wangpiao.com/API.aspx?Target=Base_FilmShow&Param=CinemaID=6069&Date=2017-03-27
	resp, err := http.Get("http://dataservices1.wangpiao.com/API.aspx?Target=Base_FilmShow&Param=CinemaID=" + strconv.Itoa(cinemaid) + "&Date=" + datestr)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	body, err := ioutil.ReadAll(resp.Body)
	body = body[1 : len(body)-1]
	var sto model.SiteShowTimeData
	err = json.Unmarshal(body, &sto)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", sto.Data)
	return sto.Data
}

//func GetShowTimeCinemas(lcinemas *list.List, datestr string) *list.List{
//	lallshowtimes := list.New().Init()
//	for e := lcinemas.Front(); e != nil; e = e.Next() {
//		cinema := e.Value.(*model.Cinema)
//		showtimes:= GetShowTimeSingleCinema(cinema.SiteId,"")
//		//todo 遍历加入list
//		lallshowtimes.PushBackList(showtimes)
//	}
//	return lallshowtimes
//}

func GetSingleShowTimeCurrentSale(lmap map[int]int, showindex int, ch chan bool) {
	if showindex != 0 {
		greq, _ := http.NewRequest("GET", "http://dataservices.wangpiao.com/Data.aspx?getpageurl=Http%3A//dataservices.wangpiao.com/Portal/ajaxcms/ajax_SeatGrid.aspx&getpageparam=SeqNo%3D"+strconv.Itoa(showindex)+"&format=json&_="+ common.GetNowTimeTsString(), nil)
		greq.Header.Add("Referer", "http://www.wangpiao.com")
		c := &http.Client{}
		resp, err := c.Do(greq)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		s := string(body[0:len(body)])
		//strings.Count(s,"#FF0000")
		saledcount := strings.Count(s, "#F")
		lmap[showindex] = saledcount
	}
	ch <- true
}
