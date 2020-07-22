// ==========================================================================
// GEAGO自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-18 15:44:13
// 生成路径: app/service/module/job/job_service.go
// 生成人：yunjie
// ==========================================================================
package job

import (
	jobModel "gea/app/model/monitor/job"
	userModel "gea/app/model/system/user"
	"gea/app/task"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gtime"
	"strings"
)

//根据主键查询数据
func SelectRecordById(id int64) (*jobModel.Entity, error) {
	return jobModel.FindOne("job_id", id)
}
//根据任务名称查询数据
func SelectRecordByJobName(jobName string) (*jobModel.Entity, error) {
	return jobModel.FindOne("job_name", jobName)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := jobModel.Delete("job_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}
	return false
}

//批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := jobModel.Delete("job_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//添加数据
func AddSave(req *jobModel.AddReq, r *ghttp.Request) (int64, error) {
	//检查任务名称是否存在
	rs := gcron.Search(req.JobName)
	if rs != nil {
		return 0, gerror.New("任务名称已经存在")
	}

	//可以task目录下是否绑定对应的方法
	//f := task.GetByName(req.JobName)
	//if f == nil {
	//	return 0, gerror.New("当前task目录下没有绑定这个方法")
	//}

	var entity jobModel.Entity
	entity.JobName = req.JobName
	entity.JobParams = req.JobParams
	entity.JobGroup = req.JobGroup
	entity.InvokeTarget = req.InvokeTarget
	entity.CronExpression = req.CronExpression
	entity.MisfirePolicy = req.MisfirePolicy
	entity.Concurrent = req.Concurrent
	entity.Status = req.Status
	entity.Remark = req.Remark
	entity.CreateTime = gtime.Now()
	entity.CreateBy = ""

	user, _ := userModel.SelectUserByUid(r.GetInt64("jwtUid"))

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	result, err := entity.Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil || id <= 0 {
		return 0, err
	}
	if req.Status == "0" {
		// 启动
		if err = Start(&entity); err != nil {
			return id,gerror.New("启动任务失败，您可以手动启动")
		}
	}
	return id, nil
}

//修改数据
func EditSave(req *jobModel.EditReq, r *ghttp.Request) (int64, error) {
	//检查任务名称是否存在
	tmp := gcron.Search(req.JobName)

	if tmp != nil {
		gcron.Remove(req.JobName)
	}

	//可以task目录下是否绑定对应的方法
	//f := task.GetByName(req.JobName)
	//if f.Func == nil {
	//	return 0, gerror.New("当前task目录下没有绑定这个方法")
	//}

	entity, err := jobModel.FindOne("job_id=?", req.JobId)

	if err != nil {
		return 0, err
	}

	if entity == nil {
		return 0, gerror.New("数据不存在")
	}

	entity.JobName = req.JobName
	entity.InvokeTarget = req.InvokeTarget
	entity.JobParams = req.JobParams
	entity.CronExpression = req.CronExpression
	entity.MisfirePolicy = req.MisfirePolicy
	entity.Concurrent = req.Concurrent
	entity.Status = req.Status
	entity.Remark = req.Remark
	entity.UpdateTime = gtime.Now()
	entity.UpdateBy = ""

	user,_ := userModel.SelectUserByUid(r.GetInt64("jwtUid"))

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	result, err := entity.Update()

	if err != nil {
		return 0, err
	}

	rs, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}
	if req.Status == "0" {
		// 启动
		if err = Start(entity); err != nil {
			return rs,gerror.New("启动任务失败，您可以手动启动")
		}
	}

	return rs, nil
}

//启动任务
func Start(entity *jobModel.Entity) error {
	//可以task目录下是否绑定对应的方法
	f := task.GetByName(entity.JobName)
	if f.Func == nil {
		return gerror.New("当前task目录下没有绑定这个方法")
	}

	//传参
	paramArr := strings.Split(entity.JobParams, "|")
	f.EditParams(f.JobName, paramArr)
	rs := gcron.Search(entity.JobName)
	if rs == nil {
		var err error
		if entity.MisfirePolicy == "1" {
			// 判断是否并发执行
			if entity.Concurrent == "1"{
				// 允许并发
				_, err = gcron.Add(entity.CronExpression, f.Run, entity.JobName)
			}else{
				// 同时只允许有一个该任务运行
				_, err = gcron.AddSingleton(entity.CronExpression, f.Run, entity.JobName)
			}
			if err != nil {
				return err
			}

		} else {
			_, err := gcron.AddOnce(entity.CronExpression, f.Run, entity.JobName)
			if err != nil {
				return err
			}
		}
	}
	gcron.Start(entity.JobName)
	if entity.MisfirePolicy == "1" {
		entity.Status = "0"
		entity.Update()
	}

	return nil
}

//停止任务
func Stop(entity *jobModel.Entity) error {
	//可以task目录下是否绑定对应的方法
	f := task.GetByName(entity.JobName)
	if f == nil {
		return gerror.New("当前task目录下没有绑定这个方法")
	}

	rs := gcron.Search(entity.JobName)

	if rs != nil {
		gcron.Stop(entity.JobName)
	}

	entity.Status = "1"
	entity.Update()
	return nil
}

//根据条件查询数据
func SelectListAll(params *jobModel.SelectPageReq) ([]jobModel.Entity, error) {
	return jobModel.SelectListAll(params)
}

//根据条件分页查询数据
func SelectListByPage(params *jobModel.SelectPageReq) ([]jobModel.Entity, *page.Paging, error) {
	return jobModel.SelectListByPage(params)
}

// 导出excel
func Export(param *jobModel.SelectPageReq) (string, error) {
	result, err := jobModel.SelectListExport(param)
	if err != nil {
		return "", err
	}

	head := []string{"任务ID", "任务名称", "任务组名", "调用目标字符串", "cron执行表达式", "计划执行错误策略（1立即执行 2执行一次 3放弃执行）", "是否并发执行（0允许 1禁止）", "状态（0正常 1暂停）", "创建者", "创建时间", "更新者", "更新时间", "备注信息"}
	key := []string{"job_id", "job_name", "job_group", "invoke_target", "cron_expression", "misfire_policy", "concurrent", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	url, err := excel.DownlaodExcel(head, key, result)

	if err != nil {
		return "", err
	}

	return url, nil
}
