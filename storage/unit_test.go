package storage
import (
	"testing"
	"fmt"
	"log"
)

func init() {
	NewSQLiteStorage("test.db")
}
func TestSet(t *testing.T) {
	// 保存翻译
	err := SaveTranslation("Google", "Hello", "你好")
	if err != nil {
		log.Printf("保存翻译失败: %v", err)
		return
	}
}

func TestGet(t *testing.T) {
	// 查询翻译
	translation, err := GetTranslation("Hello")
	if err != nil {
		log.Printf("查询翻译失败: %v", err)
		return
	}

	if translation != "" {
		fmt.Printf("翻译结果: %s\n", translation)
	} else {
		fmt.Println("未找到翻译")
	}
}
