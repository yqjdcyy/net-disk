package utils

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var (
	ts = []string{"jpg", "png", "jpeg", "bmp"}
)

func TestContainType(t *testing.T) {

	s := time.Now()

	checkContain(t, ts[0], true)
	checkContain(t, strings.ToUpper(ts[1]), true)
	checkContain(t, ts[2]+"2", false)

	d := time.Since(s)
	fmt.Print(d)
}

func checkContain(t *testing.T, arg string, expt bool) {

	if ContainType(ts, arg) == expt {
		t.Logf("Type[%s] pass", arg)
	} else {
		t.Errorf("Type[%s] fail", arg)
	}
}
