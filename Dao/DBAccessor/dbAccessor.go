package DBAccessor

import (
	"TheLabSystem/Config/DBConf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit() (*gorm.DB, error) {
	dsn, err := DBConf.GetDsn()
	if err != nil {
		return nil, err
	}
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database Connection Error,Happened in Dao.DBAccessor.MysqlInit()")
		fmt.Println(err)
		return nil, err
	}
	return db, err
}
