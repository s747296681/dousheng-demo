package main

import (
	"github.com/RaymondCode/simple-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	//mysqlDsn := "root:zz19980722@tcp(127.0.0.1:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(mysqlDsn))
	//if err != nil {
	//	panic("mysql init err")
	//}
	//repository.DB = db

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
