package model

import (
	"time"

	"github.com/zhangyiming748/MultiTranslatorUnifier/storage"
)

type TranslateHistory struct {
	Id   int64  `xorm:"pk autoincr notnull comment('主键id') INT(11)"`
	Src  string `xorm:"varchar(255) comment(原文)"`
	Dst  string `xorm:"varchar(255) comment(译文)"`
	From string `xorm:"varchar(255) comment(来源)"`
	//Source_lang string    `xorm:"varchar(255) comment(源语言)"`
	//Target_lang string    `xorm:"varchar(255) comment(目标语言)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (t *TranslateHistory) InsertOne() (int64, error) {
	storage.GetMysql()
	return storage.GetMysql().InsertOne(t)
}

func (t *TranslateHistory) FindBySrc() (bool, error) {
	return storage.GetMysql().Where("src = ?", t.Src).Get(t)
}

func (t *TranslateHistory) InsertAll(histories []TranslateHistory) (int64, error) {
	return storage.GetMysql().Insert(histories)
}
