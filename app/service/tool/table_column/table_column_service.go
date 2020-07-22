package table_column

import (
	tableColumn "gea/app/model/tool/table_column"
	"gea/app/utils/convert"
)

//新增业务字段
func Insert(entity *tableColumn.Entity) (int64, error) {
	result, err := entity.Insert()
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改业务字段
func Update(entity *tableColumn.Entity) (int64, error) {
	result, err := entity.Update()
	if err != nil {
		return 0, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return nums, nil
}

//根据主键查询数据
func SelectRecordById(id int64) (*tableColumn.Entity, error) {
	return tableColumn.FindOne("column_id", id)
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	result, err := tableColumn.Delete("column_id", id)
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
	result, err := tableColumn.Delete("column_id in (?)", idarr)
	if err != nil {
		return 0
	}

	nums, _ := result.RowsAffected()

	return nums
}

//查询业务字段列表
func SelectGenTableColumnListByTableId(tableId int64) (*[]tableColumn.Entity, error) {
	return tableColumn.SelectGenTableColumnListByTableId(tableId)
}

//根据表名称查询列信息
func SelectDbTableColumnsByName(tableName string) (*[]tableColumn.Entity, error) {
	return tableColumn.SelectDbTableColumnsByName(tableName)
}
