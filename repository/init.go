package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func init() {
	GetDB()
	initTable()
}

func GetDB() *gorm.DB {
	once.Do(func() {
		//change your dsn
		mysqlDsn := "root:zz19980722@tcp(49.232.87.168:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(mysqlDsn))
		DB = db
		if err != nil {
			panic("mysql init err")
		}
	})
	return DB
}

func initTable() {
	err := DB.AutoMigrate(model.User{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(model.Video{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(model.Comment{})
	if err != nil {
		return
	}
	//第一次如果上面报错，把上面注释掉，尝试下面建表逻辑。
	//err := DB.Migrator().CreateTable(model.User{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = DB.Migrator().CreateTable(model.Video{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = DB.Migrator().CreateTable(model.Comment{})
	//if err != nil {
	//	log.Fatal(err)
	//}
}
