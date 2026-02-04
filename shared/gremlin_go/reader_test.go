//go:build integration

package gremlin

import (
	"encoding/hex"
	"testing"
)

func TestDecodeInt64(t *testing.T) {
	var data = []byte{212, 1}

	buf := NewReader(data)
	if buf.ReadSInt64(0) != 106 {
		t.Errorf("Expected 106, got %d", buf.ReadSInt64(0))
	}
}

func TestVectorOfEnums(t *testing.T) {
	request := "0a03000102"
	data, err := hex.DecodeString(request)
	if err != nil {
		t.Fatalf("failed to decode hex string: %v", err)
	}

	buf := NewReader(data)
	offset := 0
	for buf.HasNext(offset, 0) {
		tag, wire, tagSize, err := buf.ReadTagAt(offset)
		if err != nil {
			t.Fatalf("failed to read tag: %v", err)
		}

		offset += tagSize
		offset, err = buf.SkipData(offset, wire)
		if err != nil {
			t.Fatalf("failed to skip data: %v", err)
		}
	}
}
