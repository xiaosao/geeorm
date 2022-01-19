package session

import (
	"database/sql"
	"geeorm/log"
	"strings"
)

// Session 负责与数据库的交互

type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

// New
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 统一打印日志、执行完成后进行重置，这样 Session 可以复用，一个会话可以执行多次 SQL
// Exec
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.db.Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.db.QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.db.Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
