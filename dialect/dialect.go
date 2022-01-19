package dialect

import "reflect"

// dialect 是用来区分不同数据的差异，例如数据类型，数据库表的创建、删除。

type Dialect interface {
	// 将 go 中的类型转换为数据库中的数据类型
	DataTypeOf(typ reflect.Value) string
	// 返回某个表是否存在的sql语句
	TableExistSQL(tableName string) (string, []interface{})
}

//
