package main

import (
	_ "gfast/boot"
	_ "gfast/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"

	// "database/sql"
	// "fmt"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/gin-gonic/gin"
)

// 定义一个全局对象db
// var db *sql.DB

// 定义一个初始化数据库的函数
// func initDB() (err error) {
// 	// dsn := "root:960690@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/gfast-v2-os?charset=utf8mb4"

// 	db, err = sql.Open("mysql", dsn)
// 	// 注意！！！这里不要使用:= 我们是给全局变量赋值，然后在main函数中使用全局变量db
// 	// open函数只是验证格式是否正确,不会校验账号密码是否正确,并不是创建数据库连接
// 	if err != nil {
// 		return err
// 	}

// 	// 尝试与数据库建立连接，即校验dsn是否正确
// 	err2 := db.Ping()
// 	if err2 != nil {
// 		return err2
// 	}
// 	return nil
// }

// @title       GFast
// @version     2.0
// @description GFast后台管理框架
// @schemes     http https


// import (
// 	//引入mysql的处理
//     "database/sql"
//     _"github.com/go-sql-driver/mysql"
//     "fmt"
//     //引入json格式化
// 	// "encoding/json"
// 	//引入接口的处理的包
// 	"github.com/gin-gonic/gin"
// )


func main() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	s.Run()

	// err := initDB() // 调用初始化数据库的函数
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// } else {
	// 	fmt.Println("连接成功")
	// }
	// fmt.Printf("db: %v\n", db)
}