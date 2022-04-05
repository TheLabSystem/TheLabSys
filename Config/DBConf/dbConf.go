package DBConf

//package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DataBase struct {
	DbType   string `yaml:"type"`
	DbName   string `yaml:"name"`
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type Config struct {
	DataBase DataBase `yaml:"database"`
}

var db *DataBase

func getConf() error {
	data, err := ioutil.ReadFile("./Config/config.yaml")
	if err != nil {
		fmt.Println("Error happened when reading config file in function DBConf.getConf()")
		return err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error happened when parsing config file in function DBConf.getConf()")
		return err
	}
	db = &config.DataBase
	return nil
}

func GetDsn() (string, error) {
	err := getConf()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.IP, db.Port, db.DbName), nil
}

//func main() {
//	dsn, err := GetDsn()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(dsn)
//}
