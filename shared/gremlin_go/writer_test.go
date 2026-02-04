package gremlin

import "testing"

func TestSizeBool(t *testing.T) {
	buf := NewWriter(2)
	buf.AppendBool(1, true)
	if len(buf.Bytes()) != 2 {
		t.Errorf("Expected size 2, got %d", buf.Bytes())
	}
}
