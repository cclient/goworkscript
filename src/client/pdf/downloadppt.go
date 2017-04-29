package pdf

import (
	"strings"
	"strconv"
	"fmt"
	"container/list"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"client/common"
	"reflect"
)
const mysqlconn="user:passwd@(127.0.0.1:3306)/geekbang?charset=utf8mb4"

func DoWork(pagestart int,pageend int ){
	var w BingEngine
	var engine ISearchEngine
	engine = w
	//搜索引擎页面url
	urls:=BuildMultiPageUrls(engine,pagestart,pageend)
	fmt.Println(urls)
	//页面返回
	resps:=RequestUrls(urls)
	//页面内pdf信息
	pdflist:=WebPageContainPdf(engine,resps)
	//比对mysql,过滤出未拿过的pdf信息
	pdflist=NeedDownloadPDFList(pdflist)
	if pdflist.Len()==0{
		return
	}
	db, err := sql.Open("mysql", mysqlconn)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	fields,_,signs:=GetTypeFileds()

	DownloadPDF(pdflist,func(pdfinfo PDFInfo){
		//fmt.Println("INSERT INTO pdf(id,"+strings.Join(fields,",")+") VALUES(0, "+strings.Join(signs,",")+")")
		stmtIns, err := db.Prepare("INSERT INTO pdf(id,"+strings.Join(fields,",")+") VALUES(0, "+strings.Join(signs,",")+")") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
		immutable := reflect.ValueOf(pdfinfo)
		values:=make([]interface{},len(fields))
		for j:=0;j<len(fields);j++{
			values[j]=immutable.Field(j).Interface()
		}
		fmt.Println("values",values)
		_, err = stmtIns.Exec(values...) // Insert tuples (i, i^2)
		//if err != nil {
		//	panic(err.Error()) // proper error handling instead of panic in your app
		//}
	})
}

func GetTypeFileds() ([]string,[]bool,[]string){
	var i PDFInfo
	t:=reflect.TypeOf(i)
	keys:=make([]string,t.NumField())
	rowkeys:=make([]bool,t.NumField())
	signs:=make([]string,t.NumField())
	//values:=make([]bool,t.NumField())
	for j:=0;j<t.NumField();j++{
		keys[j]="`"+t.Field(j).Tag.Get("json")+"`"
		signs[j]="?"
		//values[j]=t.FieldByIndex(j)
		if t.Field(j).Type.Name()=="string"{
			rowkeys[j]=true
		}else{
			rowkeys[j]=false
		}
	}
	return keys,rowkeys,signs
}


func getContainSIds(rows *sql.Rows) []string {
	lhasnames := list.New().Init()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		lhasnames.PushBack(name)
	}
	return common.StringListToArray(lhasnames)
}

func getNoContainSIds(oldl *list.List, hasnames []string) *list.List {
	var linserMovienames = list.New().Init()
	for e := oldl.Front(); e != nil; e = e.Next() {
		pdfinfo := e.Value.(PDFInfo)
		isinsert := true
		for j := 0; j < len(hasnames); j++ {
			if pdfinfo.SourceId == hasnames[j] {
				isinsert = false
				break
			}
		}
		if isinsert == true {
			linserMovienames.PushBack(pdfinfo)
		}
	}
	return linserMovienames
}

func NeedDownloadPDFList(pdflist *list.List) *list.List{
	db, err := sql.Open("mysql", "root:admaster@(127.0.0.1:3306)/geekbang?charset=utf8mb4")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	sourceids := make([]string, 0)
	for e := pdflist.Front(); e != nil; e = e.Next() {
		pdfinfo:=e.Value.(PDFInfo)
		sourceids=append(sourceids,"'"+pdfinfo.SourceId+"'")
	}
	sourceidjoinstr := strings.Join(sourceids, ",")
	stmtOut, err := db.Prepare("SELECT source_id FROM pdf WHERE source_id in (" + sourceidjoinstr + ")")
	//fmt.Print("SELECT source_id FROM pdf WHERE source_id in (" + sourceidjoinstr + ")")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
	rows, err := stmtOut.Query()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	containsids := getContainSIds(rows)
	fmt.Println("has count" + strconv.Itoa(len(containsids)))
	nocontainspdfs := getNoContainSIds(pdflist, containsids)
	fmt.Println("need download" + strconv.Itoa(nocontainspdfs.Len()))
	return nocontainspdfs
}


