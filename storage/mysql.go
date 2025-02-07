package storage

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func ConnectToMySQL() (*xorm.Engine, error) {
	// 构建数据库连接字符串, 先连接到mysql数据库, 这里不指定数据库名
	dsn := fmt.Sprintf("root:123456@(mysql:3306)/?charset=utf8mb4&parseTime=True&loc=Local")

	// 创建 XORM 引擎
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("创建 XORM 引擎失败: %w", err)
	}

	// 创建数据库
	_, err = engine.Exec("CREATE DATABASE IF NOT EXISTS translate DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
	if err != nil {
		return nil, fmt.Errorf("创建数据库失败: %w", err)
	}

	// 关闭之前的engine
	engine.Close()
	// 构建数据库连接字符串
	dsn = fmt.Sprintf("root:123456@(mysql:3306)/translate?charset=utf8mb4&parseTime=True&loc=Local")

	// 创建 XORM 引擎
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("创建 XORM 引擎失败: %w", err)
	}

	// 可选: 设置连接池参数 (根据需要调整)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	//engine.SetConnMaxLifetime(time.Hour)

	// 可选: 启用日志记录 (根据需要调整)
	//engine.ShowSQL(true)
	//engine.Logger().SetLevel(core.LOG_DEBUG)

	// 测试连接
	if err := engine.Ping(); err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	fmt.Println("数据库连接成功!")

	return engine, nil
}

func GetMysql() *xorm.Engine {
	return engine
}
