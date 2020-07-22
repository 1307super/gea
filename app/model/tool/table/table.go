package table

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"gea/app/model/tool/table_column"
	"gea/app/utils/page"
)

// Fill with you ideas below.

// Entity is the golang structure for table gen_table.
type EntityExtend struct {
	Entity
	TreeCode       string                `json:"treeCode"`       // 树编码字段
	TreeParentCode string                `json:"treeParentCode"` // 树父编码字段
	TreeName       string                `json:"treeName"`       // 树名称字段
	Columns        []table_column.Entity `json:"columns"`        // 表列信息
	PkColumn       table_column.Entity   `json:"pkColumn"`       // 表列信息
}

type Params struct {
	TreeCode       string `p:"treeCode"`
	TreeParentCode string `p:"treeParentCode"`
	TreeName       string `p:"treeName"`
}

//修改页面请求参数
type EditReq struct {
	TableId        int64  `p:"tableId" v:"required#主键ID不能为空"`
	TableName      string `p:"tableName"  v:"required#表名称不能为空"`
	TableComment   string `p:"tableComment"  v:"required#表描述不能为空"`
	ClassName      string `p:"className" v:"required#实体类名称不能为空"`
	FunctionAuthor string `p:"functionAuthor"  v:"required#作者不能为空"`
	TplCategory    string `p:"tplCategory"`
	PackageName    string `p:"packageName" v:"required#生成包路径不能为空"`
	ModuleName     string `p:"moduleName" v:"required#生成模块名不能为空"`
	BusinessName   string `p:"businessName" v:"required#生成业务名不能为空"`
	FunctionName   string `p:"functionName" v:"required#生成功能名不能为空"`
	Remark         string `p:"remark"`
	Params         string `p:"params"`
	Columns        string `p:"columns"`
}

//分页请求参数
type SelectPageReq struct {
	TableName    string `p:"tableName"`    //表名称
	TableComment string `p:"tableComment"` //表描述
	BeginTime    string `p:"beginTime"`    //开始时间
	EndTime      string `p:"endTime"`      //结束时间
	PageNum      int    `p:"pageNum"`      //当前页码
	PageSize     int    `p:"pageSize"`     //每页数
}

//根据ID获取记录
func SelectRecordById(id int64) (*EntityExtend, error) {
	db, err := gdb.Instance()
	var result EntityExtend
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("gen_table").Where("table_id=?", id)
	model.Struct(&result)

	//表数据列
	columModel := db.Table("gen_table_column").Where("table_id=?", id)

	var columList []table_column.Entity
	columModel.Structs(&columList)

	if err != nil {
		return nil, err
	}
	result.Columns = columList
	return &result, nil
}

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("gen_table t")

	if param != nil {
		if param.TableName != "" {
			model.Where("t.table_name like ?", "%"+param.TableName+"%")
		}

		if param.TableComment != "" {
			model.Where("t.table_comment like ?", "%"+param.TableComment+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Limit(page.StartNum, page.Pagesize)

	var result []Entity
	model.Structs(&result)
	return result, page, nil
}

//查询据库列表
func SelectDbTableList(param *SelectPageReq) ([]Entity, *page.Paging, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("information_schema.tables")
	model.Where("table_schema = (select database())")
	model.Where("table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%'")
	model.Where("table_name NOT IN (select table_name from gen_table)")
	if param != nil {
		if param.TableName != "" {
			model.Where("lower(table_name) like lower(?)", "%"+param.TableName+"%")
		}

		if param.TableComment != "" {
			model.Where("lower(table_comment) like lower(?)", "%"+param.TableComment+"%")
		}

		if param.BeginTime != "" {
			model.Where("date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := model.Count()

	if err != nil {
		return nil, nil, gerror.New("读取行数失败")
	}

	page := page.CreatePaging(param.PageNum, param.PageSize, total)

	model.Fields("table_name, table_comment, create_time, update_time")
	model.Limit(page.StartNum, page.Pagesize)

	var result []Entity
	model.Structs(&result)
	return result, page, nil
}

//查询据库列表
func SelectDbTableListByNames(tableNames []string) ([]Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	//tableNameStr := ""
	//
	//if len(tableNames) > 0 {
	//	for _, str := range tableNames {
	//		if str != "" {
	//			if tableNameStr == "" {
	//				tableNameStr =  str
	//			} else {
	//				tableNameStr = tableNameStr + "," + str
	//			}
	//		}
	//	}
	//}

	model := db.Table("information_schema.tables")
	model.Where("table_name NOT LIKE 'qrtz_%'")
	model.Where("table_name NOT LIKE 'gen_%'")
	model.Where("table_schema = (select database())")
	if len(tableNames) > 0 {
		model.Where("table_name in (?)", tableNames)
	}

	var result []Entity
	err = model.Structs(&result)
	return result, err
}

//查询据库列表
func SelectTableByName(tableName string) (*Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("information_schema.tables")
	model.Where("table_comment <> ''")
	model.Where("table_schema = (select database())")
	if tableName != "" {
		model.Where("table_name = ?", tableName)
	}

	var result Entity
	err = model.Struct(&result)
	return &result, err
}

//查询表ID业务信息
func SelectGenTableById(tableId int64) (*Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("gen_table t")
	model.LeftJoin("gen_table_column c", "t.table_id = c.table_id")
	model.Where("t.table_id = ?", tableId)
	model.Fields("t.table_id, t.table_name, t.table_comment, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.options, t.remark,c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort")

	var result Entity
	err = model.Struct(&result)
	return &result, err
}

func SelectableExtendById(tableId int64) (*EntityExtend, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("gen_table t")
	model.Where("t.table_id = ?", tableId)

	var result EntityExtend
	err = model.Struct(&result)
	return &result, err
}

//查询表名称业务信息
func SelectGenTableByName(tableName string) (*Entity, error) {
	db, err := gdb.Instance()

	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}

	model := db.Table("gen_table t")
	model.LeftJoin("gen_table_column c", "t.table_id = c.table_id")
	model.Where("t.table_name = ?", tableName)
	model.Fields("t.table_id, t.table_name, t.table_comment, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.options, t.remark,c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort")

	var result Entity
	err = model.Struct(&result)
	return &result, err
}
