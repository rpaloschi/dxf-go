package sections

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rpaloschi/dxf-go/core"
	"github.com/rpaloschi/dxf-go/entities"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewBlock(t *testing.T) {
	expected := Block{
		Name:         "LINE-BLOCK",
		Handle:       "B1",
		LayerName:    "0",
		SecondName:   "LINE-BLOCK-2",
		BasePoint:    core.Point{X: 0.5, Y: 0.1, Z: 1.1},
		XrefPathName: "",
		Description:  "This is a test block.",
		Entities:     entities.EntitySlice{},
	}

	next := core.Tagger(strings.NewReader(dxfBlock))
	tags := core.TagSlice(core.AllTags(next))

	block, err := NewBlock(tags)

	assert.Equal(t, nil, err)
	assert.True(t, expected.Equals(block),
		"Expected %+v and %+v to be equals",
		spew.Sdump(expected), spew.Sdump(block))
}

func TestBlockEquality(t *testing.T) {
	testCases := []struct {
		b1     *Block
		b2     *Block
		equals bool
	}{
		{
			&Block{
				Name:         "LINE-BLOCK",
				Handle:       "B1",
				LayerName:    "0",
				SecondName:   "LINE-BLOCK-2",
				BasePoint:    core.Point{X: 0.5, Y: 0.1, Z: 1.1},
				XrefPathName: "",
				Description:  "This is a test block.",
				Entities: entities.EntitySlice{
					&entities.Vertex{Location: core.Point{X: 1.5, Y: 2.9}},
					&entities.Vertex{Location: core.Point{X: 2.5, Y: 1.9}},
				},
			},
			&Block{
				Name:         "LINE-BLOCK",
				Handle:       "B1",
				LayerName:    "0",
				SecondName:   "LINE-BLOCK-2",
				BasePoint:    core.Point{X: 0.5, Y: 0.1, Z: 1.1},
				XrefPathName: "",
				Description:  "This is a test block.",
				Entities: entities.EntitySlice{
					&entities.Vertex{Location: core.Point{X: 1.5, Y: 2.9}},
					&entities.Vertex{Location: core.Point{X: 2.5, Y: 1.9}},
				},
			},
			true,
		},

		{
			&Block{
				Name:         "LINE-BLOCK",
				Handle:       "B1",
				LayerName:    "0",
				SecondName:   "LINE-BLOCK-2",
				BasePoint:    core.Point{X: 0.5, Y: 0.1, Z: 1.1},
				XrefPathName: "",
				Description:  "This is a test block.",
				Entities: entities.EntitySlice{
					&entities.Vertex{Location: core.Point{X: 1.5, Y: 2.9}},
					&entities.Vertex{Location: core.Point{X: 2.5, Y: 1.9}},
				},
			},
			&Block{
				Name:         "LINE-BLOCK",
				Handle:       "B1",
				LayerName:    "0",
				SecondName:   "LINE-BLOCK-2",
				BasePoint:    core.Point{X: 0.5, Y: 0.1, Z: 1.1},
				XrefPathName: "",
				Description:  "This is a test block.",
				Entities: entities.EntitySlice{
					&entities.Vertex{Location: core.Point{X: 1.5, Y: 2.9}},
					&entities.Point{},
				},
			},
			false,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.b1.Equals(test.b2), test.equals)
	}
}

func TestBlockNotEqualToDifferentType(t *testing.T) {
	assert.False(t, Block{}.Equals(core.NewIntegerValue(0)))
}

