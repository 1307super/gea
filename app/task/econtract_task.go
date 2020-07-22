package task

import (
	"fmt"
	"time"
)

func init() {
	//fmt.Println("=========== 电子合同定时任务 ========= ")
	var task TaskEntity
	task.JobName = "电子合同定时同步状态"
	task.Param = nil
	task.Func = Test
	Add(task)
}

type TestS struct {
	A string
	B int
}

func Test() interface{} {
	time.Sleep(1*time.Second)
	fmt.Println("我是test任务")
	time.Sleep(1*time.Second)

	return TestS{
		A: "12",
		B: 10,
	}
}