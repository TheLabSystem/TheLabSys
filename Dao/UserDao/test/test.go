//package test

package main

import (
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/ServiceType/User"
	"fmt"
)

// tested successfully
func testInsertUser() {
	var user = User.User{
		UserType:    4,
		Username:    "OceanCT",
		DisplayName: "OceanCT",
		Password:    "123456",
		Money:       121212.2349,
	}
	err := UserDao.InsertUser(user)
	if err != nil {
		fmt.Println(err)
	}
}

// tested successfully
func testDeleteUser() {
	var user = User.User{
		UserID: 2,
	}
	err := UserDao.DeleteUser(user)
	if err != nil {
		fmt.Println(err)
	}
}

// tested successfully
func testFindUser() {
	fmt.Println(UserDao.FindUserByUsername("OceanCT"))
	fmt.Println(UserDao.FindUserByID(3))
}

//func test UpdateUser()
func testUpdateUser() {
	user := User.User{UserID: 3, UserType: 2, Username: "sdafs"}
	fmt.Println(UserDao.UpdateUser(user))
}
func main() {
	testInsertUser()
}
