# machineryDemo
learn machinery

### machinery的简单使用
```go
├── main.go
├── sendtask.go
└── worker
    ├── tasks.go
    └── worker.go

```
#### main
启动服务和worker，发送执行请求，发送请求后可以看到函数的执行。
```go
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
		log.Println("send task")
		worker.SendHelloWorldTask(context.Background(), AsyncTaskCenter)
	}(server)
	//启动worker
	taskWorker := worker.NewAsyncTaskWorker(0, server)
	taskWorker.Launch()
	time.Sleep(3 * time.Second)
}
```
#### worker.go
`NewTaskCenter`主要是根据配置文件启动`server`，并且`RegisterTasks`注册函数，也就是`machinery`是先注册函数名和实现，然后再根据请求来决定调用时机。
```go
func NewTaskCenter() (*machinery.Server, error) {
	...
	server, err := machinery.NewServer(confg)
	...
	return server, server.RegisterTasks(asyncTaskMap)
}
```
`NewAsyncTaskWorker`启动`worker`，`worker`作为执行者并不需要太多配置，设置并发数量和一些处理函数即可。
```go
func NewAsyncTaskWorker(concurrency int, AsyncTaskCenter *machinery.Server) *machinery.Worker {
	...
	worker := AsyncTaskCenter.NewWorker(consumerTag, concurrency)
	...
	worker.SetPostTaskHandler(posttaskhandler)
	worker.SetErrorHandler(errorhandler)
	worker.SetPreTaskHandler(pretaskhandler)
	return worker
}
```
#### tasks.go
`SendHelloWorldTask`发送函数执行请求，请求是以`函数名+参数`构造的签名为单位进行发送。即请求可以设置函数入参。
```go
func SendHelloWorldTask(ctx context.Context, server *machinery.Server) {
	args := make([]tasks.Arg, 0)
	task, _ := tasks.NewSignature(HelloWorldTaskName, args)
	task.RetryCount = 3
	server.SendTaskWithContext(ctx, task)
}
```
#### 执行结果
```text
...
INFO: 2022/10/29 21:27:14 worker.go:58 Launching a worker with the following settings:
INFO: 2022/10/29 21:27:14 worker.go:59 - Broker: redis://localhost:6379
INFO: 2022/10/29 21:27:14 worker.go:61 - DefaultQueue: ServerTasksQueue
INFO: 2022/10/29 21:27:14 worker.go:65 - ResultBackend: redis://localhost:6379
2022/10/29 21:27:14 send task
INFO: 2022/10/29 21:27:14 redis.go:105 [*] Waiting for messages. To exit press CTRL+C
DEBUG: 2022/10/29 21:27:14 redis.go:347 Received new message: {"UUID":"task_7bde3698-6d24-4bbe-a4ea-9eb64bfd9422","Name":"HelloWorldTask","RoutingKey":"ServerTasksQueue","ETA":null,"GroupUUID":"","GroupTaskCount":0,"Args":[],"Headers":{},"Priority":0,"Immutable":false,"RetryCount":3,"RetryTimeout":0,"OnSuccess":null,"OnError":null,"ChordCallback":null,"BrokerMessageGroupId":"","SQSReceiptHandle":"","StopTaskDeletionOnError":false,"IgnoreWhenTaskNotRegistered":false}
INFO: 2022/10/29 21:27:14 worker.go:37  I am a start of task handler for: HelloWorldTask 
Hello World!
DEBUG: 2022/10/29 21:27:14 worker.go:261 Processed task task_7bde3698-6d24-4bbe-a4ea-9eb64bfd9422. Results = []
INFO: 2022/10/29 21:27:14 worker.go:41  I am an end of task handler for: HelloWorldTask 
^CWARNING: 2022/10/29 21:27:17 worker.go:101 Signal received: interrupt
WARNING: 2022/10/29 21:27:17 worker.go:106 Waiting for running tasks to finish before shutting down
WARNING: 2022/10/29 21:27:17 broker.go:118 Stop channel

Process finished with the exit code 137 (interrupted by signal 9: SIGKILL)

```

### 代码阅读
[!读代码](./readCode.md)