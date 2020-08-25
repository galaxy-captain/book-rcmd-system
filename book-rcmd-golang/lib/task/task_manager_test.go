package task_test

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("test rcmd init")
}

func TestGoroutine(t *testing.T) {

	t.Run("test-condition-1", func(t *testing.T) {
		t.Logf("condition-1 is ok.")
	})

}
