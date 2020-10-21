package logininfor

import (
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
	"time"
	"gea/app/model/monitor/logininfor"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
)

const USER_NOPASS_TIME string = "user_nopass_"
const USER_LOCK string = "user_lock_"

// 根据条件分页查询用户列表
func SelectPageList(param *logininfor.SelectPageReq) ([]logininfor.Entity, *page.Paging, error) {
	return logininfor.SelectPageList(param)
}

//根据主键查询用户信息
func SelectRecordById(id int64) (*logininfor.Entity, error) {
	return logininfor.FindOne("info_id", id)
}

//根据主键删除用户信息
func DeleteRecordById(id int64) bool {
	result, err := logininfor.Delete("info_id", id)
	if err == nil {
		affected, _ := result.RowsAffected()
		if affected > 0 {
			return true
		}
	}

	return false
}

//批量删除记录
func DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := logininfor.Delete("info_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//清空记录
func DeleteRecordAll() int64 {
	result, err := logininfor.Delete()
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

// 导出excel
func Export(param *logininfor.SelectPageReq) (string, error) {
	result, err := logininfor.SelectExportList(param)
	if err != nil {
		return "", err
	}

	head := []string{"访问编号", "登录名称", "登录地址", "登录地点", "浏览器", "操作系统", "登录状态", "操作信息", "登录时间"}
	key := []string{"info_id", "login_name", "ipaddr", "login_location", "browser", "os", "status", "msg", "login_time"}
	url, err := excel.DownlaodExcel(head, key, result)

	if err != nil {
		return "", err
	}

	return url, nil
}

//记录密码尝试次数
func SetPasswordCounts(loginName string) int {
	curTimes := 0
	curTimeObj, _ := gcache.Get(USER_NOPASS_TIME + loginName)
	if curTimeObj != nil {
		curTimes = gconv.Int(curTimeObj)
	}
	curTimes = curTimes + 1
	gcache.Set(USER_NOPASS_TIME+loginName, curTimes, 1*time.Minute)

	if curTimes >= 5 {
		Lock(loginName)
	}
	return curTimes
}

//移除密码错误次数
func RemovePasswordCounts(loginName string) {
	gcache.Remove(USER_NOPASS_TIME + loginName)
}

//锁定账号
func Lock(loginName string) {
	gcache.Set(USER_LOCK+loginName, true, 30*time.Minute)
}

//解除锁定
func Unlock(loginName string) {
	gcache.Remove(USER_LOCK + loginName)
}

//检查账号是否锁定
func CheckLock(loginName string) bool {
	result := false
	rs,_ := gcache.Get(USER_LOCK + loginName)
	if rs != nil {
		result = true
	}
	return result
}
