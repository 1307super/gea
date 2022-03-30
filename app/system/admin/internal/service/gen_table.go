package service

import (
	"bytes"
	"context"
	"fmt"
	"gea/app/dao"
	"gea/app/model"
	"gea/app/shared"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/convert"
	"gea/app/utils/page"
	"gea/app/utils/zip"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gparser"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gview"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

var GenTable = &genTableService{}

type genTableService struct{}

//根据主键查询数据
func (s *genTableService) Info(id int64) (*model.GenTableExtend, error) {
	var result *model.GenTableExtend
	if err := dao.GenTable.Where(dao.GenTable.Columns.TableId, id).Struct(&result); err != nil {
		return nil, gerror.New("获取数据失败")
	}
	//表数据列
	if err := dao.GenTableColumn.Where(dao.GenTableColumn.Columns.TableId, id).Structs(&result.Columns); err != nil {
		return nil, gerror.New("获取数据失败")
	}

	//表附加属性
	s.SetTableFromOptions(result)
	return result, nil
}

// 预览
func (s *genTableService) Preview(tableId int64) g.Map {
	entity, err := s.Info(tableId)
	if err != nil || entity == nil {
		return nil
	}
	s.SetPkColumn(entity, entity.Columns)
	listKey := "vm/vue/index.vue.vm"
	listValue := ""
	listTmp := "vm/vue/index.html"
	appJsKey := "vm/js/api.js.vm"
	appJsValue := ""
	appJsTmp := "vm/js/api.html"
	//treeKey := "vm/vue/tree.html.vm"
	//treeValue := ""
	if entity.TplCategory == "tree" {
		listTmp = "vm/vue/index-tree.html"
	}
	sqlKey := "vm/sql/menu.sql.vm"
	sqlValue := ""
	modelKey := "vm/go/model.go.vm"
	modelValue := ""
	modelInternalKey := "vm/go/model_internal.go.vm"
	modelInternalValue := ""
	daoKey := "vm/go/dao.go.vm"
	daoValue := ""
	daolInternalKey := "vm/go/dao_internal.go.vm"
	daolInternalValue := ""
	controllerKey := "vm/go/controller.go.vm"
	controllerValue := ""
	defineKey := "vm/go/define.go.vm"
	defineValue := ""
	serviceKey := "vm/go/service.go.vm"
	serviceValue := ""
	view := g.View()
	if tmpList, err := view.Parse(context.TODO(), listTmp, g.Map{"table": entity}); err == nil {
		listValue = tmpList
	}

	//if entity.TplCategory == "tree" {
	//	if tmpTree, err := view.Parse(context.TODO(),"vm/vue/index-tree.html", g.Map{"table": entity}); err == nil {
	//		treeValue = tmpTree
	//	}
	//}

	if tmpAppJs, err := view.Parse(context.TODO(), appJsTmp, g.Map{"table": entity}); err == nil {
		appJsValue = tmpAppJs
	}

	if tmpModel, err := view.Parse(context.TODO(), "vm/go/model.html", g.Map{"table": entity}); err == nil {
		modelValue = tmpModel
	}

	if tmpModelInternal, err := view.Parse(context.TODO(), "vm/go/model_internal.html", g.Map{"table": entity}); err == nil {
		modelInternalValue = tmpModelInternal
	}

	if tmpDao, err := view.Parse(context.TODO(), "vm/go/dao.html", g.Map{"table": entity}); err == nil {
		daoValue = tmpDao
	}
	if tmpDaoInternal, err := view.Parse(context.TODO(), "vm/go/dao_internal.html", g.Map{"table": entity}); err == nil {
		daolInternalValue = tmpDaoInternal
	}

	if tmpService, err := view.Parse(context.TODO(), "vm/go/service.html", g.Map{"table": entity}); err == nil {
		serviceValue = tmpService
	}

	if tmpController, err := view.Parse(context.TODO(), "vm/go/controller.html", g.Map{"table": entity}); err == nil {
		controllerValue = tmpController
	}
	if tmpDefine, err := view.Parse(context.TODO(), "vm/go/define.html", g.Map{"table": entity}); err == nil {
		defineValue = tmpDefine
	}

	if tmpSql, err := view.Parse(context.TODO(), "vm/sql/sql.html", g.Map{"table": entity}); err == nil {
		sqlValue = tmpSql
	}

	if entity.TplCategory == "tree" {
		return g.Map{
			listKey:  listValue,
			appJsKey: appJsValue,
			//treeKey:          treeValue,
			sqlKey:           sqlValue,
			modelKey:         modelValue,
			modelInternalKey: modelInternalValue,
			daoKey:           daoValue,
			daolInternalKey:  daolInternalValue,
			serviceKey:       serviceValue,
			controllerKey:    controllerValue,
			defineKey:        defineValue,
		}
	} else {
		return g.Map{
			listKey:          listValue,
			appJsKey:         appJsValue,
			sqlKey:           sqlValue,
			modelKey:         modelValue,
			modelInternalKey: modelInternalValue,
			daoKey:           daoValue,
			daolInternalKey:  daolInternalValue,
			serviceKey:       serviceValue,
			controllerKey:    controllerValue,
			defineKey:        defineValue,
		}
	}
}

//根据条件分页查询数据
func (s *genTableService) GetList(param *define.GenTableApiSelectPageReq) *define.GenTableServiceList {
	m := dao.GenTable.As("t")
	if param != nil {
		if param.TableName != "" {
			m = m.Where("t.table_name like ?", "%"+param.TableName+"%")
		}

		if param.TableComment != "" {
			m = m.Where("t.table_comment like ?", "%"+param.TableComment+"%")
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	total, err := m.Count()
	if err != nil {
		return nil
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)
	m = m.Limit(page.StartNum, page.Pagesize)
	result := &define.GenTableServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

// 添加
func (s *genTableService) Create(ctx context.Context, tables string) error {
	user := shared.Context.Get(ctx).User
	operName := user.UserExtend.LoginName
	tableArr := strings.Split(tables, ",")
	tableList, err := s.GetAllByName(tableArr)
	if err != nil {
		return gerror.New("获取表失败")
	}

	if tableList == nil {
		return gerror.New("请选择需要导入的表")
	}

	return s.ImportGenTable(tableList, operName)
}

//修改数据
func (s *genTableService) Update(ctx context.Context, req *define.GenTableApiUpdateReq) (int64, error) {
	user := shared.Context.Get(ctx).User
	if req == nil {
		return 0, gerror.New("参数错误")
	}
	table, err := dao.GenTable.FindOne(dao.GenTable.Columns.TableId, req.TableId)
	if err != nil || table == nil {
		return 0, gerror.New("数据不存在")
	}
	if req.TableName != "" {
		table.TableName = req.TableName
	}
	if req.TableComment != "" {
		table.TableComment = req.TableComment
	}
	if req.BusinessName != "" {
		table.BusinessName = req.BusinessName
	}
	if req.ClassName != "" {
		table.ClassName = req.ClassName
	}
	if req.FunctionAuthor != "" {
		table.FunctionAuthor = req.FunctionAuthor
	}
	if req.FunctionName != "" {
		table.FunctionName = req.FunctionName
	}
	if req.ModuleName != "" {
		table.ModuleName = req.ModuleName
	}
	if req.PackageName != "" {
		table.PackageName = req.PackageName
	}
	if req.Remark != "" {
		table.Remark = req.Remark
	}
	if req.TplCategory != "" {
		table.TplCategory = req.TplCategory
	}
	if req.Params != "" {
		table.Options = req.Params
	}
	table.UpdateBy = user.UserExtend.LoginName
	table.UpdateTime = gtime.Now()

	tx, err := g.DB().Begin()
	if err != nil {
		return 0, err
	}

	_, err = dao.GenTable.TX(tx).Data(table).Save()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	//保存列数据
	if req.Columns != "" {
		if j, err := gjson.DecodeToJson([]byte(req.Columns)); err != nil {
			glog.Error(err)
		} else {
			var columnList []model.GenTableColumn
			err = j.Structs(&columnList)
			if err == nil && columnList != nil && len(columnList) > 0 {
				for _, column := range columnList {
					if column.ColumnId > 0 {
						tmp, _ := dao.GenTableColumn.FindOne(dao.GenTableColumn.Columns.ColumnId, column.ColumnId)
						if tmp != nil {
							tmp.ColumnComment = column.ColumnComment
							tmp.GoType = column.GoType
							tmp.HtmlType = column.HtmlType
							tmp.QueryType = column.QueryType
							tmp.GoField = column.GoField
							tmp.DictType = column.DictType
							tmp.IsInsert = column.IsInsert
							tmp.IsEdit = column.IsEdit
							tmp.IsList = column.IsList
							tmp.IsQuery = column.IsQuery
							_, err = dao.GenTableColumn.TX(tx).Data(tmp).Save()
							if err != nil {
								tx.Rollback()
								return 0, err
							}
						}
					}
				}
			}
		}
	}
	return 1, tx.Commit()
}

//批量删除数据记录
func (s *genTableService) Delete(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dao.GenTable.Delete("table_id in (?)", idarr)
	if err != nil {
		return 0
	}
	nums, _ := result.RowsAffected()
	if nums > 0 {
		dao.GenTableColumn.Delete("table_id in (?)", idarr)
	}
	return nums
}

//查询据库列表
func (s *genTableService) GetTables(param *define.GenTableApiSelectPageReq) *define.GenTableServiceList {
	db, err := gdb.Instance()
	if err != nil {
		return nil
	}
	var whereSlice []string
	whereSlice = append(whereSlice, "table_schema = (select database())")
	whereSlice = append(whereSlice, "table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%'")
	whereSlice = append(whereSlice, "table_name NOT IN (select table_name from gen_table)")

	if param != nil {
		if param.TableName != "" {
			whereSlice = append(whereSlice, fmt.Sprintf("lower(table_name) like lower('%s')", "%"+param.TableName+"%"))
		}
		if param.TableComment != "" {
			whereSlice = append(whereSlice, fmt.Sprintf("lower(table_comment) like lower('%s')", "%"+param.TableComment+"%"))
		}
		if param.BeginTime != "" {
			whereSlice = append(whereSlice, "date_format(create_time,'%y%m%d') >= date_format('"+param.BeginTime+"','%y%m%d') ")
		}
		if param.EndTime != "" {
			whereSlice = append(whereSlice, "date_format(create_time,'%y%m%d') <= date_format('"+param.EndTime+"','%y%m%d') ")
		}
	}
	where := gstr.Implode(" and ", whereSlice)
	countSql := fmt.Sprintf("select count(*) from information_schema.tables where %s ", where)
	fmt.Println(countSql)
	total, err := db.GetCount(countSql)
	if err != nil {
		return nil
	}
	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	listSql := fmt.Sprintf("select table_name, table_comment, create_time, update_time from information_schema.tables where %s limit ?,?", where)
	var result = &define.GenTableServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	rows, err := db.GetAll(listSql, g.Slice{page.StartNum, page.Pagesize})
	if err != nil {
		return nil
	}
	if err = rows.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

//根据名称查询据库列表
func (s *genTableService) GetAllByName(tableNames []string) ([]model.GenTable, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}
	var whereSlice []string
	whereSlice = append(whereSlice, "table_name NOT LIKE 'qrtz_%'")
	whereSlice = append(whereSlice, "table_name NOT LIKE 'gen_%'")
	whereSlice = append(whereSlice, "table_schema = (select database())")
	if len(tableNames) > 0 {
		whereSlice = append(whereSlice, fmt.Sprintf("table_name in ('%s')", gstr.Implode("','", tableNames)))
	}
	sql := fmt.Sprintf("select * from information_schema.tables where %s", gstr.Implode(" and ", whereSlice))
	rows, err := db.GetAll(sql)
	if err != nil {
		return nil, gerror.New("未查询到数据")
	}

	var result []model.GenTable
	err = rows.Structs(&result)
	return result, err
}

//导入表结构
func (s *genTableService) ImportGenTable(tableList []model.GenTable, operName string) error {
	if tableList != nil && operName != "" {
		tx, err := g.DB().Begin()
		if err != nil {
			return err
		}

		for _, table := range tableList {
			tableName := table.TableName
			s.InitTable(&table, operName)
			result, err := dao.GenTable.TX(tx).Data(table).Insert()
			if err != nil {
				return err
			}
			tmpid, err := result.LastInsertId()
			if err != nil || tmpid <= 0 {
				tx.Rollback()
				return gerror.New("保存数据失败")
			}
			table.TableId = tmpid
			// 保存列信息
			genTableColumns, err := GenTableColumn.GetAllByName(tableName)
			if err != nil || len(genTableColumns) <= 0 {
				tx.Rollback()
				return gerror.New("获取列数据失败")
			}

			for _, column := range genTableColumns {
				s.InitColumnField(&column, &table)
				_, err = dao.GenTableColumn.TX(tx).Data(column).Insert()
				if err != nil {
					tx.Rollback()
					return gerror.New("保存列数据失败")
				}
			}
		}
		return tx.Commit()
	} else {
		return gerror.New("参数错误")
	}
}

//初始化表信息
func (s *genTableService) InitTable(table *model.GenTable, operName string) {
	table.ClassName = s.ConvertClassName(table.TableName)
	table.PackageName = g.Cfg().GetString("gen.packageName")
	table.ModuleName = g.Cfg().GetString("gen.moduleName")
	table.BusinessName = s.GetBusinessName(table.TableName)
	table.FunctionName = strings.ReplaceAll(table.TableComment, "表", "")
	table.FunctionAuthor = g.Cfg().GetString("gen.author")
	table.CreateBy = operName
	table.TplCategory = "crud"
	table.CreateTime = gtime.Now()
}

//初始化列属性字段
func (s *genTableService) InitColumnField(column *model.GenTableColumn, table *model.GenTable) {
	dataType := s.GetDbType(column.ColumnType)
	columnName := column.ColumnName
	column.TableId = table.TableId
	column.CreateBy = table.CreateBy
	//设置字段名
	column.GoField = s.ConvertToCamelCase(columnName)
	column.HtmlField = s.ConvertToCamelCase1(columnName)

	if GenTableColumn.IsStringObject(dataType) {
		//字段为字符串类型
		column.GoType = "string"
		if strings.EqualFold(dataType, "text") || strings.EqualFold(dataType, "tinytext") || strings.EqualFold(dataType, "mediumtext") || strings.EqualFold(dataType, "longtext") {
			column.HtmlType = "textarea"
		} else {
			columnLength := s.GetColumnLength(column.ColumnType)
			if columnLength >= 500 {
				column.HtmlType = "textarea"
			} else {
				column.HtmlType = "input"
			}
		}
	} else if GenTableColumn.IsTimeObject(dataType) {
		//字段为时间类型
		column.GoType = "Time"
		column.HtmlType = "datatime"
	} else if GenTableColumn.IsNumberObject(dataType) {
		//字段为数字类型
		column.HtmlType = "input"
		// 如果是浮点型
		tmp := column.ColumnType
		if tmp == "float" || tmp == "double" {
			column.GoType = "float64"
		} else {
			start := strings.Index(tmp, "(")
			end := strings.Index(tmp, ")")
			result := tmp[start+1 : end]
			arr := strings.Split(result, ",")
			if len(arr) == 2 && gconv.Int(arr[1]) > 0 {
				column.GoType = "float64"
			} else if len(arr) == 1 && gconv.Int(arr[0]) <= 10 {
				column.GoType = "int"
			} else {
				column.GoType = "int64"
			}
		}
	}
	//新增字段
	if columnName == "create_by" || columnName == "create_time" || columnName == "update_by" || columnName == "update_time" {
		column.IsRequired = "0"
		column.IsInsert = "0"
	} else {
		column.IsRequired = "0"
		column.IsInsert = "1"
		if strings.Index(columnName, "name") >= 0 || strings.Index(columnName, "status") >= 0 {
			column.IsRequired = "1"
		}
	}

	// 编辑字段
	if GenTableColumn.IsNotEdit(columnName) {
		if column.IsPk == "1" {
			column.IsEdit = "0"
		} else {
			column.IsEdit = "1"
		}
	} else {
		column.IsEdit = "0"
	}
	// 列表字段
	if GenTableColumn.IsNotList(columnName) {
		column.IsList = "1"
	} else {
		column.IsList = "0"
	}
	// 查询字段
	if GenTableColumn.IsNotQuery(columnName) {
		column.IsQuery = "1"
	} else {
		column.IsQuery = "0"
	}

	// 查询字段类型
	if s.CheckNameColumn(columnName) {
		column.QueryType = "LIKE"
	} else {
		column.QueryType = "EQ"
	}

	// 状态字段设置单选框
	if s.CheckStatusColumn(columnName) {
		column.HtmlType = "radio"
	} else if s.CheckTypeColumn(columnName) || s.CheckSexColumn(columnName) {
		// 类型&性别字段设置下拉框
		column.HtmlType = "select"
	}
}

func (s *genTableService) GenCode(r *ghttp.Request, tableId string) error {
	tableIds := convert.ToInt64Array(tableId, ",")
	if len(tableIds) <= 0 {
		return gerror.New("参数错误")
	}
	for _, tid := range tableIds {
		entity, err := s.Info(tid)
		if err != nil || entity == nil {
			return gerror.New("数据不存在")
		}

		s.SetPkColumn(entity, entity.Columns)
		listTmp := "vm/vue/index.html"
		if entity.TplCategory == "tree" {
			listTmp = "vm/vue/index-tree.html"
		}

		template := g.MapStrStr{
			listTmp:                     strings.Join([]string{"/template/", "business", "/", entity.BusinessName, "/index.vue"}, ""),
			"vm/js/api.html":            strings.Join([]string{"/template/", "business", "/", entity.BusinessName, "/index.js"}, ""),
			"vm/go/model.html":          strings.Join([]string{"/app/model/", entity.TableName, ".go"}, ""),
			"vm/go/model_internal.html": strings.Join([]string{"/app/model/", "internal", "/", entity.TableName, ".go"}, ""),
			"vm/go/dao.html":            strings.Join([]string{"/app/dao/", entity.TableName, ".go"}, ""),
			"vm/go/dao_internal.html":   strings.Join([]string{"/app/dao/", "internal", "/", entity.TableName, ".go"}, ""),
			"vm/go/controller.html":     strings.Join([]string{"/app/system/", entity.ModuleName, "/internal/api/", entity.TableName, ".go"}, ""),
			"vm/go/define.html":         strings.Join([]string{"/app/system/", entity.ModuleName, "/internal/define/", entity.TableName, ".go"}, ""),
			"vm/go/service.html":        strings.Join([]string{"/app/system/", entity.ModuleName, "/internal/service/", entity.TableName, ".go"}, ""),
			//"vm/go/router.html":  strings.Join([]string{ "/app/controller/", entity.ModuleName, "/", entity.BusinessName, "_router.go"}, ""),
			"vm/sql/sql.html": strings.Join([]string{"/document/sql/", "business", "/", entity.BusinessName, "_menu.sql"}, ""),
		}

		view := g.View()
		buf := new(bytes.Buffer)
		zipUtil := zip.New(buf)
		for k, v := range template {
			s.genFile(view, zipUtil, &define.GenTableServiceGenFile{
				TemplatePath:   k,
				GenTableEntity: entity,
				GenFileName:    v,
			})
		}
		zipUtil.Close()
		r.Response.Header().Set("Content-Type", "application/zip")
		r.Response.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "gea.zip"))
		r.Response.WriteExit(buf.Bytes())
	}
	return nil
}

func (s *genTableService) genFile(view *gview.View, zip *zip.ZipUtils, genParam *define.GenTableServiceGenFile) {
	if tmpList, err := view.Parse(context.TODO(), genParam.TemplatePath, g.Map{"table": genParam.GenTableEntity}); err == nil {
		//生成zip
		zip.PackToBuffer(genParam.GenFileName, gconv.Bytes(tmpList))
	}
}

//检查字段名后3位是否是sex
func (s *genTableService) CheckSexColumn(columnName string) bool {
	if len(columnName) >= 3 {
		end := len(columnName)
		start := end - 3

		if start <= 0 {
			start = 0
		}

		if columnName[start:end] == "sex" {
			return true
		}
	}
	return false
}

//检查字段名后4位是否是type
func (s *genTableService) CheckTypeColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4

		if start <= 0 {
			start = 0
		}

		if columnName[start:end] == "type" {
			return true
		}
	}
	return false
}

