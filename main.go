package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"mysql2excel/service"
)

type DbConfig struct {
	Db     string `yaml:"connectstr"`
	File   string `yaml:"file"`
	DbName string `yaml:"dbname"`
}

func main() {

	dbconfig := new(DbConfig)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	err = yaml.Unmarshal(yamlFile, dbconfig)
	if err != nil {
		fmt.Println("Unmarshal: %v", err)
	}
	objs := []*service.TableModel{}
	database := service.GetDB(dbconfig.Db)
	sql := "SELECT \nf.TABLE_NAME TableName,\nf.TABLE_COMMENT TableComment,\nt.COLUMN_NAME ColumnName,\nt.COLUMN_COMMENT ColumnComment,\nt.COLUMN_TYPE ColumnType,\nt.COLUMN_KEY IsPrimary,\nt.IS_NULLABLE CanNull\nFROM(\nSELECT\nsch.TABLE_SCHEMA,\nsch.TABLE_NAME,\nsch.TABLE_COMMENT\nFROM information_schema.`TABLES` sch\nWHERE TABLE_SCHEMA='" + dbconfig.DbName + "'\n) f\nLEFT JOIN information_schema.COLUMNS t ON f.TABLE_SCHEMA=t.TABLE_SCHEMA AND f.TABLE_NAME=t.TABLE_NAME"
	err = database.Select(&objs, sql)
	if err != nil {
		fmt.Println(err)
	}
	service.SaveExcel(dbconfig.File, objs)
}
