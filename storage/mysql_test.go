package storage

import (
	"testing"
)

func TestMysql(t *testing.T) {
	// 初始化数据库
	SetMysql()
	GetMysql().Sync2(History{})
	// 插入数据
	h := new(History)
	h.Src = "hello"
	h.Dst = "你好"
	h.InsertOne()
	// 查询数据
	h1 := new(History)
	h1.Src = "hello"
	h1.FindBySrc()
	t.Log(h1.Dst)
}
