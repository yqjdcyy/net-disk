package utils

import (
	"testing"
)

func TestParse(t *testing.T) {

	url := "http://127.0.0.1:5518/api/ppt?extension=pptx&format=0&indexes=&args=extension_pptx-id_276-type_0"
	m, err := Parse(url)
	if nil != err {
		t.Error(err.Error())
	}

	for key, val := range *m {
		t.Log(key, "=", val)
	}
}
