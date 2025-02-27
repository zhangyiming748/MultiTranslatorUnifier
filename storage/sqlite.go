package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var	db *sql.DB
var err error


// NewSQLiteStorage 创建新的 SQLite 存储实例
func NewSQLiteStorage(dbPath string)  {
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("Open %s failed: %v", dbPath, err)
	}

	// 创建翻译记录表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS translations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			from_source TEXT NOT NULL,
			src TEXT NOT NULL,
			dst TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Printf("Create table failed: %v", err)
	}
}

func GetSqlite() *sql.DB {
	return db
}
// SaveTranslation 保存翻译记录
func SaveTranslation(from, src, dst string) error {
	_, err := db.Exec(
		"INSERT INTO translations (from_source, src, dst) VALUES (?, ?, ?)",
		from, src, dst,
	)
	log.Printf("SaveTranslation %s %s %s", from, src, dst)
	return err
}

// GetTranslation 根据原文查询译文
func  GetTranslation(src string) (string, error) {
	var dst string
	err := db.QueryRow(
		"SELECT dst FROM translations WHERE src = ? ORDER BY created_at DESC LIMIT 1",
		src,
	).Scan(&dst)
	if err == sql.ErrNoRows {
		log.Printf("GetTranslation %s not found", src)
		return "", nil
	}
	if err!= nil {
		log.Printf("GetTranslation %s failed: %v", src, err)
		return "", err
	}
	return dst, err
}

// Close 关闭数据库连接
func Close() {
	db.Close()
}