package api

import (
	"gea/app/system/admin/internal/define"
	"gea/app/system/admin/internal/service"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

var Post = new(postApi)

type postApi struct {
	BaseController
}

func (a *postApi) Init(r *ghttp.Request) {
	a.Module = "岗位管理"
}

//列表分页数据
func (a *postApi) Get(r *ghttp.Request) {
	var req *define.PostApiPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	result := service.Post.GetList(req)
	if result != nil {
		a.SuccTable(r, result.Total, result.List)
	}else{
		a.SuccTable(r, 0, nil)
	}
}

//新增页面保存
func (a *postApi) Post(r *ghttp.Request) {
	var req *define.PostApiCreateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r, err.Error())
	}
	pid, err := service.Post.Create(r.Context(),req)
	if err != nil || pid <= 0 {
		a.Err(r, err.Error())
	}
	if pid <= 0{
		a.Err(r, "新增失败")
	}
	a.Succ(r, pid)
}


//修改页面保存
func (a *postApi) Put(r *ghttp.Request) {
	var req *define.PostApiUpdateReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	rs, err := service.Post.Update(r.Context(),req)
	if err != nil || rs <= 0 {
		a.Err(r,err.Error())
	}
	if rs <= 0 {
		a.Err(r,"修改失败")
	}
	a.Succ(r,rs)
}

//删除数据
func(a *postApi) Delete(r *ghttp.Request) {
	var req *define.PostApiDeleteReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	rs := service.Post.Delete(req.Ids)
	if rs > 0 {
		a.Succ(r)
	} else {
		a.Err(r,"删除失败")
	}
}

func (a *postApi) Info(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id <= 0 {
		a.Err(r,"参数错误")
	}
	post, err := service.Post.Info(id)
	if err != nil || post == nil {
		a.Err(r,"岗位不存在")
	}
	a.Succ(r,post)
}


//导出
func (a *postApi)Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,response.Buniss_Export)
	var req *define.PostApiPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		a.Err(r,err.Error())
	}
	url, err := service.Post.Export(req)

	if err != nil {
		a.Err(r,err.Error())
	}
	a.Succ(r,url)
}
