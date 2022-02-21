package semanticSort

import (
	"fmt"
	"strings"
	"testing"
)

func TestInitialize(t *testing.T) {

	var inputSemanticVersions = [5]string{"10.0.1", "5.2.0", "1.1.1", "1.3.0", "1.2.11"}

	initializedSemanticVersions, err := Initialize(inputSemanticVersions[:])

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(initializedSemanticVersions) != 5 {
		t.Log("Length should be 5, but length is " + fmt.Sprint(len(initializedSemanticVersions)))
		t.Fail()
	}
}

/*
depends on initialize
*/
func TestSemanticSort(t *testing.T) {

	var inputSemanticVersions = [5]string{"10.0.1", "5.2.0", "1.1.1", "1.3.0", "1.2.11"}

	initializedSemanticVersions, err := Initialize(inputSemanticVersions[:])

	if err != nil {
		t.Log("Test dependedncy failed: " + err.Error())
		t.Fail()
	}

	var expectSemanticVersions = [5]string{"1.1.1", "1.2.11", "1.3.0", "5.2.0", "10.0.1"}

	//var sortedSemanticVersions = [5]string

	var sortedSemanticVersions = SemanticSort(initializedSemanticVersions)

	fmt.Println(fmt.Sprint(len(sortedSemanticVersions)))

	// check each string
	for i := 0; i < len(expectSemanticVersions); i++ {
		matches := strings.Compare(expectSemanticVersions[i], sortedSemanticVersions[i].Version)

		if matches == -1 {
			t.Log("expectSemanticVersions " + fmt.Sprint(i) + " not expected")
			t.Fail()
		}
	}

}
