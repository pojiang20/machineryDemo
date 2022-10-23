package main

import (
	"context"
	"github.com/RichardKnop/machinery/v1"
	"log"
	"machneryDemo/worker"
	"time"
)

func main() {
	//启动服务
	server, err := worker.NewTaskCenter()
	if err != nil {
		log.Println(err)
		return
	}
	//发送任务
	go func(AsyncTaskCenter *machinery.Server) {
		time.Sleep(3 * time.Second)
		log.Println("send task")
		worker.SendHelloWorldTask(context.Background(), AsyncTaskCenter)
	}(server)
	//启动worker
	taskWorker := worker.NewAsyncTaskWorker(0, server)
	taskWorker.Launch()
	time.Sleep(10 * time.Second)
}
