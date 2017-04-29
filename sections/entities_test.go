package sections

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rpaloschi/dxf-go/core"
	"github.com/rpaloschi/dxf-go/entities"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewEntitiesSection(t *testing.T) {
	base := entities.BaseEntity{
		Handle:    "3E5",
		LayerName: "0",
		On:        true,
		Visible:   true,
	}

	expected := EntitiesSection{
		Entities: entities.EntitySlice{
			&entities.Arc{
				BaseEntity:         base,
				Center:             core.Point{X: 1.1, Y: 1.2, Z: 1.3},
				Radius:             5.0,
				StartAngle:         10.0,
				EndAngle:           90.0,
				ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
			},
			&entities.Insert{
				BaseEntity:       base,
				InsertionPoint:   core.Point{X: 1.1, Y: 1.2, Z: 1.3},
				ScaleFactorX:     1.0,
				ScaleFactorY:     1.0,
				ScaleFactorZ:     1.0,
				RotationAngle:    0.0,
				ColumnCount:      1,
				RowCount:         1,
				ColumnSpacing:    0.0,
				RowSpacing:       0.0,
				AttributesFollow: true,
				Entities: entities.EntitySlice{
					&entities.Vertex{
						BaseEntity: entities.BaseEntity{
							Handle:    "LH",
							LayerName: "0",
							On:        true,
							Visible:   true,
						},
						Location: core.Point{X: 1.1, Y: 1.2}},
					&entities.Vertex{
						BaseEntity: entities.BaseEntity{
							Handle:    "LH1",
							LayerName: "0",
							On:        true,
							Visible:   true,
						},
						Location: core.Point{X: 1.3, Y: 5.2}},
				},
				ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
			},
			&entities.Line{
				BaseEntity: entities.BaseEntity{
					Handle:    "LH",
					LayerName: "0",
					On:        true,
					Visible:   true,
				},
				Start:              core.Point{X: 1.1, Y: 1.2, Z: 1.3},
				End:                core.Point{X: 2.0, Y: 5.0, Z: 7.0},
				ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
			},
			&entities.Point{
				BaseEntity:         base,
				Location:           core.Point{X: 1.1, Y: 1.2, Z: 1.3},
				ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
			},
			&entities.Circle{
				BaseEntity:         base,
				Center:             core.Point{X: 1.1, Y: 1.2, Z: 1.3},
				Radius:             5.0,
				ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
			},
			&entities.Text{
				BaseEntity:         base,
				RelativeXScale:     1.0,
				StyleName:          "STANDARD",
				ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
			},
			&entities.Polyline{
				BaseEntity:         base,
				Vertices:           entities.VertexSlice{},
				ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
			},
			&entities.LWPolyline{
				BaseEntity:         base,
				ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
			},
			&entities.Ellipse{
				BaseEntity:            base,
				Center:                core.Point{X: 1.1, Y: 1.2, Z: 1.3},
				MajorAxisEnd:          core.Point{X: 2.0, Y: 5.0, Z: 7.0},
				ExtrusionDirection:    core.Point{X: 0.0, Y: 0.0, Z: 1.0},
				MinorToMajorAxisRatio: 1.0,
				StartParameter:        0.0,
				EndParameter:          360.0,
			},
			&entities.Spline{
				BaseEntity:            base,
				KnotValues:            []float64{},
				Weights:               []float64{},
				ControlPoints:         core.PointSlice{},
				FitPoints:             core.PointSlice{},
				KnotTolerance:         0.0000001,
				ControlPointTolerance: 0.0000001,
				FitTolerance:          0.0000000001,
			},
		},
	}

	next := core.Tagger(strings.NewReader(dxfEntitiesSection))
	tags := core.TagSlice(core.AllTags(next))

	section, err := NewEntitiesSection(tags)

	assert.Equal(t, nil, err)
	assert.True(t, expected.Equals(section),
		"Expected %+v and %+v to be equals",
		spew.Sdump(expected), spew.Sdump(section))
}

