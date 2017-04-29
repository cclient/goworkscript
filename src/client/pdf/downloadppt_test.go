package pdf

import (
	"testing"
	"fmt"
)

func Test_BuildSinglePageUrl(t *testing.T) {
	var w BingEngine  // 定义类型为watch的变量w
	var engine ISearchEngine  // 定义类型为iWatch的变量t
	engine = w             // 把类型为watch的变量w赋值给类型为iWatch的变量t，这样能行的通吗？
	url:=engine.BuildSinglePageUrl(1000)
	fmt.Println(url)
}


func Test_BuildPageUrls(t *testing.T){
	var w BingEngine  // 定义类型为watch的变量w
	var engine ISearchEngine  // 定义类型为iWatch的变量t
	engine = w             // 把类型为watch的变量w赋值给类型为iWatch的变量t，这样能行的通吗？
	urls:=BuildMultiPageUrls(engine,1,2)
	fmt.Println(urls)
}

func Test_GetUId(t *testing.T){
	var w BingEngine  // 定义类型为watch的变量w
	var engine ISearchEngine  // 定义类型为iWatch的变量t
	engine = w             // 把类型为watch的变量w赋值给类型为iWatch的变量t，这样能行的通吗？
	id:=engine.getId(PDFInfo{Engine:"http://ppt.geekbang.org/slide/download/427/58086a1fe2ede.pdf"})
	fmt.Println(id)
}


func Test_RequestUrls(t *testing.T){
	var w BingEngine  // 定义类型为watch的变量w
	var engine ISearchEngine  // 定义类型为iWatch的变量t
	engine = w             // 把类型为watch的变量w赋值给类型为iWatch的变量t，这样能行的通吗？
	urls:=BuildMultiPageUrls(engine,35,36)
	fmt.Println(urls)
	resps:=RequestUrls(urls)
	fmt.Println(resps)
	fmt.Println(resps[0])
	fmt.Println(*resps[0])
	info:=engine.ParseContainPdfInfo(resps[0])
	fmt.Println(info)
}

func Test_DownloadPDF(t *testing.T){
	downloadPDF(PDFInfo{Url:"http://ppt.geekbang.org/slide/download/427/58086a1fe2ede.pdf",SourceId:"58086a1fe2ede",Title:"testdownload.pdf"})
}

func Test_doWork(t *testing.T){
	DoWork(1,40)
}

func Test_Start(t *testing.T){
	//无限循环
	Start()
}