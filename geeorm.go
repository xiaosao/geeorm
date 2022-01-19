package geeorm

import (
	"database/sql"
	"geeorm/log"
	"geeorm/session"
)

// Engine 负责连接数据库，交互后的关闭连接

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	e = &Engine{db}
	log.Info("Connect database sucess")
	return

}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error(err)
	}
	log.Info("Close database sucess")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
