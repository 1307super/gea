package service

import (
	"gea/app/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
)

var GenTableColumn = &genTableColumnService{
	COLUMNTYPE_STR:       []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"},
	COLUMNTYPE_TIME:      []string{"datetime", "time", "date", "timestamp"},
	COLUMNTYPE_NUMBER:    []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"},
	COLUMNNAME_NOT_EDIT:  []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"},
	COLUMNNAME_NOT_LIST:  []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"},
	COLUMNNAME_NOT_QUERY: []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time", "remark"},
}

type genTableColumnService struct {
	//数据库字符串类型
	COLUMNTYPE_STR []string
	//数据库时间类型
	COLUMNTYPE_TIME []string
	//数据库数字类型
	COLUMNTYPE_NUMBER []string
	//页面不需要编辑字段
	COLUMNNAME_NOT_EDIT []string
	//页面不需要显示的列表字段
	COLUMNNAME_NOT_LIST []string
	//页面不需要显示的列表字段
	COLUMNNAME_NOT_QUERY []string
}

func (s *genTableColumnService) GetAllByName(tableName string) ([]model.GenTableColumn, error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("生成失败")
	}
	var result []model.GenTableColumn
	m := db.Table("information_schema.columns")
	m.Where("table_schema = (select database())")
	m.Where("table_name=?", tableName).Order("ordinal_position")
	m.Fields("column_name, (case when (is_nullable = 'no' && column_key != 'PRI') then '1' else null end) as is_required, (case when column_key = 'PRI' then '1' else '0' end) as is_pk, ordinal_position as sort, column_comment, (case when extra = 'auto_increment' then '1' else '0' end) as is_increment, column_type")
	m.Structs(&result)
	return result, nil
}

//判断string 是否存在在数组中
func (s *genTableColumnService) IsExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

//判断是否是数据库字符串类型
func (s *genTableColumnService) IsStringObject(value string) bool {
	return s.IsExistInArray(value, s.COLUMNTYPE_STR)
}

//判断是否是数据库时间类型
func (s *genTableColumnService) IsTimeObject(value string) bool {
	return s.IsExistInArray(value, s.COLUMNTYPE_TIME)
}

//判断是否是数据库数字类型
func (s *genTableColumnService) IsNumberObject(value string) bool {
	return s.IsExistInArray(value, s.COLUMNTYPE_NUMBER)
}

//页面不需要编辑字段
func (s *genTableColumnService) IsNotEdit(value string) bool {
	return !s.IsExistInArray(value, s.COLUMNNAME_NOT_EDIT)
}

//页面不需要显示的列表字段
func (s *genTableColumnService) IsNotList(value string) bool {
	return !s.IsExistInArray(value, s.COLUMNNAME_NOT_LIST)
}

//页面不需要查询字段
func (s *genTableColumnService) IsNotQuery(value string) bool {
	return !s.IsExistInArray(value, s.COLUMNNAME_NOT_QUERY)
}
