package gen

import (
	"gea/app/controller"
	"gea/app/model"
	tableModel "gea/app/model/tool/table"
	userService "gea/app/service/system/user"
	tableService "gea/app/service/tool/table"
	"gea/app/utils/convert"
	"gea/app/utils/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"os"
	"strings"
)

type Controller struct {
	controller.BaseController
}

func (c *Controller) Init(r *ghttp.Request) {
	c.Module = "代码生成管理"
}


func (c *Controller) Get (r *ghttp.Request) {
	var req *tableModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	rows := make([]tableModel.Entity, 0)
	result, page, err := tableService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	c.SuccTable(r,page.Total,rows)
}
//导入表结构（保存）
func (c *Controller) Post(r *ghttp.Request) {
	tables := r.GetQueryString("tables")
	if tables == "" {
		c.Err(r,"参数错误")
	}

	user,_ := userService.GetProfileApi(r.GetInt64("jwtUid"))
	if user == nil {
		c.Err(r,"登录超时")
	}

	operName := user.LoginName

	tableArr := strings.Split(tables, ",")
	tableList, err := tableService.SelectDbTableListByNames(tableArr)
	if err != nil {
		c.Err(r,err.Error())
	}

	if tableList == nil {
		c.Err(r,"请选择需要导入的表")
	}

	tableService.ImportGenTable(tableList, operName)
	c.Succ(r)
}
//修改数据保存
func (c *Controller)Put(r *ghttp.Request) {
	var req tableModel.EditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	_, err := tableService.SaveEdit(&req, r)
	if err != nil {
		c.Err(r,err.Error())
	}
	c.Succ(r)
}
//删除数据
func (c *Controller) Delete(r *ghttp.Request) {
	var req *model.RemoveReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}

	rs := tableService.DeleteRecordByIds(req.Ids)

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

	entity, err := tableService.SelectRecordById(id)

	if err != nil || entity == nil {
		c.Err(r,"参数不存在")
	}

	//goTypeTpl := tableService.GoTypeTpl()
	//queryTypeTpl := tableService.QueryTypeTpl()
	//htmlTypeTpl := tableService.HtmlTypeTpl()

	c.Succ(r,g.Map{
		"info":        entity,
		"rows": entity.Columns,
		//"goTypeTpl":    goTypeTpl,
		//"queryTypeTpl": queryTypeTpl,
		//"htmlTypeTpl":  htmlTypeTpl,
	})
}

//预览代码
func (c *Controller) Preview(r *ghttp.Request) {
	tableId := r.GetInt64("tableId")
	if tableId <= 0 {
		c.Err(r,"参数错误")
	}

	entity, err := tableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		c.Err(r,"数据不存在")
	}
	tableService.SetPkColumn(entity, entity.Columns)
	listKey := "vm/vue/index.vue.vm"
	listValue := ""
	listTmp := "vm/vue/index.html"

	appJsKey := "vm/js/api.js.vm"
	appJsValue := ""
	appJsTmp := "vm/js/api.html"

	treeKey := "vm/vue/tree.html.vm"
	treeValue := ""

	if entity.TplCategory == "tree" {
		listTmp = "vm/vue/index-tree.html"
	}

	sqlKey := "vm/sql/menu.sql.vm"
	sqlValue := ""
	entityKey := "vm/go/entity.go.vm"
	entityValue := ""
	modelKey := "vm/go/model.go.vm"
	modelValue := ""
	extendKey := "vm/go/extend.go.vm"
	extendValue := ""
	serviceKey := "vm/go/service.go.vm"
	serviceValue := ""
	routerKey := "vm/go/router.go.vm"
	routerValue := ""
	controllerKey := "vm/go/controller.go.vm"
	controllerValue := ""
	// vo dto
	voKey := "vm/go/vo.go.vm"
	voValue := ""
	dtoKey := "vm/go/dto.go.vm"
	dtoValue := ""

	if tmpList, err := r.Response.ParseTpl(listTmp, g.Map{"table": entity}); err == nil {
		listValue = tmpList
	}

	if entity.TplCategory == "tree" {
		if tmpTree, err := r.Response.ParseTpl("vm/vue/index-tree.html", g.Map{"table": entity}); err == nil {
			treeValue = tmpTree
		}
	}

	if tmpAppJs, err := r.Response.ParseTpl(appJsTmp, g.Map{"table": entity}); err == nil {
		appJsValue = tmpAppJs
	}

	if tmpEntity, err := r.Response.ParseTpl("vm/go/entity.html", g.Map{"table": entity}); err == nil {
		entityValue = tmpEntity
	}

	if tmpModel, err := r.Response.ParseTpl("vm/go/model.html", g.Map{"table": entity}); err == nil {
		modelValue = tmpModel
	}

	if tmpExtend, err := r.Response.ParseTpl("vm/go/extend.html", g.Map{"table": entity}); err == nil {
		extendValue = tmpExtend
	}

	if tmpService, err := r.Response.ParseTpl("vm/go/service.html", g.Map{"table": entity}); err == nil {
		serviceValue = tmpService
	}

	if tmpRouter, err := r.Response.ParseTpl("vm/go/router.html", g.Map{"table": entity}); err == nil {
		routerValue = tmpRouter
	}

	if tmpController, err := r.Response.ParseTpl("vm/go/controller.html", g.Map{"table": entity}); err == nil {
		controllerValue = tmpController
	}

	if tmpSql, err := r.Response.ParseTpl("vm/sql/sql.html", g.Map{"table": entity}); err == nil {
		sqlValue = tmpSql
	}

	// vo
	if tmpVo, err := r.Response.ParseTpl("vm/go/vo.html", g.Map{"table": entity}); err == nil {
		voValue = tmpVo
	}
	// dto
	if tmpDto, err := r.Response.ParseTpl("vm/go/dto.html", g.Map{"table": entity}); err == nil {
		dtoValue = tmpDto
	}

	if entity.TplCategory == "tree" {
		c.Succ(r,g.Map{
			listKey:       listValue,
			appJsKey:      appJsValue,
			treeKey:       treeValue,
			sqlKey:        sqlValue,
			entityKey:     entityValue,
			modelKey:      modelValue,
			extendKey:     extendValue,
			serviceKey:    serviceValue,
			routerKey:     routerValue,
			controllerKey: controllerValue,
			voKey: voValue,
			dtoKey: dtoValue,
		})
	} else {
		c.Succ(r,g.Map{
			listKey:       listValue,
			appJsKey:      appJsValue,
			sqlKey:        sqlValue,
			entityKey:     entityValue,
			modelKey:      modelValue,
			extendKey:     extendValue,
			serviceKey:    serviceValue,
			routerKey:     routerValue,
			controllerKey: controllerValue,
			voKey: voValue,
			dtoKey: dtoValue,
		})
	}

}