//检查字段名后4位是否是name
func (s *genTableService) CheckNameColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4

		if start <= 0 {
			start = 0
		}

		tmp := columnName[start:end]

		if tmp == "name" {
			return true
		}
	}
	return false
}

//检查字段名后6位是否是status
func (s *genTableService) CheckStatusColumn(columnName string) bool {
	if len(columnName) >= 6 {
		end := len(columnName)
		start := end - 6

		if start <= 0 {
			start = 0
		}
		tmp := columnName[start:end]

		if tmp == "status" {
			return true
		}
	}

	return false
}

//获取数据库类型字段
func (s *genTableService) GetDbType(columnType string) string {
	if strings.Index(columnType, "(") > 0 {
		return columnType[0:strings.Index(columnType, "(")]
	} else {
		return columnType
	}
}

//表名转换成类名
func (s *genTableService) ConvertClassName(tableName string) string {
	autoRemovePre := g.Cfg().GetBool("gen.autoRemovePre")
	tablePrefix := g.Cfg().GetString("gen.tablePrefix")
	if autoRemovePre && tablePrefix != "" {
		searchList := strings.Split(tablePrefix, ",")
		for _, str := range searchList {
			tableName = strings.ReplaceAll(tableName, str, "")
		}
	}
	return tableName
}

