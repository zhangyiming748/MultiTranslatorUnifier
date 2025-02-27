package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

// NewSQLiteStorage 创建新的 SQLite 存储实例
func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	db, err := sql.Open("/app/sqlite3", dbPath)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

// SaveTranslation 保存翻译记录
func (s *SQLiteStorage) SaveTranslation(from, src, dst string) error {
	_, err := s.db.Exec(
		"INSERT INTO translations (from_source, src, dst) VALUES (?, ?, ?)",
		from, src, dst,
	)
	return err
}

// GetTranslation 根据原文查询译文
func (s *SQLiteStorage) GetTranslation(src string) (string, error) {
	var dst string
	err := s.db.QueryRow(
		"SELECT dst FROM translations WHERE src = ? ORDER BY created_at DESC LIMIT 1",
		src,
	).Scan(&dst)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return dst, err
}

// Close 关闭数据库连接
func (s *SQLiteStorage) Close() error {
	return s.db.Close()
}