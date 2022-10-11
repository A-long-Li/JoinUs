package main

import (
	"JoinUs/dao"
	"JoinUs/models"
	"JoinUs/rouers"
	"JoinUs/settings"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./list conf/config.ini")
		return
	}
	//连接数据库
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	if err := dao.InitMysql(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	//退出关闭数据库
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.User{})

	r := rouers.SetRouter()
	RunPort := ":" + strconv.Itoa(settings.Conf.Port)
	err := r.Run(RunPort)
	if err != nil {
		return
	}
}
