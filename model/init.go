package model

import (
	"control/config"
	"control/controller"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var GormDB *gorm.DB
var maxConnectionTime = 5

func init() {
	//初始化 prometheus http客户端
	controller.NewPrometheusClient()

	//连接数据库
	times := 1
	for err := connectDB(); err != nil; times++ {
		if times == maxConnectionTime {
			panic(fmt.Sprint("can not connect to db after ", times, " times"))
			os.Exit(1)
			// break
		}
		log.Print("connect database with error", err, "reconnecting...")
	}

}

func reConnectDB() error {
	return connectDB()
}

func connectDB() error {
	db, err := gorm.Open(config.Mysql, config.Dbconnection+"?charset=utf8&parseTime=True") //这里的True首字母要大写！
	if err != nil {
		return err
	}
	GormDB = db
	return nil
}
