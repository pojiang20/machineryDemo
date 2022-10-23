package main

import (
	"context"
	"log"
	"machneryDemo/worker"
	"time"
)

func main() {
	//发送任务
	go func() {
		time.Sleep(10 * time.Second)
		log.Println("send task")
		worker.SendHelloWorldTask(context.Background())
	}()
	//启动异步任务框架
	taskWorker := worker.NewAsyncTaskWorker(0)
	taskWorker.Launch()
}
