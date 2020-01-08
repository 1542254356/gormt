package model

import (
	"fmt"
	"regexp"

	"github.com/xxjwxc/gormt/data/config"
	"github.com/xxjwxc/gormt/data/view/cnf"
	"github.com/xxjwxc/public/mybigcamel"
)

// getCamelName Big Hump or Capital Letter.大驼峰或者首字母大写
func getCamelName(name string) string {
	if config.GetSingularTable() { // If the table name plural is globally disabled.如果全局禁用表名复数
		return titleCase(name)
	}

	return mybigcamel.Marshal(name)
}

// titleCase title case.首字母大写
func titleCase(name string) string {
	vv := []rune(name)
	if len(vv) > 0 {
		if bool(vv[0] >= 'a' && vv[0] <= 'z') { // title case.首字母大写
			vv[0] -= 32
		}
	}

	return string(vv)
}

// getTypeName Type acquisition filtering.类型获取过滤
func getTypeName(name string) string {
	// Precise matching first.先精确匹配
	if v, ok := cnf.TypeMysqlDicMp[name]; ok {
		return v
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for k, v := range cnf.TypeMysqlMatchMp {
		if ok, _ := regexp.MatchString(k, name); ok {
			return v
		}
	}

	panic(fmt.Sprintf("type (%v) not match in any way.maybe need to add on (https://github.com/xxjwxc/gormt/blob/master/data/view/cnf/def.go)", name))
}

func getUninStr(left, middle, right string) string {
	re := left
	if len(right) > 0 {
		re = left + middle + right
	}
	return re
}

func getGormModelElement() []ColumusInfo {
	var result []ColumusInfo
	result = append(result, ColumusInfo{
		BaseInfo: BaseInfo{Name: "id", Notes: "Primary key"},
		Type:     "int64",                                // Type.类型标记
		Index:    []KList{KList{Key: ColumusKeyPrimary}}, // index list.index列表
	})
	result = append(result, ColumusInfo{
		BaseInfo: BaseInfo{Name: "created_at", Notes: "created time"},
		Type:     "time.Time", // Type.类型标记
	})
	result = append(result, ColumusInfo{
		BaseInfo: BaseInfo{Name: "updated_at", Notes: "updated time"},
		Type:     "time.Time", // Type.类型标记
	})
	result = append(result, ColumusInfo{
		BaseInfo: BaseInfo{Name: "deleted_at", Notes: "deleted time"},
		Type:     "time.Time", // Type.类型标记
	})
	return result
}
