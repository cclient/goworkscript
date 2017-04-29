package pdf

import (
	"time"
	"fmt"
)

func Start() {
	ch := make(chan bool)
	for {
		go func() {
			a := time.Now().Unix()
			time.Sleep(5*time.Minute)
			fmt.Println("continue new task", time.Now().Unix() - a)
			ch <- true
		}()
		//阻塞,直到定时触发返回
		<-ch
		//继续拿数
		DoWork(1,40)
	}
}