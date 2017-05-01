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

func (suite *HeaderTestSuite) TestGetMissingValue() {
	next := core.Tagger(strings.NewReader(testEmptyHeader))
	header := NewHeaderSection(core.TagSlice(core.AllTags(next)))

	result := header.Get("MISSING")
	suite.Equal(core.TagSlice{}, result)
}

func (suite *HeaderTestSuite) TestGetSimpleTagKey() {
	result := suite.header.Get("$ACADVER")
	expected := core.NewTag(1, core.NewStringValue("AC1021"))
	suite.Equal(expected, result[0])
}

func (suite *HeaderTestSuite) TestMultipleTagsKey() {
	result := suite.header.Get("$INSBASE")
	expected := core.TagSlice{
		core.NewTag(10, core.NewFloatValue(0.1)),
		core.NewTag(20, core.NewFloatValue(22.0)),
		core.NewTag(30, core.NewFloatValue(53.5)),
	}
	suite.Equal(expected, result)
}

const testHeaderDuplicateKey = `  0
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
  9
$INSBASE
 20
22.0
 30
53.5
  0
ENDSEC
`

func (suite *HeaderTestSuite) TestDuplicateTagsKey() {
	next := core.Tagger(strings.NewReader(testHeaderDuplicateKey))
	tags := core.TagSlice(core.AllTags(next))
	header := NewHeaderSection(tags)
	result := header.Get("$INSBASE")
	expected := core.TagSlice{
		core.NewTag(10, core.NewFloatValue(0.1)),
		core.NewTag(20, core.NewFloatValue(22.0)),
		core.NewTag(30, core.NewFloatValue(53.5)),
	}
	suite.Equal(expected, result)
}

func (suite *HeaderTestSuite) TestHeaderEquality() {
	header1 := &HeaderSection{Values: map[string]core.TagSlice{
		"KEY1": {core.NewTag(0, core.NewStringValue("SECTION"))}}}
	header2 := &HeaderSection{Values: map[string]core.TagSlice{
		"KEY1": {core.NewTag(0, core.NewStringValue("SECTION"))}}}
	header3 := &HeaderSection{Values: map[string]core.TagSlice{
		"KEY2": {core.NewTag(10, core.NewFloatValue(20.17))},
		"KEY3": {core.NewTag(20, core.NewFloatValue(1.1))}}}
	header4 := &HeaderSection{Values: map[string]core.TagSlice{
		"KEY2": {core.NewTag(10, core.NewFloatValue(20.17))},
		"KEY5": {core.NewTag(60, core.NewIntegerValue(2017))}}}
	header5 := &HeaderSection{Values: map[string]core.TagSlice{
		"KEY1": {core.NewTag(0, core.NewStringValue("OTHER"))}}}

	suite.True(header1.Equals(header2))
	suite.False(header1.Equals(header4))
	suite.False(header3.Equals(header4))
	suite.False(header1.Equals(header5))
	suite.False(header2.Equals(core.NewIntegerValue(1)))
}

func TestHeaderTestSuite(t *testing.T) {
	suite.Run(t, new(HeaderTestSuite))
}
