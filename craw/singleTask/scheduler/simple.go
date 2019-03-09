package scheduler

import "moocGo/craw/singleTask/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

/*
//重构前

func (s *SimpleScheduler) Submit(r engine.Request) {
	//不启用goroutine 会在out和in之间死循环
	go func() {
		s.workerChan <- r
		}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
*/

func (s *SimpleScheduler) Submit(r engine.Request) {
	//不启用goroutine 会在out和in之间死循环
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}
