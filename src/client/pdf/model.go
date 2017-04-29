package pdf

//pdf 文本信息
type PDFInfo struct {
	SourceId string `json:"source_id" bson:"source_id"`
	Engine string `json:"engine" bson:"engine"`
	Url string `json:"url" bson:"url"`
	Title string `json:"title" bson:"title"`
	Desc string `json:"desc" bson:"desc"`
}
