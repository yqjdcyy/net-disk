package utils

import (
	"fmt"
	"testing"
)

func TestPOST(t *testing.T) {

	p1 := "test"
	f := "2018.jpg"
	p := fmt.Sprintf("‪D:\\download\\%v\\%v", p1, f)
	u := fmt.Sprintf("http://106.15.73.253:7000/upload?path=%v&filename=%v", p1, f)

	POST(p, u)
	t.Log("POST")
}