func TestNewBlocksSection(t *testing.T) {
	expected := BlocksSection{
		"1": &Block{
			Name:      "1",
			LayerName: "1",
			BasePoint: core.Point{},
			Entities: entities.EntitySlice{
				&entities.Line{
					BaseEntity: entities.BaseEntity{
						On:        true,
						Visible:   true,
						LayerName: "7",
					},
					Start: core.Point{
						X: 12.29123377799988,
						Y: 54.33330607414246},
					End: core.Point{
						X: 12.29123377799988,
						Y: 13.66311264038086},
					ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
				},
				&entities.Text{
					BaseEntity: entities.BaseEntity{
						On:        true,
						Visible:   true,
						LayerName: "1",
					},
					FirstAlignmentPoint: core.Point{
						X: 12.40123374043651,
						Y: 37.14821090829416,
					},
					Height:             0.4999999999999999,
					Rotation:           270.0000006832453,
					Value:              "S38",
					RelativeXScale:     1.0,
					StyleName:          "STANDARD",
					ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
				},
			},
		},
		"2": &Block{
			Name:      "2",
			LayerName: "1",
			BasePoint: core.Point{X: 1.0, Y: 3.0},
			Entities: entities.EntitySlice{
				&entities.Line{
					BaseEntity: entities.BaseEntity{
						On:        true,
						Visible:   true,
						LayerName: "7",
					},
					Start: core.Point{
						X: 40.87370872497559,
						Y: 54.33330440521240},
					End: core.Point{
						X: 40.87370872497559,
						Y: 13.66311091184616},
					ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
				},
			},
		},
	}

	next := core.Tagger(strings.NewReader(dxfBlocksSection))
	tags := core.TagSlice(core.AllTags(next))

	section, err := NewBlocksSection(tags)

	assert.Equal(t, nil, err)
	assert.True(t, expected.Equals(section),
		"Expected %+v and %+v to be equals",
		spew.Sdump(expected), spew.Sdump(section))
}

func TestNewBlocksErrorParsingBlock(t *testing.T) {
	next := core.Tagger(strings.NewReader(dxfBlocksSection))
	tags := core.TagSlice(core.AllTags(next))

	tags[6].Value = core.NewStringValue("ERROR")

	section, err := NewBlocksSection(tags)

	assert.Nil(t, section)
	assert.NotNil(t, err)
}

func TestNewBlocksErrorParsingBlockEntities(t *testing.T) {
	next := core.Tagger(strings.NewReader(dxfBlocksSection))
	tags := core.TagSlice(core.AllTags(next))

	tags[10].Value = core.NewStringValue("ERROR ON ENTITIES")

	section, err := NewBlocksSection(tags)

	assert.Nil(t, section)
	assert.NotNil(t, err)
}

func TestDifferentLengthBlocksSectionAreNotEquals(t *testing.T) {
	s1 := BlocksSection{
		"B": &Block{Name: "B"},
	}

	assert.False(t, s1.Equals(BlocksSection{}))
}

func TestDifferentBlocksinBlocksSectionAreNotEquals(t *testing.T) {
	s1 := BlocksSection{
		"B": &Block{Name: "B"},
	}
	s2 := BlocksSection{
		"B": &Block{Name: "A"},
	}

	assert.False(t, s1.Equals(s2))
}

const dxfBlock = `  0
BLOCK
  5
B1
  8
0
  2
LINE-BLOCK
 10
0.5
 20
0.1
 30
1.1
  3
LINE-BLOCK-2
  1

  4
This is a test block.
`

const dxfBlocksSection = `  0
SECTION
  2
BLOCKS
  0
BLOCK
  8
1
  2
1
 70
64
 10
0.0
 20
0.0
  0
LINE
  8
7
 10
12.29123377799988
 20
54.33330607414246
 11
12.29123377799988
 21
13.66311264038086
  0
TEXT
  8
1
 10
12.40123374043651
 20
37.14821090829416
 40
0.4999999999999999
 50
270.0000006832453
 51
0.0000000000000000
 71
0
  1
S38
  0
ENDBLK
  0
BLOCK
  8
1
  2
2
 70
64
 10
1.0
 20
3.0
  0
LINE
  8
7
 10
40.87370872497559
 20
54.33330440521240
 11
40.87370872497559
 21
13.66311091184616
  0
ENDBLK
  0
ENDSEC
`
