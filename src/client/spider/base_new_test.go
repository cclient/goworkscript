package spider


import (
	"testing"
	"fmt"
)

func doTask_New(tasks []interface{}, showindex int, ch chan interface{}) {
	if showindex>=0{
		ch <- ("hello "+tasks[showindex].(string))
	}else {
		ch <- "hello -1"
	}
}

func TestParRunTaskLimitMaxConcurrence(t *testing.T) {
	urls:=make([]interface{},7)
	urls[0]="1"
	urls[1]="2"
	urls[2]="3"
	urls[3]="4"
	urls[4]="5"
	urls[5]="6"
	urls[6]="7"
	resulttasks:=ParRunTaskLimitMaxConcurrence(urls,5,doTask_New)
	fmt.Println("resulttasks end")
	if resulttasks.Len() != 0 {
		for e := resulttasks.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	}
}
