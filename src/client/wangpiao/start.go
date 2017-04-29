package wangpiao

import (
	_ "github.com/go-sql-driver/mysql"
	parse "client/wangpiao/parse"
	model "client/wangpiao/model"
)

func start() {
	//得到所有城市
	listCitys:=parse.GetCity()
	//go InsertCityList(listCitys)
	//再得到所有城市的影院
	listCinemas:=parse.GetCinemaByCitys(listCitys.Front().Value.(*model.City))
	go InsertCinemaList(listCinemas)
	//再得影院的场次//todo 场次官网有更新,不能用了
	//listShowTimes:=parse.GetShowTimeSingleCinema(listCinemas.Front().Value.(*model.Cinema),"")
	//go InsertShowTimeList(listShowTimes)
	//定时更新场次

}