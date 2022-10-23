package worker

import (
	"context"
	"fmt"
	"github.com/RichardKnop/machinery/v1/tasks"
)

var (
	asyncTaskMap map[string]interface{}
)

const (
	HelloWorldTaskName          = "HelloWorldTask"
	DeleteAppShareImageTaskName = "DeleteAppShareImageTask"
)

func HelloWorld() error {
	fmt.Println("Hello World!")
	return nil
}

func SendHelloWorldTask(ctx context.Context) {
	args := make([]tasks.Arg, 0)
	task, _ := tasks.NewSignature(HelloWorldTaskName, args)
	task.RetryCount = 3
	AsyncTaskCenter.SendTaskWithContext(ctx, task)
}

func initAsyncTaskMap() {
	asyncTaskMap = make(map[string]interface{})
	asyncTaskMap[HelloWorldTaskName] = HelloWorld
}
