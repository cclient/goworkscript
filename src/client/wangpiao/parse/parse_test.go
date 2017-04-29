package wangpiao

import (
	"testing"
	"fmt"
)

func Test_GetCity(t *testing.T) {
	citys:=GetCity()
	fmt.Println(citys.Front().Value)
}

func Test_GetCinemaByCitySiteId(t *testing.T) {
	cinemas:=GetCinemaByCitySiteId(1)
	//fmt.Printf("%+v", cinemas)
	fmt.Println(cinemas)
}

func Test_GetShowTimeSingleCinema(t *testing.T) {
	showtimes:=GetShowTimeSingleCinema(6069,"2017-03-27")
	//fmt.Printf("%+v", cinemas)
	fmt.Println(showtimes)
}
