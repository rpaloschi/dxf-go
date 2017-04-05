package sections

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func lineTypeFromDxfFragment(fragment string) (*LineType, error) {
	next := core.Tagger(strings.NewReader(fragment))
	tags := core.TagSlice(core.AllTags(next))
	return NewLineType(tags)
}

func TestLineTypeDefaultValues(t *testing.T) {
	lineType, err := lineTypeFromDxfFragment("")

	assert.Nil(t, err)
	assert.Equal(t, "", lineType.Name)
	assert.Equal(t, "", lineType.Description)
	assert.InDelta(t, 0.0, lineType.Length, 0.001)
	assert.Equal(t, []*LineElement{}, lineType.Pattern)
}

const dxfLineType = ` 2
DASHDOT
  3
Strange dashdot
 72
    65
 73
     4
 40
1.0
 49
0.4
 74
3
 46
1.0
 50
0.5
 44
1.5
 45
2.3
  9
Sample Text
 49
0.3
 74
5
 75
1
 46
1.5
 50
1.0
 44
0.1
 45
0.2
 49
0.3
 74
4
 75
2
`

func TestParseLineType(t *testing.T) {
	lineType, err := lineTypeFromDxfFragment(dxfLineType)

	assert.Nil(t, err)
	assert.Equal(t, "DASHDOT", lineType.Name)
	assert.Equal(t, "Strange dashdot", lineType.Description)
	assert.InDelta(t, 1.0, lineType.Length, 0.001)

	expectedElements := []*LineElement{
		{0.4, true, true, false, 0,
			1.0, 0.5, 1.5, 2.3, "Sample Text"},
		{0.3, true, false, true, 1,
			1.5, 1.0, 0.1, 0.2, ""},
		{0.3, false, false, true, 2,
			1.0, 0.0, 0.0, 0.0, ""},
	}

	assert.Equal(t, expectedElements, lineType.Pattern)
}
