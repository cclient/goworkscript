package spider

func StartSpider() {
	ch := make(chan bool)
	for {
		go DoTasks(ch,20)
		<-ch
	}
}