func TestNewEntitiesSectionInvalidEntityIsIgnored(t *testing.T) {
	expected := EntitiesSection{
		Entities: entities.EntitySlice{
			&entities.Line{
				BaseEntity: entities.BaseEntity{
					Handle:    "LH",
					LayerName: "0",
					On:        true,
					Visible:   true,
				},
				Start:              core.Point{X: 1.1, Y: 1.2, Z: 1.3},
				End:                core.Point{X: 2.0, Y: 5.0, Z: 7.0},
				ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
			},
		},
	}
	next := core.Tagger(strings.NewReader(dxfLineAndInvalidEntitiesSection))
	tags := core.TagSlice(core.AllTags(next))

	section, err := NewEntitiesSection(tags)

	assert.Equal(t, nil, err)
	assert.True(t, expected.Equals(section),
		"Expected %+v and %+v to be equals",
		spew.Sdump(expected), spew.Sdump(section))
}

//dxfLineAndInvalidEntitiesSection

func TestNewEntitiesSectionEmpty(t *testing.T) {
	next := core.Tagger(strings.NewReader(dxfEmptyEntitiesSection))
	tags := core.TagSlice(core.AllTags(next))

	section, err := NewEntitiesSection(tags)

	assert.Equal(t, nil, err)
	assert.True(t, EntitiesSection{}.Equals(section))
}

func TestNewEntitiesSectionParsingError(t *testing.T) {
	next := core.Tagger(strings.NewReader(dxfLineEntitiesSection))
	tags := core.TagSlice(core.AllTags(next))

	tags[5] = core.NewTag(10, core.NewStringValue("ERROR"))

	section, err := NewEntitiesSection(tags)

	assert.Nil(t, section)
	assert.NotNil(t, err)
}

func TestEntitiesSectionEquality(t *testing.T) {
	testCases := []struct {
		es1    EntitiesSection
		es2    EntitiesSection
		equals bool
	}{
		{
			es1:    EntitiesSection{},
			es2:    EntitiesSection{},
			equals: true,
		},
		{
			es1: EntitiesSection{
				Entities: entities.EntitySlice{
					entities.Arc{},
				},
			},
			es2: EntitiesSection{
				Entities: entities.EntitySlice{
					entities.SeqEnd{},
				},
			},
			equals: false,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.es1.Equals(&test.es2), test.equals)
	}
}

func TestEntitiesSectionNotEqualToDifferentType(t *testing.T) {
	assert.False(t, EntitiesSection{}.Equals(core.NewIntegerValue(1)))
}

const dxfEntitiesSection = `  0
SECTION
  2
ENTITIES
  0
ARC
  5
3E5
  8
0
 10
1.1
 20
1.2
 30
1.3
 40
5.0
 50
10.0
 51
90.0
  0
INSERT
  5
3E5
  8
0
 10
1.1
 20
1.2
 30
1.3
 66
1
  0
VERTEX
  5
LH
  8
0
 10
1.1
 20
1.2
  0
VERTEX
  5
LH1
  8
0
 10
1.3
 20
5.2
  0
SEQEND
  0
LINE
  5
LH
  8
0
 10
1.1
 20
1.2
 30
1.3
 11
2.0
 21
5.0
 31
7.0
  0
POINT
  5
3E5
  8
0
 10
1.1
 20
1.2
 30
1.3
  0
CIRCLE
  5
3E5
  8
0
 10
1.1
 20
1.2
 30
1.3
 40
5.0
  0
TEXT
  5
3E5
  8
0
  0
POLYLINE
  5
3E5
  8
0
  0
SEQEND
  0
LWPOLYLINE
  5
3E5
  8
0
  0
ELLIPSE
  5
3E5
  8
0
 10
1.1
 20
1.2
 30
1.3
 11
2.0
 21
5.0
 31
7.0
  0
SPLINE
  5
3E5
  8
0
  0
ENDSEC
`

const dxfEmptyEntitiesSection = `  0
SECTION
  2
ENTITIES
  0
ENDSEC
`

const dxfLineEntitiesSection = `  0
SECTION
  2
ENTITIES
  0
LINE
  5
LH
  8
0
 10
1.1
 20
1.2
 30
1.3
 11
2.0
 21
5.0
 31
7.0
  0
ENDSEC
`

const dxfLineAndInvalidEntitiesSection = `  0
SECTION
  2
ENTITIES
  0
LINE
  5
LH
  8
0
 10
1.1
 20
1.2
 30
1.3
 11
2.0
 21
5.0
 31
7.0
  0
BLABLA
  0
ENDSEC
`
