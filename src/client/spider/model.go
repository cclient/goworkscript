package spider


type RequestTasks struct {
	User  string `json:"user"bson:"user"`
	Token string `json:"token"bson:"token"`
	Tasks []TaskSimple `json:"tasks"bson:"tasks"`
}

type Tasks struct {
	User       string `json:"user"bson:"user"`
	Url        string `json:"url"bson:"url"`
	Interval   int `json:"interval"bson:"interval"`
	Handler    string `json:"handler"bson:"handler"`
	Timeout    int `json:"timeout"bson:"timeout"`
	Priority   int `json:"priority"bson:"priority"`
	Index      int `json:"index"bson:"index"`
	Depth      int `json:"depth"bson:"depth"`
	Force      int `json:"force"bson:"force"`
	Extra      string `json:"extra"bson:"extra"`
	Js_script  string `json:"js_script"bson:"js_script"`
	Noderegion string `json:"noderegion"bson:"noderegion"`
}

type TaskSimple struct {
	Url   string `json:"url"bson:"url"`
	Extra string `json:"extra"bson:"extra"`
}

type ResponseTask struct {
	Tasks []TaskStatus `json:"tasks"bson:"tasks"`
}
type TaskStatus struct {
	Status string `json:"status"bson:"status"`
	Id     string `json:"id"bson:"id"`
}

type TaskResult struct {
	Id          string `json:"id"bson:"id"`
	Status      string `json:"status"bson:"status"`
	User        string `json:"user"bson:"user"`
	Url         string `json:"url"bson:"url"`
	Status_code int `json:"status_code"bson:"status_code"`
	Headers     Headers `json:"headers"bson:"headers"`
	Content     string `json:"content"bson:"content"`
	Time        int `json:"time"bson:"time"`
	Elapse      int `json:"elapse"bson:"elapse"`
	Depth       int `json:"depth"bson:"depth"`
	Extra       string `json:"extra"bson:"extra"`
}

type NullResult struct {
}

type Headers struct {
}