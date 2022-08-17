package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var dsn = "root:Password.@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=true"

var db *sql.DB

func gormInit() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	//create
	var p Product
	db.AutoMigrate(&p)

	//insert
	//db.Create(&Product{
	//	Code:  "1001",
	//	Price: 200,
	//})

	//read
	//db.First(&p, 2)
	//db.First(&p, "code=?", "D22")
	//fmt.Printf("p:%v\n", p)

	//update
	//db.First(&p, 1)
	//单值
	//db.Model(&p).Update("price", 200)
	//多值
	//db.Model(&p).Updates(Product{Price: 1000, Code: "kk"})

	//delete
	//db.First(&p, 3)
	//db.Delete(&p)

}
func insert() {

}

func main() {
	//err := initDB()
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//} else {
	//	fmt.Println("连接成功")
	//}
	gormInit()
}

func initDB() (err error) {
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	err = db.Ping()

	if err != nil {
		return err
	}

	return nil
}

type User struct {
	id       int
	username string
	password string
}

func QueryOneRowById(id int) {
	s := "select * from user_tb1 where id = ?"
	var u User
	err := db.QueryRow(s, id).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}
	fmt.Printf("%v", u)
}
func QueryAllRow() {
	s := "select * from user_tb1"
	rows, err := db.Query(s)

	var u User

	// 关闭连接
	defer rows.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for rows.Next() {
			err := rows.Scan(&u.id, &u.username, &u.password)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(u)
		}
	}
}

func insertData(username string, password string) {
	sqlStr := "insert into user_tb1(username,password) values (?,?)"
	r, err := db.Exec(sqlStr, username, password)

	if err != nil {
		println("插入失败")
		return
	}
	println("插入成功")
	i, _ := r.LastInsertId()
	fmt.Printf("i:%v\n", i)
}

func updateData(id int, username string, password string) {
	sql := "update user_tb1 set username=?,password=? where id = ?"
	r, err := db.Exec(sql, username, password, id)

	if err != nil {
		println("update fail")
		return
	}
	rows, err := r.RowsAffected()
	if err != nil {
		println("get RowsAffected fail")
		return
	}
	fmt.Printf("AffectedRows:%v\n", rows)
}

func deleteData(id int) {
	sql := "delete from user_tb1 where id = ?"
	r, err := db.Exec(sql, id)

	if err != nil {
		println("delete fail")
		return
	}
	rows, err := r.RowsAffected()
	if err != nil {
		println("get RowsAffected fail")
		return
	}
	fmt.Printf("AffectedRows:%v\n", rows)
}

func initGin() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
