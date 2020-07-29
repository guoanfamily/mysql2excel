// generated by gcg
package service

import "github.com/guoanfamily/sqlx"

import (
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbInfo *sqlx.DB
)

func GetDB(dbstr string) *sqlx.DB {
	if dbInfo != nil {
		return dbInfo
	}

	//sqlx.OpenLog()
	dbInfo, err := sqlx.Connect("mysql", dbstr+"/information_schema?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	dbInfo.SetMaxIdleConns(0)
	dbInfo.SetMaxOpenConns(100)
	dbInfo.Ping()
	return dbInfo
}

func CloseDB() {
	if dbInfo != nil {
		dbInfo.DB.Close()
	}
	dbInfo = nil
}