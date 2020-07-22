package post

import (
	"gea/app/controller"
	"gea/app/model"
	postModel "gea/app/model/system/post"
	postService "gea/app/service/system/post"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller) Init(r *ghttp.Request) {
	c.Module = "岗位管理"
}

//列表分页数据
func (c *Controller) Get(r *ghttp.Request) {
	var req *postModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r, err.Error())
	}
	rows := make([]postModel.Entity, 0)
	result, page, err := postService.SelectListByPage(req)
	if err == nil && len(result) > 0 {
		rows = result
	}
	c.SuccTable(r, page.Total, rows)
}

//新增页面保存
func (c *Controller) Post(r *ghttp.Request) {
	var req *postModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r, err.Error())
	}

	if postService.CheckPostNameUniqueAll(req.PostName) == "1" {
		c.Err(r, "岗位名称已存在")
	}

	if postService.CheckPostCodeUniqueAll(req.PostCode) == "1" {
		c.Err(r, "岗位编码已存在")
	}

	pid, err := postService.AddSave(req, r)

	if err != nil || pid <= 0 {
		c.Err(r, "新增失败")
	}
	c.Succ(r, pid)
}


//修改页面保存
func (c *Controller) Put(r *ghttp.Request) {
	var req *postModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	if postService.CheckPostNameUnique(req.PostName, req.PostId) == "1" {
		c.Err(r,"岗位名称已存在")
	}

	if postService.CheckPostCodeUnique(req.PostCode, req.PostId) == "1" {
		c.Err(r,"岗位编码已存在")
	}

	rs, err := postService.EditSave(req, r)

	if err != nil || rs <= 0 {
		c.Err(r,"修改失败")
	}
	c.Succ(r,rs)
}

//删除数据
func(c *Controller) Delete(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	rs := postService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		c.Succ(r)
	} else {
		c.Err(r,"删除失败")
	}
}

func (c *Controller) Info(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		c.Err(r,"参数错误")
	}
	post, err := postService.SelectRecordById(id)
	if err != nil || post == nil {
		c.Err(r,"岗位不存在")
	}
	c.Succ(r,post)
}


//导出
func (c *Controller)Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Export)
	var req *postModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	url, err := postService.Export(req)

	if err != nil {
		c.Err(r,err.Error())
	}
	c.Succ(r,url)
}