package t

import (
	"testing"

	"github.com/zhangyiming748/MultiTranslatorUnifier/model"
	"github.com/zhangyiming748/MultiTranslatorUnifier/storage"
)

// go test -v -run TestConnectMysql
func TestConnectMysql(t *testing.T) {
	_, err := storage.ConnectToMySQL()
	if err != nil {
		t.Log(err)
	}
	h := new(model.TranslateHistory)
	h.Src = "hello"
	h.Dst = "你好"
	h.From = "脑补"
	h.InsertOne()
}

// go test -v -run TestConnectRedis
func TestConnectRedis(t *testing.T) {
	err := storage.InitRedis()
	if err != nil {
		t.Log(err)
	}
}
