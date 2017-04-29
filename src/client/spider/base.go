package spider

import (
	"server/common/tool"
	ll "container/list"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"time"
)

func GetTasks() []TaskSimple {
	resultstrs, _ := tool.LRange(tool.GetClient(), "spider", 0, 100)
	tasks := make([]TaskSimple, len(resultstrs))
	for i, v := range resultstrs {
		var task TaskSimple
		value :=[]byte(v)
		_ = json.Unmarshal(value, &task)
		tasks[i] = task
	}
	return tasks
}

func RemoveTasks() {
	tool.LTrim(tool.GetClient(), "spider", 100, -1)
}

func doTask(tasks []TaskSimple, showindex int, ch chan TaskResult) {
	if showindex != -1 {
		task := tasks[showindex]
		greq, _ := http.NewRequest("GET", task.Url, nil)
		//部分站点需要加referer,或其他header
		//greq.Header.Add("Referer", "http://www.wangpiao.com")
		c := &http.Client{}
		resp, err := c.Do(greq)
		if err != nil {
			fmt.Println(resp, err)
			ch <- TaskResult{}
			return
			//panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		res := TaskResult{}
		res.Url = task.Url
		res.Status_code = resp.StatusCode
		res.Extra = task.Extra
		res.Content = tool.Base64EncodeToString(tool.DoZlibCompress(body))
		ch <- res
	}
	ch <- TaskResult{}
}


func DoTasks(pch chan bool,pinum int) {
	//一次从redis拿100条,分批请求,5组,每组20条
	tasks := GetTasks()
	if len(tasks) == 0 {
		go func() {
			a := time.Now().Unix()
			time.Sleep(10*time.Second)
			fmt.Println("start request ", time.Now().Unix() - a)
			pch <- true
		}()
	} else {
		itemcount := len(tasks)
		picount := itemcount / pinum //21 5
		yu := itemcount % pinum
		if yu != 0 {
			picount++ //6
		} else if picount == 1 {
			yu = pinum
		}
		chs := make([]chan TaskResult, pinum)
		for i := 0; i < picount; i++ {
			l := ll.New().Init()
			for j := 0; j < pinum; j++ {
				chs[j] = make(chan TaskResult)
				if (i != (picount - 1)) || j < yu {
					go doTask(tasks, pinum * i + j, chs[j])
				} else {
					go doTask(tasks, -1, chs[j])
				}
			}
			for _, ch := range chs {
				res := <-ch
				if res.Url != "" {
					fmt.Println(res.Url)
					b, err := json.Marshal(res)
					if err == nil {
						l.PushBack(string(b))
					}
				}
			}
			var insertstrs = make([]interface{}, l.Len())
			o := 0
			if l.Len() != 0 {
				for e := l.Front(); e != nil; e = e.Next() {
					insertstrs[o] = e.Value
					o++
				}
			}
			tool.RPushArr(tool.GetClient(), "spider_end", insertstrs)
		}
		RemoveTasks()
		pch <- true
	}
}

