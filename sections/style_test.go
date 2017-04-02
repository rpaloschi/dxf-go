package sections

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func styleFromDxfFragment(fragment string) *Style {
	next := core.Tagger(strings.NewReader(fragment))
	tags := core.TagSlice(core.AllTags(next))
	return NewStyle(tags)
}

func TestStyleDefaultValues(t *testing.T) {
	style := styleFromDxfFragment("")

	assert.Equal(t, "", style.Name)
	assert.InDelta(t, 1.0, style.Height, 0.001)
	assert.InDelta(t, 1.0, style.Width, 0.001)
	assert.InDelta(t, 0.0, style.Oblique, 0.001)
	assert.False(t, style.IsBackwards)
	assert.False(t, style.IsUpsideDown)
	assert.False(t, style.IsShape)
	assert.False(t, style.IsVerticalText)
	assert.Equal(t, "", style.Font)
	assert.Equal(t, "", style.BigFont)
}

const dxfStyle = `  2
STANDARD
 70
     5
 40
3.55
 41
1.1
 50
6.0
 71
     6
 42
0.2
  3
txt
  4
Arial.ttf
`

func TestDxfStyle(t *testing.T) {
	style := styleFromDxfFragment(dxfStyle)

	assert.Equal(t, "STANDARD", style.Name)
	assert.InDelta(t, 3.55, style.Height, 0.001)
	assert.InDelta(t, 1.1, style.Width, 0.001)
	assert.InDelta(t, 6.0, style.Oblique, 0.001)
	assert.True(t, style.IsBackwards)
	assert.True(t, style.IsUpsideDown)
	assert.True(t, style.IsShape)
	assert.True(t, style.IsVerticalText)
	assert.Equal(t, "txt", style.Font)
	assert.Equal(t, "Arial.ttf", style.BigFont)
}
