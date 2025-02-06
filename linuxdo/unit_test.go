package linuxdo

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	apikey := os.Getenv("LINUXDO")
	ret, _ := Req("hello", apikey)
	t.Logf("测试结果 : %v\n", ret)
}
