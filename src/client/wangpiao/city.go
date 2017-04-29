package wangpiao

import (
	"client/common"
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"container/list"
	"strconv"
	//parse "client/wangpiao/parse"
	model "client/wangpiao/model"

)


func InsertCityList(l *list.List) {
	common.InsertList(l, "insert into city(id, name, Siteindex, Sitename)  values ", func(any interface{}) string {
		city := any.(*model.City)
		return "(0,'" + city.Name + "','" + strconv.Itoa(city.SiteId) + "','')"
	})
}

func GetCitysFromDB() *list.List {
	lcitys := list.New().Init()
	db, err := sql.Open("mysql", "root:1CUI@/piaofang")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	stmtOut, err := db.Prepare("SELECT id,name,Siteindex FROM city")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	rows, err := stmtOut.Query()
	for rows.Next() {
		city := new(model.City)
		err := rows.Scan(&city.Id, &city.Name, &city.SiteId)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		lcitys.PushBack(city)
	}
	return lcitys
}
