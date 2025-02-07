package storage

import (
	"testing"

	"github.com/zhangyiming748/MultiTranslatorUnifier/model"
)

// go test -v -run TestConnectMysql
func TestConnectMysql(t *testing.T) {
	_, err := ConnectToMySQL()
	if err != nil {
		t.Log(err)
	}
	h := new(model.TranslateHistory)
}

// go test -v -run TestConnectRedis
func TestConnectRedis(t *testing.T) {
	err := InitRedis()
	if err != nil {
		t.Log(err)
	}
	InsertTranslationToRedis("hello", "你好", "我猜的")
}
