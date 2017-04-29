package wangpiao

import (
	"container/list"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
	"client/common"
	model "client/wangpiao/model"
	parse "client/wangpiao/parse"
)

func getShowTimeSiteIndexFromDB() []int {
	lshowindex := list.New().Init()
	db, err := sql.Open("mysql", "root:1CUI@/piaofang")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	stmtOut, err := db.Prepare("SELECT Siteshowindex FROM showtime")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	rows, err := stmtOut.Query()
	for rows.Next() {
		var index int
		err := rows.Scan(&index)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		//		fmt.Println(city.Name)
		lshowindex.PushBack(index)
	}
	return common.IntListToArray(lshowindex)
}


func InsertShowTimeList(l *list.List) {
	common.InsertList(l, "insert into showtime (id,Site,Sitename,Siteshowindex,Sitecinemaindex,Sitemovieindex,Sitemoviename,Sitehallid,Sitesaleendtime,price,seatcount,Sitecityindex)  values ", func(any interface{}) string {
		showtime := any.(*model.ShowTime)
		//		fmt.Println("(0,0,'网票'," + strconv.Itoa(showtime.SiteShowIndex) + "," + strconv.Itoa(showtime.SiteCinemaIndex) + "," + strconv.Itoa(showtime.SiteMovieIndex) + ",'" + showtime.SiteMovieName + "'," + strconv.Itoa(showtime.SiteHallID) + "," + strconv.FormatInt(showtime.SiteSaleEndTime, 10) + "," + strconv.Itoa(showtime.Price) + "," + strconv.Itoa(showtime.SeatCount) + "," + strconv.Itoa(showtime.SiteCityIndex) + ")")
		return "(0,0,'网票'," + strconv.Itoa(showtime.ShowSiteID) + "," + strconv.Itoa(showtime.ShowSiteID) + "," + strconv.Itoa(showtime.SiteMovieIndex) + ",'" + showtime.SiteMovieName + "'," + strconv.Itoa(showtime.SiteHallID) + "," + strconv.FormatInt(showtime.SiteSaleEndTime, 10) + "," + strconv.Itoa(showtime.Price) + "," + strconv.Itoa(showtime.SeatCount) + "," + strconv.Itoa(showtime.SiteCityIndex) + ")"
	})
}

func GetShowTime(lcinema *list.List) {
	for e := lcinema.Front(); e != nil; e = e.Next() {
		lshowtimes := list.New().Init()
		cinema := e.Value.(*model.Cinema)
		showtimes := parse.GetShowTimeSingleCinema(cinema.SiteId, "2015-10-15")
		for i := 0; i < len(showtimes); i++ {
			jshowtime := showtimes[i]
			showtime := new(model.ShowTime)
			showtime.SeatCount = jshowtime.SeatCount
			showtime.SiteId = 0
			showtime.CinemaSiteID = jshowtime.CinemaID
			showtime.CitySiteId = jshowtime.CityID
			showtime.HallSiteID = jshowtime.HallID
			showtime.MovieSiteID = jshowtime.FilmID
			showtime.SiteMovieName = jshowtime.FilmName
			showtime.SiteName = "wangpiao"
			//showtime.SiteSaleEndTimeS = jshowtime.SaleEndTime
			ts, _ := time.Parse("2006-01-02 15:04:05", jshowtime.SaleEndTime)
			showtime.SiteSaleEndTimeS = ts
			showtime.SiteSaleEndTime = ts.Unix()
			showtime.ShowSiteID= jshowtime.ShowIndex
			lshowtimes.PushBack(showtime)
		}
		InsertShowTimeList(lshowtimes)
	}
}




//func SaveTodayShowTime() {
//	lcinemas := getCinemasFromDB()
//	GetShowTime(lcinemas)
//}


func UpdateAllShowTimePeople() {
	pinum := 10
	arr := getShowTimeSiteIndexFromDB()
	itemcount := len(arr)
	picount := itemcount / pinum //21 5
	yu := itemcount % pinum
	if yu != 0 {
		picount++ //6
	}
	chs := make([]chan bool, pinum)
	for i := 0; i < picount; i++ {
		pimap := make(map[int]int)
		for j := 0; j < pinum; j++ {
			chs[j] = make(chan bool)
			if (i != (picount - 1)) || j < yu {
				go parse.GetSingleShowTimeCurrentSale(pimap, arr[pinum*i+j], chs[j])
			} else {
				go parse.GetSingleShowTimeCurrentSale(pimap, 0, chs[j])
			}
		}
		for _, ch := range chs {
			<-ch
		}
		UpdateShowTimeSaledByShowIndex(pimap)
	}
	fmt.Println("all show time  update finished")
}

func UpdateShowTimeSaledByShowIndex(mapinfo map[int]int) int64 {
	db, err := sql.Open("mysql", "root:1CUI@/piaofang")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	//UPDATE categories
	//    SET display_order = CASE id
	//        WHEN 1 THEN 3
	//        WHEN 2 THEN 4
	//        WHEN 3 THEN 5
	//    END
	//WHERE id IN (1,2,3)
	exestrpre := "UPDATE piaofang.showtime set salecount= CASE Siteshowindex\n"
	exestrsub := ""
	lwhere := list.New().Init()
	for k, v := range mapinfo {
		if k != 0 && v != 0 {
			exestrsub += "WHEN " + strconv.Itoa(k) + " THEN " + strconv.Itoa(v) + "\n"
			lwhere.PushBack(strconv.Itoa(k))
		}
	}
	res, err := db.Exec(exestrpre + exestrsub + "\n End Where Siteshowindex In (" + strings.Join(model.StringListToArray(lwhere), ",") + ")")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	count, _ := res.RowsAffected()
	return count
}
