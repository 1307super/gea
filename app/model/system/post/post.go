package post

import (
	"gea/app/utils/page"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
)

// Fill with you ideas below.
// Entity is the golang structure for table sys_post.
type EntityFlag struct {
	PostId     int64       `orm:"post_id,primary" json:"post_id"`     // 岗位ID
	PostCode   string      `orm:"post_code"       json:"post_code"`   // 岗位编码
	PostName   string      `orm:"post_name"       json:"post_name"`   // 岗位名称
	PostSort   int         `orm:"post_sort"       json:"post_sort"`   // 显示顺序
	Status     string      `orm:"status"          json:"status"`      // 状态（0正常 1停用）
	CreateBy   string      `orm:"create_by"       json:"create_by"`   // 创建者
	CreateTime *gtime.Time `orm:"create_time"     json:"create_time"` // 创建时间
	UpdateBy   string      `orm:"update_by"       json:"update_by"`   // 更新者
	UpdateTime *gtime.Time `orm:"update_time"     json:"update_time"` // 更新时间
	Remark     string      `orm:"remark"          json:"remark"`      // 备注
	Flag       bool        `json:"flag"`                              // 标记
}

//新增页面请求参数
type AddReq struct {
	PostName string `p:"postName"  v:"required#岗位名称不能为空"`
	PostCode string `p:"postCode"  v:"required#岗位编码不能为空"`
	PostSort int    `p:"postSort"  v:"required#显示顺序不能为空"`
	Status   string `p:"status"    v:"required#状态不能为空"`
	Remark   string `p:"remark"`
}

//修改页面请求参数
type EditReq struct {
	PostId   int64  `p:"postId" v:"required#主键ID不能为空"`
	PostName string `p:"postName"  v:"required#岗位名称不能为空"`
	PostCode string `p:"postCode"  v:"required#岗位编码不能为空"`
	PostSort int    `p:"postSort"  v:"required#显示顺序不能为空"`
	Status   string `p:"status"    v:"required#状态不能为空"`
	Remark   string `p:"remark"`
}

//分页请求参数
type SelectPageReq struct {
	PostCode      string `p:"postCode"`      //岗位编码
	Status        string `p:"status"`        //状态
	PostName      string `p:"postName"`      //岗位名称
	BeginTime     string `p:"beginTime"`     //开始时间
	EndTime       string `p:"endTime"`       //结束时间
	OrderByColumn string `p:"orderByColumn"` //排序字段
	IsAsc         string `p:"isAsc"`         //排序方式
	PageNum       int    `p:"pageNum"`       //当前页码
	PageSize      int    `p:"pageSize"`      //每页数
}

//检查编码请求参数
type CheckPostCodeReq struct {
	PostId   int64  `p:"postId"  v:"required#岗位ID不能为空"`
	PostCode string `p:"postCode"  v:"required#岗位编码不能为空"`
}

//检查编码请求参数
type CheckPostCodeALLReq struct {
	PostCode string `p:"postCode"  v:"required#岗位编码不能为空"`
}

//检查名称请求参数
type CheckPostNameReq struct {
	PostId   int64  `p:"postId"  v:"required#岗位ID不能为空"`
	PostName string `p:"postName"  v:"required#岗位名称不能为空"`
}

//检查名称请求参数
type CheckPostNameALLReq struct {
	PostName string `p:"postName"  v:"required#岗位名称不能为空"`
}

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_post p")

	if param != nil {
		if param.PostCode != "" {
			model.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}

		if param.Status != "" {
			model.Where("p.status = ", param.Status)
		}

		if param.PostName != "" {
			model.Where("p.post_name like ?", "%"+param.PostName+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Limit(page.StartNum, page.Pagesize)

	if param.OrderByColumn != "" {
		model.Order(param.OrderByColumn + " " + param.IsAsc)
	}

	var result []Entity
	model.Structs(&result)
	return result, page, nil
}

// 导出excel
func SelectListExport(param *SelectPageReq) (gdb.Result, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_post p")

	if param != nil {
		if param.PostCode != "" {
			model.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}

		if param.Status != "" {
			model.Where("p.status = ", param.Status)
		}

		if param.PostName != "" {
			model.Where("p.post_name like ?", "%"+param.PostName+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	//"岗位序号","岗位名称","岗位编码","岗位排序","状态"
	model.Fields("p.post_id,p.post_name,p.post_code,p.post_sort,p.status")

	result, _ := model.All()
	return result, nil
}

//获取所有数据
func SelectListAll(param *SelectPageReq) ([]EntityFlag, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_post p")
	if param != nil {

		if param.PostCode != "" {
			model.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}

		if param.Status != "" {
			model.Where("p.status = ", param.Status)
		}

		if param.PostName != "" {
			model.Where("p.post_name like ?", "%"+param.PostName+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	var result []EntityFlag
	model.Structs(&result)
	return result, nil
}

//根据用户ID查询岗位
func SelectPostsByUserId(userId int64) ([]EntityFlag, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("sys_post p")
	model.LeftJoin("sys_user_post up", "p.post_id = up.post_id")
	model.LeftJoin("sys_user u", "u.user_id = up.user_id")
	model.Where("u.user_id = ?", userId)
	model.Fields("p.post_id, p.post_name, p.post_code")

	var result []EntityFlag
	err = model.Structs(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//校验岗位名称是否唯一
func CheckPostNameUnique(postName string, postId int64) (*Entity, error) {
	return FindOne("post_id !=? and post_name=?", postId, postName)
}

//校验岗位名称是否唯一
func CheckPostNameUniqueAll(postName string) (*Entity, error) {
	return FindOne("post_name=?", postName)
}

//校验岗位名称是否唯一
func CheckPostCodeUnique(postCode string, postId int64) (*Entity, error) {
	return FindOne("post_id !=? and post_code=?", postId, postCode)
}

//校验岗位名称是否唯一
func CheckPostCodeUniqueAll(postCode string) (*Entity, error) {
	return FindOne("post_code=?", postCode)
}
