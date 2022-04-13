//package test

package main

import (
	"TheLabSystem/Dao/VerifyCodeDao"
)

func main() {
	VerifyCodeDao.InsertVerifyCode(1234, 255)
	//VerifyCodeDao.DeleteVerifyCode(10203)
}
