package wangpiao

import (
	"client/common"
	"container/list"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	parse "client/wangpiao/parse"
	model "client/wangpiao/model"

)

func GetCinemasFromDB() *list.List {
	lcitys := list.New().Init()
	db, err := sql.Open("mysql", "root:1CUI@/piaofang")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	stmtOut, err := db.Prepare("SELECT id,name,Siteindex FROM cinema")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	rows, err := stmtOut.Query()
	for rows.Next() {
		cinema := new(model.Cinema)
		err := rows.Scan(&cinema.Id, &cinema.Name, &cinema.SiteId)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		//		fmt.Println(city.Name)
		lcitys.PushBack(cinema)
	}
	return lcitys
}

func InsertCinemaList(l *list.List) {
	common.InsertList(l, "insert into cinema(id, name, Siteindex, Siteindexs)  values ", func(any interface{}) string {
		cinema := any.(*model.Cinema)
		//		fmt.Println(cinema.Name)
		fmt.Println("(0,'" + cinema.Name + "'," + strconv.Itoa(cinema.SiteId) + ",'" + strconv.Itoa(cinema.SiteId) + "')")
		return "(0,'" + cinema.Name + "'," + strconv.Itoa(cinema.SiteId) + ",'" + strconv.Itoa(cinema.SiteId) + "')"
	})
}

func GetCinemaByCitys(lcitys *list.List) {
	for e := lcitys.Front(); e != nil; e = e.Next() {
		lcinemas := list.New().Init()
		city := e.Value.(*model.City)
		cinemas := parse.GetCinemaByCitySiteId(city.SiteId)
		for i := 0; i < len(cinemas); i++ {
			cinema := new(model.Cinema)
			cinema.Name = cinemas[i].ICinemaName
			cinema.SiteId = cinemas[i].CinemaIndex
			lcinemas.PushBack(cinema)
		}
		InsertCinemaList(lcinemas)
	}
}
