package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToEncoding(t *testing.T) {
	testCases := []struct {
		dxfCodePage string
		encoding    string
	}{
		{"932", "cp932"},
		{"a1257", "cp1257"},
		{"", "cp1252"},
	}

	for _, test := range testCases {
		assert.Equal(t, toEncoding(test.dxfCodePage), test.encoding)
	}
}
