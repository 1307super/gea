package post

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	postModel "gea/app/model/system/post"
	userService "gea/app/service/system/user"
	"gea/app/utils/convert"
	"gea/app/utils/excel"
	"gea/app/utils/page"
)

//根据主键查询数据
func SelectRecordById(id int64) (*postModel.Entity, error) {
	return postModel.FindOne("post_id", id)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := postModel.Delete("post_id", id)
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
	result, err := postModel.Delete("post_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//添加数据
func AddSave(req *postModel.AddReq, r *ghttp.Request) (int64, error) {
	var post postModel.Entity
	post.PostName = req.PostName
	post.PostCode = req.PostCode
	post.Status = req.Status
	post.PostSort = req.PostSort
	post.Remark = req.Remark
	post.CreateTime = gtime.Now()
	post.CreateBy = ""

	user, _ := userService.GetProfileApi(r.GetInt64("jwtUid"))

	if user != nil {
		post.CreateBy = user.LoginName
	}

	result, err := post.Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}

//修改数据
func EditSave(req *postModel.EditReq, r *ghttp.Request) (int64, error) {

	post, err := postModel.FindOne("post_id=?", req.PostId)

	if err != nil {
		return 0, err
	}

	if post == nil {
		return 0, gerror.New("数据不存在")
	}

	post.PostName = req.PostName
	post.PostCode = req.PostCode
	post.Status = req.Status
	post.Remark = req.Remark
	post.PostSort = req.PostSort
	post.UpdateTime = gtime.Now()
	post.UpdateBy = ""

	user,_ := userService.GetProfileApi(r.GetInt64("jwtUid"))

	if user == nil {
		post.UpdateBy = user.LoginName
	}

	result, err := post.Update()

	if err != nil {
		return 0, err
	}

	rs, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rs, nil
}

//根据条件分页查询角色数据
func SelectListAll(params *postModel.SelectPageReq) ([]postModel.EntityFlag, error) {
	return postModel.SelectListAll(params)
}

//根据条件分页查询角色数据
func SelectListByPage(params *postModel.SelectPageReq) ([]postModel.Entity, *page.Paging, error) {
	return postModel.SelectListByPage(params)
}

// 导出excel
func Export(param *postModel.SelectPageReq) (string, error) {
	result, err := postModel.SelectListExport(param)
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

//根据用户ID查询岗位
func SelectPostsByUserId(userId int64) ([]postModel.EntityFlag, error) {
	userPosts, err := postModel.SelectPostsByUserId(userId)

	return userPosts, err
}
//根据用户ID查询岗位
func SelectPostsIdByUserId(userId int64) ([]int64, error) {
	userPosts, err := postModel.SelectPostsByUserId(userId)
	var userPostIds []int64
	for _, userPost := range userPosts {
		userPostIds = append(userPostIds, userPost.PostId)
	}
	return userPostIds, err
}

//检查角色名是否唯一
func CheckPostNameUniqueAll(postName string) string {
	post, err := postModel.CheckPostNameUniqueAll(postName)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

//检查岗位名称是否唯一
func CheckPostNameUnique(postName string, postId int64) string {
	post, err := postModel.CheckPostNameUnique(postName, postId)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

//检查岗位编码是否唯一
func CheckPostCodeUniqueAll(postCode string) string {
	post, err := postModel.CheckPostCodeUniqueAll(postCode)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}

//检查岗位编码是否唯一
func CheckPostCodeUnique(postCode string, postId int64) string {
	post, err := postModel.CheckPostCodeUnique(postCode, postId)
	if err != nil {
		return "1"
	}
	if post != nil && post.PostId > 0 {
		return "1"
	}
	return "0"
}
