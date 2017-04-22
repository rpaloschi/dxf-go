package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrueColorRgb(t *testing.T) {
	color := TrueColor(0xA0B0C0)
	r, g, b := color.Rgb()

	assert.Equal(t, byte(0xA0), r)
	assert.Equal(t, byte(0xB0), g)
	assert.Equal(t, byte(0xC0), b)
}

func TestTrueColorRgbAcessors(t *testing.T) {
	color := TrueColor(0xA0B0C0)

	assert.Equal(t, byte(0xA0), color.R())
	assert.Equal(t, byte(0xB0), color.G())
	assert.Equal(t, byte(0xC0), color.B())
}

func TestTrueColorFromRgb(t *testing.T) {
	color := TrueColorFromRGB(0xA0, 0xB0, 0xC0)

	assert.Equal(t, uint(0xA0B0C0), uint(color))
}
