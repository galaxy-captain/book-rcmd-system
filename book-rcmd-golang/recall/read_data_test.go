package recall_test

import "testing"
import "book-rcmd/recall"

func TestInit(t *testing.T) {
	if recall.DB == nil {
		t.Error("DB not except nil.")
	} else {
		t.Log("DB is init succeed.")
	}
}
