package simplelog

import (
	"testing"
)

func TestLog(t *testing.T) {

	var filename = "logTest.txt"

	_, err := Config(filename)

	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
