package wangpiao

import "client/model"

type City struct {
	model.BaseCity
}


type Cinema struct {
	model.BaseCinema
	CinemaIndex int
	ICinemaName string
}
type SiteCinema struct {
	CinemaIndex int
	ICinemaName string
}


type ShowTime struct {
	model.BaseShowTime
}


type Movie struct {
	model.BaseMovie
	tid string
}


//用来包装json对象,转换用
type SiteShowTime struct {
	ShowIndex int
	CinemaID  int
	HallID    int
	FilmID    int
	FilmName  string
	//LG: 原版,
	//ShowTime: 2015-10-15 22:00:00,
	SaleEndTime string
	//Status: 1,
	//UPrice: 45,
	//VPrice: 45,
	CityID int
	//UWPrice: 50,
	//SPSite: 1|2|5,
	//SPPrice: 5|0|0,
	//HallName: 12号厅,
	//CPrice: 50,
	//IsImax: false,
	//Dimensional: 3D,
	SeatCount int
}

//用来包装json对象,转换用
type SiteShowTimeData struct {
	ErrNo int
	Sign  string
	Msg   string
	Data  []SiteShowTime
}