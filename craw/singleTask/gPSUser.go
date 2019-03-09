/*
获取城市列表第一页信息
 */
package main

import (
	"moocGo/craw/singleTask/engine"
	"moocGo/craw/singleTask/scheduler"
	"moocGo/craw/singleTask/zhenai/parser"
)

//单任务版本
/*
func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
*/

/*
//多线程并发版本
func main() {
	e := engine.ConcurrentEngin{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100}

	e.Run(engine.Request{Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,})
}
*/

//下一页扩展和数据去重
func main() {
	e := engine.ConcurrentEngin{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100}

	e.Run(engine.Request{Url:"http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,})
}