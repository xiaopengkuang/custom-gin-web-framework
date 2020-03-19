package util

import (
	"fmt"
	"testing"
)

func TestPathExist(t *testing.T) {
	gopath, err := GetGoPathSrc()
	if err != nil {
		t.FailNow()
	}

	fmt.Println(gopath)
}
