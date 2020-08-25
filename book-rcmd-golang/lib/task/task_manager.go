package task

import (
	"context"
	"fmt"
	"time"
)

func init() {

}

type Task func(exchange *DataExchange)

type DataExchange struct {
	ctx        context.Context
	notifyChan chan interface{}
}

type TaskManager struct {
	ctx         context.Context
	taskProcess chan string
	CurrentNum  int
}

func (tm *TaskManager) Init(ctx context.Context) {
	tm.ctx = ctx
	tm.taskProcess = make(chan string, tm.CurrentNum)
	for i := 0; i < tm.CurrentNum; i++ {
		tm.taskProcess <- fmt.Sprintf("task-%d", i+1)
	}
}

func (tm *TaskManager) doTask(name string, task Task, timeout time.Duration) {

	timeoutCtx, _ := context.WithTimeout(tm.ctx, timeout)
	dataExchange := &DataExchange{
		ctx:        timeoutCtx,
		notifyChan: make(chan interface{}),
	}

	go task(dataExchange)

	select {
	case <-dataExchange.notifyChan:
		return
	case <-timeoutCtx.Done():
		return
	}

}
