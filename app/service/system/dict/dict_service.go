package dict

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"time"
	dictDataModel "gea/app/model/system/dict_data"
)

//根据字典类型和字典键值查询字典数据信息
func GetDictLabel(dictType string, dictValue interface{}) string {
	result := ""
	value, err := dictDataModel.FindValue("dict_label", "dict_type = '"+dictType+"'", "dict_value = '"+gconv.String(dictValue)+"'")
	if err != nil {
		result = value.String()
	}
	return result
}

//通用的字典单选框控件  dictType 字典类别  value 默认值
func GetDictTypeRadio(dictType, name string, value interface{}) string {
	result, err := SelectDictDataByType(dictType)
	if err != nil {
		return ""
	}

	if result == nil || len(result) <= 0 {
		return ""
	}

	htmlstr := ``

	for i := range result {
		if strings.EqualFold(result[i].DictValue, gconv.String(value)) {
			htmlstr += `<div class="radio-box"><option value="` + result[i].DictValue + `">` + result[i].DictLabel + `</option>`
			htmlstr += `<input type="radio" id="` + gconv.String(result[i].DictCode) + `" name="` + name + `" value="` + result[i].DictValue + `"
                           checked="checked">
                    <label for="` + gconv.String(result[i].DictCode) + `" text="` + result[i].DictLabel + `"></label></div>`
		} else {
			htmlstr += `<div class="radio-box"><option value="` + result[i].DictValue + `">` + result[i].DictLabel + `</option>`
			htmlstr += `<input type="radio" id="` + gconv.String(result[i].DictCode) + `" name="` + name + `" value="` + result[i].DictValue + `">
                    <label for="` + gconv.String(result[i].DictCode) + `" text="` + result[i].DictLabel + `"></label></div>`
		}
	}

	htmlstr += ``
	return htmlstr
}

//通用的字典下拉框控件  字典类别   html控件id  html控件name html控件class  html控件value  html控件空值标签 是否可以多选
func GetDictTypeSelect(dictType, id, name, className, value, emptyLabel, multiple string) string {
	startT := time.Now() //计算当前时间
	result, err := SelectDictDataByType(dictType)
	if err != nil {
		return ""
	}

	if result == nil || len(result) <= 0 {
		return ""
	}

	htmlstr := `<select id="` + id + `" name="` + name + `" class="` + className + `" ` + multiple + `>`

	if emptyLabel != "" {
		htmlstr += `<option value="">` + emptyLabel + `</option>`
	}

	for i := range result {
		if strings.EqualFold(result[i].DictValue, value) {
			htmlstr += `<option selected value="` + result[i].DictValue + `">` + result[i].DictLabel + `</option>`
		} else {
			htmlstr += `<option value="` + result[i].DictValue + `">` + result[i].DictLabel + `</option>`
		}
	}

	htmlstr += `</select>`
	tc := time.Since(startT) //计算耗时
	fmt.Printf("time cost = %v\n", tc)
	return htmlstr
}

//通用的字典下拉框控件
func GetDictTypeData(dictType string) string {
	startT := time.Now() //计算当前时间
	result := make([]dictDataModel.Entity, 0)
	rs, err := SelectDictDataByType(dictType)
	if err == nil || len(rs) > 0 {
		result = rs
	}

	jsonstr := ""

	jsonbyte, err := json.Marshal(result)

	if err == nil {
		jsonstr = string(jsonbyte)
	}
	tc := time.Since(startT) //计算耗时
	fmt.Printf("time cost = %v\n", tc)
	return jsonstr
}

//根据字典类型查询字典数据
func SelectDictDataByType(dictType string) ([]dictDataModel.Entity, error) {
	var result []dictDataModel.Entity
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("获取数据库连接失败")
	}
	err = db.Table("sys_dict_data").Where("status = '0' and dict_type = ?", dictType).Order("dict_sort asc").Structs(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
