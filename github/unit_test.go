package github

import (
	"github.com/OwO-Network/DeepLX/translate"
	"testing"
)

func TestAsk(t *testing.T) {
	src := "hello world"
	lx, err := translate.TranslateByDeepLX("auto", "zh", src, "", "", "")
	if err != nil {
		return
	}
	t.Logf("%+v\t%+v\n", lx, err)
}
