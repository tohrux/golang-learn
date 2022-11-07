package main

import (
	"database/sql"
	"fmt"
	cfg "wire-example/internal/config"
)

type App struct { // 最终需要的对象
	db *sql.DB // db可以自定义命名，*sql.DB 需要和 internal/db/db.go 中 NewDb 方法返回的类型相同
}

func NewApp(db *sql.DB) *App {
	return &App{
		db: db,
	}
}

func main() {
	fmt.Println(cfg.Provider)
}
