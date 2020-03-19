package util

import "testing"

func TestIsMapEmpty(t *testing.T) {
	var data map[string]interface{}

	if !IsMapEmpty(data) {
		t.FailNow()
	}
}
