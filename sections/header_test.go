package sections

import (
	"testing"

	"github.com/rpaloschi/dxf-go/core"
	"github.com/stretchr/testify/suite"
	"strings"
)

const testHeader = `  0
SECTION
  2
HEADER
  9
$ACADVER
  1
AC1021
  9
$INSBASE
 10
0.1
 20
22.0
 30
53.5
  0
ENDSEC
`

const testEmptyHeader = `  0
SECTION
  2
HEADER
  0
ENDSEC
`

type HeaderTestSuite struct {
	suite.Suite
	tags   core.TagSlice
	header *HeaderSection
}

func (suite *HeaderTestSuite) SetupTest() {
	next := core.Tagger(strings.NewReader(testHeader))
	suite.tags = core.TagSlice(core.AllTags(next))
	suite.header = NewHeaderSection(suite.tags)
}

func (suite *HeaderTestSuite) TestDefaultValues() {
	next := core.Tagger(strings.NewReader(testEmptyHeader))
	header := NewHeaderSection(core.TagSlice(core.AllTags(next)))

	result := header.Get("$ACADVER")
	expected := core.NewTag(1, core.NewStringValue("AC1009"))
	suite.Equal(expected, result[0])

	result = header.Get("$DWGCODEPAGE")
	expected = core.NewTag(3, core.NewStringValue("ANSI_1252"))
	suite.Equal(expected, result[0])
}

func (suite *HeaderTestSuite) TestGetSimpleTagKey() {
	result := suite.header.Get("$ACADVER")
	expected := core.NewTag(1, core.NewStringValue("AC1021"))
	suite.Equal(expected, result[0])
}

func (suite *HeaderTestSuite) TestMultipleTagsKey() {
	result := suite.header.Get("$INSBASE")
	expected := []*core.Tag{
		core.NewTag(10, core.NewFloatValue(0.1)),
		core.NewTag(20, core.NewFloatValue(22.0)),
		core.NewTag(30, core.NewFloatValue(53.5)),
	}
	suite.Equal(expected, result)
}

func TestHeaderTestSuite(t *testing.T) {
	suite.Run(t, new(HeaderTestSuite))
}
