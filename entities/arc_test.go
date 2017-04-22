package entities

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type ArcTestSuite struct {
	suite.Suite
}

func (suite *ArcTestSuite) SetupTest() {
	//next := core.Tagger(strings.NewReader(testHeader))
	//suite.tags = core.TagSlice(core.AllTags(next))
	//suite.header = NewHeaderSection(suite.tags)
}

func (suite *ArcTestSuite) TestMinimalArc() {
	expected := Arc{
		BaseEntity: BaseEntity{
			Handle:    "3E5",
			LayerName: "0",
			On:        true,
			Visible:   true,
		},
		Center:             core.Point{X: 1.1, Y: 1.2, Z: 1.3},
		Radius:             5.0,
		StartAngle:         10.0,
		EndAngle:           90.0,
		ExtrusionDirection: core.Point{X: 0.0, Y: 0.0, Z: 1.0},
	}

	next := core.Tagger(strings.NewReader(testMinimalArc))
	arc, err := NewArc(core.TagSlice(core.AllTags(next)))

	suite.Nil(err)
	suite.True(expected.Equals(arc))
}

func TestArcTestSuite(t *testing.T) {
	suite.Run(t, new(ArcTestSuite))
}

const testMinimalArc = `  0
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
`

const testArcAllAttribs = `  0
ARC
  5
ALL_ARGS
  8
L1
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
`
