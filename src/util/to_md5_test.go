package util

import (
	"fmt"
	"testing"
)

func TestToMd5(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(ToMd5("P@ssw0rd"))
	}

}
