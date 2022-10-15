package models

import (
	"JoinUs/dao"
	"fmt"
)

type User struct {
	//姓名
	Name string `json:"name"`
	//学院
	College string `json:"college"`
	//专业
	Major string `json:"major"`
	//学号
	ID string `json:"id"`
	//qq号
	Qq string `json:"qq"`
	//电话号码
	Tel string `json:"tel"`
	//意向1
	Intention1 string `json:"intention1"`
	//意向2
	Intention2 string `json:"intention2"`
}

func CreateUser(user *User) error {
	return dao.DB.Create(&user).Error
}

func FindUserById(user *User) error {
	fmt.Println(dao.DB.First(&user).Value)
	return dao.DB.First(&user).Error
}
func UpdaterUser(user *User) error {
	return dao.DB.Save(&user).Error
}
