package sections

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func tableEntryTagsFromDxfFragment(fragment string) ([]core.TagSlice, error) {
	next := core.Tagger(strings.NewReader(fragment))
	tags := core.TagSlice(core.AllTags(next))
	return TableEntryTags(tags)
}

const minimalTable = `  0
TABLE
  0
ENDTAB
`

func TestTableEntryTagsMinimalTable(t *testing.T) {
	groups, err := tableEntryTagsFromDxfFragment(minimalTable)

	assert.Nil(t, err)
	assert.Equal(t, make([]core.TagSlice, 0), groups)
}

const invalidStartTable = `  0
WRONG
  0
ENDTAB
`

const invalidEndTable = `  0
TABLE
  0
ZZZZ
`

func TestTableEntryTagsInvalidTables(t *testing.T) {
	groups, err := tableEntryTagsFromDxfFragment(invalidStartTable)
	expectedMessage := "Invalid table. Missing TABLE AND/OR ENDTAB tags."

	assert.Equal(t, []core.TagSlice{}, groups)
	assert.Equal(t, expectedMessage, err.Error())

	groups, err = tableEntryTagsFromDxfFragment(invalidEndTable)

	assert.Equal(t, []core.TagSlice{}, groups)
	assert.Equal(t, expectedMessage, err.Error())
}

const tableTags = `  0
TABLE
  2
LAYER
  0
LAYER
  20
1.1
  0
ENDTAB
`

func TestTableEntryTagsValidTable(t *testing.T) {
	expected := []core.TagSlice{
		[]*core.Tag{
			core.NewTag(0, core.NewStringValue("LAYER")),
			core.NewTag(20, core.NewFloatValue(1.1)),
		},
	}
	groups, err := tableEntryTagsFromDxfFragment(tableTags)

	assert.Nil(t, err)
	assert.Equal(t, expected, groups)
}

const tagsToSplit = `  0
TABLE
  2
LAYER
  0
TABLE
  0
LAYER
  20
1.1
  0
TABLE
  60
1001
  0
ENDTAB
`

func TestSplitTagChunks(t *testing.T) {
	expected := []core.TagSlice{
		[]*core.Tag{
			core.NewTag(0, core.NewStringValue("TABLE")),
			core.NewTag(2, core.NewStringValue("LAYER")),
		},
		[]*core.Tag{
			core.NewTag(0, core.NewStringValue("TABLE")),
			core.NewTag(0, core.NewStringValue("LAYER")),
			core.NewTag(20, core.NewFloatValue(1.1)),
		},
		[]*core.Tag{
			core.NewTag(0, core.NewStringValue("TABLE")),
			core.NewTag(60, core.NewIntegerValue(1001)),
		},
	}
	next := core.Tagger(strings.NewReader(tagsToSplit))
	tags := core.TagSlice(core.AllTags(next))

	chunks := SplitTagChunks(tags,
		core.NewTag(0, core.NewStringValue("ENDTAB")),
		core.NewTag(0, core.NewStringValue("TABLE")))

	assert.Equal(t, expected, chunks)
}
