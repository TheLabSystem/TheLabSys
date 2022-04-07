//package test

package main

import (
	"TheLabSystem/Dao/VerifyCodeDao"
	"fmt"
)

func main() {
	//VerifyCodeDao.InsertVerifyCode(1, 4)
	//VerifyCodeDao.DeleteVerifyCode(10203)
	fmt.Println(VerifyCodeDao.CheckVerifyCode(1, 4))
}
