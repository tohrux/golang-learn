package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"wire-example/internal/config"
)

var Provider = wire.NewSet(NewDb) // 将New方法声明为Provider，表示New方法可以创建一个被别人依赖的对象

func NewDb(cfg *config.Config) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", cfg.Database.Dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return db, nil
}
