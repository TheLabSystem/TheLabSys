package test

//package main

import (
	"TheLabSystem/Dao/DBAccessor"
	"fmt"
)

// test successfully
func testMysqlInit() {
	_, err := DBAccessor.MysqlInit()
	if err != nil {
		fmt.Println("Connection has been established successfully.")
		fmt.Println(err)
	}
}

//func main() {
//	testMysqlInit()
//}
