package task

import (
	"fmt"
	"gea/app/dao"
	"gea/app/model"
	"github.com/gogf/gf/os/gtime"
)

//任务信息
type TaskEntity struct {
	JobName   string
	Param     []string
	Func      func() interface{}
	StartTime int64
}

var taskList = make([]TaskEntity, 0)

//检查方法名是否存在
func GetByName(jobName string) *TaskEntity {
	var result TaskEntity
	for _, task := range taskList {
		if task.JobName == jobName {
			result = task
			break
		}
	}
	return &result
}

//增加Task方法
func Add(task TaskEntity) {
	if task.JobName == "" {
		return
	}
	if task.Func == nil {
		return
	}
	taskList = append(taskList, task)
}

//修改参数
func (t *TaskEntity) EditParams(jobName string, params []string) {
	for index := range taskList {
		if taskList[index].JobName == jobName {
			taskList[index].Param = params
			break
		}
	}
}

func (t *TaskEntity) Run() {
	t.before()
	res := t.Func()
	t.after(res)
}

// 任务执行前
func (t *TaskEntity) before() {
	// 记录开始时间
	t.StartTime = gtime.TimestampMilli()
}

// 任务执行后
func (t *TaskEntity) after(res interface{}) {
	endTime := gtime.TimestampMilli()
	var jobLog model.SysJobLog
	jobLog.JobName = t.JobName
	jobLog.Status = "0"
	jobLog.CreateTime = gtime.Now()
	//jobLogService.DeleteRecordById()
	//获取任务信息
	job, err := dao.SysJob.FindOne(dao.SysJob.Columns.JobName, t.JobName)
	if err != nil {
		// 执行失败
		jobLog.Status = "1"
		message := fmt.Sprintf("未查询到任务 【%s】", t.JobName)
		jobLog.ExceptionInfo = message
		jobLog.JobMessage = message
	} else {
		// 计算运行时间
		runTime := endTime - t.StartTime
		jobLog.JobMessage = fmt.Sprintf("%s，执行结果：%#v 总共耗时：%d毫秒", t.JobName, res, runTime)
		jobLog.JobGroup = job.JobGroup
		jobLog.InvokeTarget = job.InvokeTarget
	}
	dao.SysJobLog.Data(jobLog).Insert()
}
