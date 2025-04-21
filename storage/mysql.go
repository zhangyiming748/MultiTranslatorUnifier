package storage

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	useMysql bool
	engine   *xorm.Engine
)

func SetMysql() {
	var err error
	user := os.Getenv("MYSQL_USER")
	if user == "" {
		user = "root"
	}
	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		password = "163453"
	}
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "192.168.2.10"
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}
	dbName := os.Getenv("MYSQL_DATABASE")
	if dbName == "" {
		dbName = "translate"
	}
	log.Printf("连接mysql使用的各种参数-> %v:%v@%v:%v\n", user, password, host, port)
	// 先连接到 MySQL 服务器（不指定数据库）
	rootDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", user, password, host, port)
	log.Printf("第一次连接数据库的参数%s\n", rootDSN)
	tempEngine, err := xorm.NewEngine("mysql", rootDSN)
	if err != nil {
		log.Printf("连接MySQL服务器失败: %v\n", err)
		useMysql = false
		return
	}
	// 修改这里：使用 tempEngine 而不是 engine
	if err = tempEngine.Ping(); err != nil {
		log.Printf("连接数据库失败: %v\n", err)
		useMysql = false
		return
	} else {
		log.Printf("成功Ping到数据库\n")
		useMysql = true
	}
	// 检查数据库是否存在
	query := fmt.Sprintf("SELECT SCHEMA_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = '%s'", dbName)
	rows, err := tempEngine.QueryString(query)
	if err != nil {
		log.Printf("查询数据库失败: %v\n", err)
		useMysql = false
		return
	}
	// 如果数据库不存在，创建它
	if len(rows) == 0 {
		create := fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci'", dbName)
		_, err = tempEngine.Exec(create)
		if err != nil {
			log.Printf("创建数据库失败: %v\n", err)
			useMysql = false
			return
		}
		log.Printf("成功创建数据库:%s\n", dbName)
	}
	// 关闭临时连接
	tempEngine.Close()
	// 连接到 dbName 数据库
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", user, password, host, port, dbName)
	log.Printf("第二次连接数据库的参数%s\n", dataSourceName)
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Printf("连接%s数据库失败: %v\n", dbName, err)
		useMysql = false
		return
	}
	log.Printf("成功连接到数据库:%s\n", dbName)
}

func GetMysql() *xorm.Engine {
	return engine
}

func UseMysql() bool {
	return useMysql
}