//生成代码
func (c *Controller) GenCode(r *ghttp.Request) {
	r.SetCtxVar(response.ResponseBunissType,model.Buniss_Gen)
	tableId := r.GetQueryString("tableId")
	if tableId == "" {
		c.Err(r,"参数错误")
	}

	tableIds := convert.ToInt64Array(tableId, ",")
	if len(tableIds) <= 0 {
		c.Err(r,"参数错误")
	}
	for _,tid := range tableIds {
		entity, err := tableService.SelectRecordById(tid)

		if err != nil || entity == nil {
			c.Err(r,"数据不存在")
		}

		tableService.SetPkColumn(entity, entity.Columns)

		listTmp := "vm/vue/index.html"
		if entity.TplCategory == "tree" {
			listTmp = "vm/vue/index-tree.html"
		}

		//获取当前运行时目录
		curDir, err := os.Getwd()
		if err != nil {
			c.Err(r,err.Error())
		}


		if tmpList, err := r.Response.ParseTpl(listTmp, g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", entity.BusinessName, "/index.vue"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpList)
				}
				f.Close()
			}
		}
		if tmpJs, err := r.Response.ParseTpl("vm/js/api.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", entity.BusinessName, "/index.js"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpJs)
				}
				f.Close()
			}
		}

		if tmpEntity, err := r.Response.ParseTpl("vm/go/entity.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/app/model/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, "_entity.go"}, "")
			if gfile.Exists(fileName) {
				gfile.Remove(fileName)
			}

			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpEntity)
			}
			f.Close()
		}

		if tmpModel, err := r.Response.ParseTpl("vm/go/model.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/app/model/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, "_model.go"}, "")
			if gfile.Exists(fileName) {
				gfile.Remove(fileName)
			}

			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(tmpModel)
			}
			f.Close()
		}

		if tmpExtend, err := r.Response.ParseTpl("vm/go/extend.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/app/model/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, ".go"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpExtend)
				}
				f.Close()
			}
		}

		if tmpService, err := r.Response.ParseTpl("vm/go/service.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/app/service/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, "_service.go"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpService)
				}
				f.Close()
			}
		}

		if tmpRouter, err := r.Response.ParseTpl("vm/go/router.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/app/controller/", entity.ModuleName, "/", entity.BusinessName, "_router.go"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpRouter)
				}
				f.Close()
			}
		}

		if tmpController, err := r.Response.ParseTpl("vm/go/controller.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/app/controller/", entity.ModuleName, "/", entity.BusinessName, "/", entity.BusinessName, "_controller.go"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpController)
				}
				f.Close()
			}
		}

		if tmpSql, err := r.Response.ParseTpl("vm/sql/sql.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/document/sql/", entity.ModuleName, "/", entity.BusinessName, "_menu.sql"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpSql)
				}
				f.Close()
			}
		}

		if tmpVo, err := r.Response.ParseTpl("vm/go/vo.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/app/vo/", entity.ModuleName, "/", entity.BusinessName, "vo/", entity.BusinessName, "_vo.go"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpVo)
				}
				f.Close()
			}
		}

		if tmpDto, err := r.Response.ParseTpl("vm/go/dto.html", g.Map{"table": entity}); err == nil {
			fileName := strings.Join([]string{curDir, "/app/dto/", entity.ModuleName, "/", entity.BusinessName, "dto/", entity.BusinessName, "_vo.go"}, "")
			if !gfile.Exists(fileName) {
				f, err := gfile.Create(fileName)
				if err == nil {
					f.WriteString(tmpDto)
				}
				f.Close()
			}
		}
	}

	c.Succ(r)
}

//查询数据库列表
func (c *Controller) DataList(r *ghttp.Request) {
	var req *tableModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.Err(r,err.Error())
	}
	result, page, err := tableService.SelectDbTableList(req)

	if err != nil {
		c.Err(r,err.Error())
	}

	c.SuccTable(r,page.Total,result)
}

