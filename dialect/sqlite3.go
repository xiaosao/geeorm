package dialect

import "reflect"

type sqlite3 struct{}

var _ Dialect = (*sqlite3)(nil) // 可以确保 sqlite3 完全实现了 Dialect 接口

func (s *sqlite3) DataTypeOf(typ reflect.Value) string {
	return ""
}

func (s *sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "", args
}
