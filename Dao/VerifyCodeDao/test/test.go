//package test

package main

import (
	"TheLabSystem/Dao/VerifyCodeDao"
)

func main() {
	VerifyCodeDao.InsertVerifyCode(12345, 255)
	//VerifyCodeDao.DeleteVerifyCode(10203)
}
