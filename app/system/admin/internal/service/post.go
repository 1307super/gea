package service

import (
	"context"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Post = &postService{}

type postService struct{}

func (s *postService) Info (id int64) (*model.SysPost,error){
	return dao.SysPost.FindOne(dao.SysPost.Columns.PostId,id)
}

// 获取分页列表
func (s *postService) GetList(param *define.PostApiPageReq) *define.PostServiceList {
	m := dao.SysPost.As("p")
	if param != nil {
		if param.PostCode != "" {
			m = m.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}
		if param.Status != "" {
			m = m.Where("p.status = ", param.Status)
		}
		if param.PostName != "" {
			m = m.Where("p.post_name like ?", "%"+param.PostName+"%")
		}
		if param.BeginTime != "" {
			m = m.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}
		if param.EndTime != "" {
			m = m.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := m.Count()
	if err != nil {
		return nil
	}
	page := page.CreatePaging(param.PageNum, param.PageSize, total)
	m = m.Limit(page.StartNum, page.Pagesize)
	if param.OrderByColumn != "" {
		m = m.Order(param.OrderByColumn + " " + param.IsAsc)
	}
	result := &define.PostServiceList{
		Page:  page.PageNum,
		Total: page.Total,
		Size:  page.Pagesize,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}
// 获取所有
func (s *postService) GetAll(param *define.PostApiPageReq) []model.SysPostFlag {
	m := dao.SysPost.As("p")
	if param != nil {
		if param.PostCode != "" {
			m = m.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}

		if param.Status != "" {
			m = m.Where("p.status = ", param.Status)
		}

		if param.PostName != "" {
			m = m.Where("p.post_name like ?", "%"+param.PostName+"%")
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	var result []model.SysPostFlag
	m.Structs(&result)
	return result
}

func (s *postService)Export(param *define.PostApiPageReq)(string, error) {

	m := dao.SysPost.As("p")
	if param != nil {
		if param.PostCode != "" {
			m = m.Where("p.post_code like ?", "%"+param.PostCode+"%")
		}

		if param.Status != "" {
			m = m.Where("p.status = ", param.Status)
		}

		if param.PostName != "" {
			m = m.Where("p.post_name like ?", "%"+param.PostName+"%")
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	//"岗位序号","岗位名称","岗位编码","岗位排序","状态"
	m = m.Fields("p.post_id,p.post_name,p.post_code,p.post_sort,p.status")
	result, err := m.M.All()
	if err != nil {
		return "", err
	}
	head := []string{"岗位序号", "岗位名称", "岗位编码", "岗位排序", "状态"}
	key := []string{"post_id", "post_name", "post_code", "post_sort", "stat"}
	url, err := excel.DownlaodExcel(head, key, result)
	if err != nil {
		return "", err
	}
	return url, nil

}

//添加数据
func (s *postService) Create(ctx context.Context, req *define.PostApiCreateReq) (int64, error) {
	if s.CheckPostNameUniqueAll(req.PostName) {
		return 0, gerror.New("岗位名称已存在")
	}

	if s.CheckPostCodeUniqueAll(req.PostCode) {
		return 0, gerror.New("岗位编码已存在")
	}

	user := shared.Context.Get(ctx).User
	var post model.SysPost
	post.CreateTime = gtime.Now()
	post.CreateBy = user.UserExtend.LoginName
	var editReq *define.PostApiUpdateReq
	gconv.Struct(req, &editReq)
	return s.save(&post, editReq)
}

//修改数据
func (s *postService) Update(ctx context.Context, req *define.PostApiUpdateReq) (int64, error) {
	if s.CheckPostNameUnique(req.PostName, req.PostId) {
		return 0, gerror.New("岗位名称已存在")
	}

	if s.CheckPostCodeUnique(req.PostCode, req.PostId) {
		return 0, gerror.New("岗位编码已存在")
	}
	user := shared.Context.Get(ctx).User
	post, err := dao.SysPost.FindOne(dao.SysPost.Columns.PostId, req.PostId)
	if err != nil {
		return 0, err
	}
	if post == nil {
		return 0, gerror.New("数据不存在")
	}

	post.UpdateTime = gtime.Now()
	post.UpdateBy = user.UserExtend.LoginName
	return s.save(post, req)
}

func (s *postService) save(post *model.SysPost, req *define.PostApiUpdateReq) (int64, error) {
	post.PostName = req.PostName
	post.PostCode = req.PostCode
	post.Status = req.Status
	post.Remark = req.Remark
	post.PostSort = req.PostSort
	result, err := dao.SysPost.Data(post).Save()
	if err != nil {
		return 0, gerror.New("操作失败")
	}
	if post.PostId == 0 {
		// 新增
		rs, err := result.LastInsertId()
		if err != nil || rs <= 0 {
			return 0, gerror.New("保存失败")
		}
	} else {
		rs, err := result.RowsAffected()
		if err != nil || rs <= 0 {
			return 0, gerror.New("保存失败")
		}
	}
	return 1, nil

}

//批量删除数据记录
func (s *postService)Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dao.SysPost.Delete("post_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

// 根据用户获取岗位
func (s *postService) GetPostListByUid(uid int64) ([]int64, error) {
	m := dao.SysPost.As("p")
	m = m.LeftJoin("sys_user_post up", "p.post_id = up.post_id")
	m = m.LeftJoin("sys_user u", "u.user_id = up.user_id")
	m = m.Where("u.user_id = ?", uid)
	m = m.Fields("p.post_id, p.post_name, p.post_code")
	var userPosts []model.SysPostFlag
	err := m.Structs(&userPosts)
	if err != nil {
		return nil, err
	}
	var userPostIds []int64
	for _, userPost := range userPosts {
		userPostIds = append(userPostIds, userPost.PostId)
	}
	return userPostIds, err
}

//检查角色名是否唯一
func (s *postService) CheckPostNameUniqueAll(postName string) bool {
	post, err := dao.SysPost.FindOne(dao.SysPost.Columns.PostName, postName)
	if err != nil {
		return true
	}
	if post != nil && post.PostId > 0 {
		return true
	}
	return false
}

//检查岗位名称是否唯一
func (s *postService) CheckPostNameUnique(postName string, postId int64) bool {
	post, err := dao.SysPost.FindOne(g.Map{
		dao.SysPost.Columns.PostId:   postId,
		dao.SysPost.Columns.PostName: postName,
	})
	if err != nil {
		return true
	}
	if post != nil && post.PostId > 0 {
		return true
	}
	return false
}

//检查岗位编码是否唯一
func (s *postService) CheckPostCodeUniqueAll(postCode string) bool {
	post, err := dao.SysPost.FindOne(dao.SysPost.Columns.PostCode, postCode)
	if err != nil {
		return true
	}
	if post != nil && post.PostId > 0 {
		return true
	}
	return false
}

//检查岗位编码是否唯一
func (s *postService) CheckPostCodeUnique(postCode string, postId int64) bool {
	post, err := dao.SysPost.FindOne(g.Map{
		dao.SysPost.Columns.PostId:   postId,
		dao.SysPost.Columns.PostCode: postCode,
	})
	if err != nil {
		return true
	}
	if post != nil && post.PostId > 0 {
		return true
	}
	return false
}
