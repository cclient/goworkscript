package pdf

import (
	"container/list"
	_ "github.com/go-sql-driver/mysql"

	"os"
	"net/http"
	"io/ioutil"
)


//func GetRedisClient() *redis.Client{
//	client := redis.NewClient(&redis.Options{
//		Addr:         "10.10.3.111:6379",
//		DialTimeout:  10 * time.Second,
//		ReadTimeout:  30 * time.Second,
//		WriteTimeout: 30 * time.Second,
//		PoolSize:     50,
//		PoolTimeout:  30 * time.Second,
//		DB:redisDb,
//	})
//	return client
//}

func StringListToArray(l *list.List) []string {
	strs := make([]string, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		strs[i] = e.Value.(string)
		i++
	}
	return strs
}

func IntListToArray(l *list.List) []int {
	strs := make([]int, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		strs[i] = e.Value.(int)
		i++
	}
	return strs
}


func StringArrToInterfaceArr(values []string) []interface{}{
	vs:=make([]interface{},len(values))
	for i:=0;i<len(values);i++{
		vs[i]=values[i]
	}
	return vs
}

func WebPageContainPdf(engine ISearchEngine,resps []*http.Response) *list.List{
	l := list.New().Init()
	for i:=0;i<len(resps);i++{
		//urls:=BuildMultiPageUrls(engine,1,1)
		//resps:=RequestUrls(urls)
		pdfinfos:=engine.ParseContainPdfInfo(resps[i])
		for i := 0; i < len(pdfinfos); i++ {
			l.PushBack(pdfinfos[i])
			//todo
			//return l
		}
	}
	return l
}

func downloadPDF(pdfinfo PDFInfo) bool{
	//filename:=GetFileName(url)
	resp, err := http.Get(pdfinfo.Url)
	//defer resp.Body.Close()
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return false;
	}
	body, err := ioutil.ReadAll(resp.Body)
	// 首先看一下如何将一个字符串写入文件
	err=ioutil.WriteFile(pdfinfo.Title+"_"+pdfinfo.SourceId+".pdf", body, os.ModePerm)
	if err==nil{
		return true
	}
	return false
	//err := ioutil.WriteFile("dat1", body, 0644)
}

func DownloadPDF(pdflist *list.List,insert func(pdfinfo PDFInfo)) {
	//filename:=GetFileName(url)
	for e := pdflist.Front(); e != nil; e = e.Next() {
		pdfinfo:=e.Value.(PDFInfo)
		success:=downloadPDF(pdfinfo)
		if success{
			insert(pdfinfo)
		}
	}
}


func BuildMultiPageUrls(engine ISearchEngine,pagenum int,pageend int ) []string{
	urls:=make([]string,pageend-pagenum+1)
	for i:=0;pagenum<=pageend;pagenum++{
		urls[i]=engine.BuildSinglePageUrl(pagenum)
		i++
	}
	return urls
}

func RequestUrls(urls []string)[]*http.Response{
	resps:=make([]*http.Response,0)
	for i:=0;i<len(urls);i++{
		resp, err := http.Get(urls[i])
		if err ==nil{
			resps=append(resps,resp)
		}
	}
	return resps
}