package storage

import "testing"

// go test -v -run TestConnectMysql
func TestConnectMysql(t *testing.T) {
	_, err := ConnectToMySQL()
	if err != nil {
		t.Log(err)
	}
}

// go test -v -run TestConnectRedis
func TestConnectRedis(t *testing.T) {
	err := InitRedis()
	if err != nil {
		t.Log(err)
	}
}
