package pdf

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)


//不同搜索引擎不同实现
type ISearchEngine interface{
	//单页构造查询url
	getId(pdfInfo PDFInfo) string
	//单页构造查询url
	BuildSinglePageUrl(pagenum int) string
	//单页解析pdfinfo
	ParseContainPdfInfo(resp *http.Response) []PDFInfo
}

type BingEngine struct{
}

func (t BingEngine) BuildSinglePageUrl(pagenum int) string{
	sUrl:="http://cn.bing.com/search?q=site:ppt.geekbang.org+ppt"
	if pagenum==1{
		return sUrl
	}
	return sUrl+"&first="+strconv.Itoa((pagenum-1)*10+1)+"&FORM=PERE"+strconv.Itoa(pagenum-2)
}

func (t BingEngine) getId(pdfinfo PDFInfo) string{
	arr1:=strings.Split(pdfinfo.Url,"/")
	tstr:=arr1[len(arr1)-1]
	arr2:=strings.Split(tstr,".")
	return arr2[0]
}

func (t BingEngine) ParseContainPdfInfo(resp *http.Response) []PDFInfo{
	pdfInfos:=make([]PDFInfo,0)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		//panic(err)
		return pdfInfos
	}
	//<li class="b_algo" data-bm="6"><div class="b_title"><div class="b_imagePair square_mi"><div class="inner"><a target="_blank" href="http://ppt.geekbang.org/slide/download/409/5808242245ec5.pdf" h="ID=SERP,5117.1"><img class="rms_img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAA8UlEQVQ4y5WTQQqCQBSGvUCb7tCqm3SGjtFRahG0adOyO7hspUT7JCEkaCEhaSoT38gTy3GaBn6Gef7ve8+n46nO2vuJWu4uP9VdXvewWPnqVVaDIjnNmv1wzBtANB4pFxVF0QIE8hdAkgQgZ2dAkpbGeXjb6cS5AyBSXV6pN8Qsr7VRdN+s1XU+U+db3kIQQ8Vv/ApiIgHARwdpqnfxDwIwUpXqImJxHNs7qOtKJ2J6ngKdKF2IOFs7ACDJ7MCIoSAMdMw6A9qkkgzPJIZtBDAsqdSt/C18VsAvDQLC6KEr86OYLhRxnuPrAbjOBF2F/w0iWUsSETZLDgAAAABJRU5ErkJggg==" data-bm="23"></a></div><h2><a target="_blank" href="http://ppt.geekbang.org/slide/download/409/5808242245ec5.pdf" h="ID=SERP,5117.2">构建微服务体系下的全链路监控体系</a></h2></div></div><div class="b_caption"><p>唯品会 构建微服务体系下的全链路监控体系 姚捷@唯品会</p><div class="b_attribution"><cite><strong>ppt.geekbang.org</strong>/slide/download/409/5808242245ec5.<strong>pdf</strong></cite>&nbsp;· 2016-10-20</div></div></li>
	htmlbody := doc.Find("body")
	maindiv := htmlbody.Find("ol[id=b_results]")
	maindiv.Children().Each(func(num int, s *goquery.Selection) {
		a:=s.Find("h2").First().Children().First()
		pdfInfo := new(PDFInfo)
		pdfInfo.Engine="Bing"
		pdfInfo.Url=a.AttrOr("href","")
		pdfInfo.Title=a.Text()
		pdfInfo.Desc=s.Find("p").First().Text()
		pdfInfo.SourceId=t.getId(*pdfInfo)
		if len(pdfInfo.SourceId)==13{
			pdfInfos=append(pdfInfos,*pdfInfo)
		}

	})
	return pdfInfos
}
