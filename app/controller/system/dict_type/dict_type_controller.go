package dict_type

import (
	"gea/app/controller"
	"gea/app/model"
	dictTypeModel "gea/app/model/system/dict_type"
	dictTypeService "gea/app/service/system/dict_type"
	"gea/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller)Init(r *ghttp.Request) {
	c.Module = "字典管理"
}

//列表分页数据
func (c *Controller) Get(r *ghttp.Request) {
	var req *dictTypeModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	result, page, err := dictTypeService.SelectListByPage(req)
	if err != nil {
		 c.Err(r,err.Error())
	}
	rows := make([]dictTypeModel.Entity, 0)
	rows = result
	c.SuccTable(r,page.Total,rows)
}

//新增页面保存
func (c *Controller) Post(r *ghttp.Request) {
	var req *dictTypeModel.AddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	if dictTypeService.CheckDictTypeUniqueAll(req.DictType) == "1" {
		 c.Err(r,"字典类型已存在")
	}

	rid, err := dictTypeService.AddSave(req, r)

	if err != nil || rid <= 0 {
		 c.Err(r,"")
	}
	 c.Succ(r,rid)
}

//修改页面保存
func (c *Controller) Put(r *ghttp.Request) {
	var req *dictTypeModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	if dictTypeService.CheckDictTypeUnique(req.DictType, req.DictId) == "1" {
		 c.Err(r,"字典类型已存在")
	}

	rs, err := dictTypeService.EditSave(req, r)

	if err != nil || rs <= 0 {
		 c.Err(r,"")
	}
	 c.Succ(r,)
}

//删除数据
func (c *Controller) Delete(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}

	rs := dictTypeService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		 c.Succ(r,)
	} else {
		 c.Err(r,"删除失败")
	}
}
//数据详情
func (c *Controller) Info(r *ghttp.Request) {
	dictId := gconv.Int64(r.Get("dictId"))
	if dictId <= 0 {
		 c.Err(r,"请选择字典类型")
	}
	dict, _ := dictTypeService.SelectRecordById(dictId)

	if dict == nil {
		 c.Err(r,"字典不存在")
	}
	 c.Succ(r,dict)
}

func (c *Controller) Optionselect(r *ghttp.Request) {
	var req *dictTypeModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	result, err := dictTypeService.SelectListAll(req)
	if err != nil {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,result)
}

//导出
func (c *Controller) Export(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Export)
	var req *dictTypeModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		 c.Err(r,err.Error())
	}
	url, err := dictTypeService.Export(req)
	if err != nil {
		 c.Err(r,err.Error())
	}
	 c.Succ(r,url)
}