//获取业务名
func (s *genTableService) GetBusinessName(tableName string) string {
	lastIndex := strings.LastIndex(tableName, "_")
	nameLength := len(tableName)
	businessName := tableName[lastIndex+1 : nameLength]
	return businessName
}

//将下划线大写方式命名的字符串转换为驼峰式。如果转换前的下划线大写方式命名的字符串为空，则返回空字符串。 例如：HELLO_WORLD->HelloWorld
func (s *genTableService) ConvertToCamelCase(name string) string {
	if name == "" {
		return ""
	} else if !strings.Contains(name, "_") {
		// 不含下划线，仅将首字母大写
		return strings.ToUpper(name[0:1]) + name[1:len(name)]
	}
	var result string = ""
	camels := strings.Split(name, "_")
	for index := range camels {
		if camels[index] == "" {
			continue
		}
		camel := camels[index]
		result = result + strings.ToUpper(camel[0:1]) + strings.ToLower(camel[1:len(camel)])
	}
	return result
}

////将下划线大写方式命名的字符串转换为驼峰式,首字母小写。如果转换前的下划线大写方式命名的字符串为空，则返回空字符串。 例如：HELLO_WORLD->helloWorld
func (s *genTableService) ConvertToCamelCase1(name string) string {
	if name == "" {
		return ""
	} else if !strings.Contains(name, "_") {
		// 不含下划线，原值返回
		return name
	}
	var result string = ""
	camels := strings.Split(name, "_")
	for index := range camels {
		if camels[index] == "" {
			continue
		}
		camel := camels[index]
		if result == "" {
			result = strings.ToLower(camel[0:1]) + strings.ToLower(camel[1:len(camel)])
		} else {
			result = result + strings.ToUpper(camel[0:1]) + strings.ToLower(camel[1:len(camel)])
		}
	}
	return result
}

//获取字段长度
func (s *genTableService) GetColumnLength(columnType string) int {
	start := strings.Index(columnType, "(")
	end := strings.Index(columnType, ")")
	result := columnType[start+1 : end-1]
	return gconv.Int(result)
}

//设置代码生成其他选项值
func (s *genTableService) SetTableFromOptions(entity *model.GenTableExtend) {
	if entity != nil && entity.Options != "" {
		if p, e := gparser.LoadJson([]byte(entity.Options)); e != nil {
			glog.Error(e)
		} else {
			treeCode := p.GetString("treeCode")
			treeParentCode := p.GetString("treeParentCode")
			treeName := p.GetString("treeName")
			entity.TreeCode = treeCode
			entity.TreeParentCode = treeParentCode
			entity.TreeName = treeName
		}
	}
}

//设置主键列信息
func (s *genTableService) SetPkColumn(table *model.GenTableExtend, columns []model.GenTableColumn) {
	for _, column := range columns {
		if column.IsPk == "1" {
			table.PkColumn = column
			break
		}
	}
	if &(table.PkColumn) == nil {
		table.PkColumn = columns[0]
	}
}
