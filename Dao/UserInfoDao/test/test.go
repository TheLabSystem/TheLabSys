package test

//package main

import (
	"TheLabSystem/Dao/UserInfoDao"
	"fmt"
)

func main() {
	fmt.Println(UserInfoDao.ChangeUserInfo(3, "hello world,I am OceanCT"))
	fmt.Println(UserInfoDao.FindUserInfoByID(3))
}
