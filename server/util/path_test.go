package util_test

import (
	"fmt"
	"testing"

	"github.com/he-wen-yao/my-blog/server/util"
)

func TestPath_GetCurrentAbPath(t *testing.T) {
	path := util.GetProjectRootPath()
	fmt.Println(path)
}
