package spider

import (
	ll "container/list"

)


func ParRunTaskLimitMaxConcurrence(tasks []interface{},gourp_url_count int,ff func([]interface{},int,chan interface{})) *ll.List{
	// tasks 所有要请求的url
	// gourp_url_count每批次并行请求的数量（这里就是指同时request的数量，根据网络和机型设置）
	allurlscount := len(tasks)
	//计算出一共需要执行几组，注意边界处理，例 21个url 每组5条，则要分5组（最后一组只有1条）
	group_count := allurlscount / gourp_url_count //21 5
	//余数
	remainder := allurlscount % gourp_url_count
	if remainder != 0 {
		//有余数则多算一组
		group_count++ //6
	} else if group_count == 1 {
		//正好是一组
		remainder = gourp_url_count
	}
	result := ll.New().Init()
	//第个小组内 任务channel array
	//var chs []chan interface{}
	chs := make([]chan interface{}, gourp_url_count)

	//遍历每组 这里是串行的
	//例100个url 每组20条 分5组,组内20条并行,5个组则是串行
	for i := 0; i < group_count; i++ {
		//遍历该批次内的任务，请求url
		for j := 0; j < gourp_url_count; j++ {
			chs[j] = make(chan interface{})
			//遍历组内的url项
			//不是最后一组则请求组内所有url || 是最后一组，但序号小于余数的,请求该url。
			if (i != (group_count - 1)) || j < remainder {
				go ff(tasks, gourp_url_count * i + j, chs[j])
			} else {
				//最后一组 序号大于余数的
				go ff(tasks, -1, chs[j])
				//正好一组,实际都在这里执行
			}
		}
		//阻塞在这里，直到该批次内所有url都请求完毕。
		for _, ch := range chs {
			res := <-ch
			result.PushBack(res)
		}
	}
	return result
}
