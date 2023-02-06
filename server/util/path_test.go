package util_test

import (
	"fmt"
	"github.com/he-wen-yao/my-blog/server/util"
	"testing"
)

func TestPath_GetCurrentAbPath(t *testing.T) {
	path := util.GetCurrentAbPath()
	fmt.Println(path)
}
