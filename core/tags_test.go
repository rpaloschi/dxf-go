package core

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const regularDXF = `  0
SECTION
  2
HEADER
  9
$ACADVER
  1
AC1018
  9
$DWGCODEPAGE
  3
ANSI_1252
  0
ENDSEC
  0
EOF
`

const noEOFDXF = `  0
SECTION
  2
HEADER
  9
$ACADVER
  1
AC1018
  9
$DWGCODEPAGE
  3
ANSI_1252
  0
ENDSEC
`

const regularDXFComments = `999
Comment0
  0
SECTION
  2
HEADER
  9
$ACADVER
999
Comment1
  1
AC1018
  9
$DWGCODEPAGE
  3
ANSI_1252
  0
ENDSEC
  0
EOF
`

const regularDXFAppDataAndXData = `  0
SECTION
  2
HEADER
  9
$ACADVER
  1
AC1018
  9
$DWGCODEPAGE
  3
ANSI_1252
  102
{DXFGrabber
330
999
  102
}
  0
ENDSEC
  1001
DXFGRABBER
  1000
XDATA_STRING
  0
EOF
`

const dxfEllipse = ` 0
ELLIPSE
5
3D2
330
1F
100
AcDbEntity
8
0
100
AcDbEllipse
10
0.0
20
0.0
30
0.0
11
2.60
21
1.50
31
0.0
210
0.0
220
0.0
230
1.0
40
0.33
41
0.0
42
6.28
`

var expectedRegularDxfTags = []*Tag{
	NewTag(0, NewStringValue("SECTION")),
	NewTag(2, NewStringValue("HEADER")),
	NewTag(9, NewStringValue("$ACADVER")),
	NewTag(1, NewStringValue("AC1018")),
	NewTag(9, NewStringValue("$DWGCODEPAGE")),
	NewTag(3, NewStringValue("ANSI_1252")),
	NewTag(0, NewStringValue("ENDSEC")),
	NewTag(0, NewStringValue("EOF")),
}

func TestTagEquality(t *testing.T) {
	str1 := NewStringValue("TEST")
	str2 := NewStringValue("STRING")
	int1 := NewIntegerValue(1001)
	int2 := NewIntegerValue(9)
	float1 := NewFloatValue(10.01)
	float2 := NewFloatValue(0.33)

	assert.Equal(t, NewTag(1, str1), NewTag(1, str1))
	assert.Equal(t, NewTag(2, int1), NewTag(2, int1))
	assert.Equal(t, NewTag(3, float1), NewTag(3, float1))

	assert.NotEqual(t, NewTag(1, str1), NewTag(1, str2))
	assert.NotEqual(t, NewTag(0, str1), NewTag(1, str1))
	assert.NotEqual(t, NewTag(0, str1), NewTag(0, int1))
	assert.NotEqual(t, NewTag(0, str1), NewTag(0, float1))

	assert.NotEqual(t, NewTag(1, int1), NewTag(1, int2))
	assert.NotEqual(t, NewTag(0, int1), NewTag(1, int1))
	assert.NotEqual(t, NewTag(0, int1), NewTag(0, str1))
	assert.NotEqual(t, NewTag(0, int1), NewTag(0, float1))

	assert.NotEqual(t, NewTag(1, float1), NewTag(1, float2))
	assert.NotEqual(t, NewTag(0, float1), NewTag(1, float1))
	assert.NotEqual(t, NewTag(0, float1), NewTag(0, str1))
	assert.NotEqual(t, NewTag(0, float1), NewTag(0, int1))
}

type TaggerTestSuite struct {
	suite.Suite
}

func (suite *TaggerTestSuite) SetupTest() {
}

func (suite *TaggerTestSuite) TestEmptyStream() {
	next := Tagger(strings.NewReader(""))
	tag, err := next()

	suite.Equal(&NoneTag, tag)
	suite.Equal(nil, err)
}

func (suite *TaggerTestSuite) TestNextTag() {
	next := Tagger(strings.NewReader(regularDXF))
	tag, err := next()

	suite.Equal(nil, err)
	suite.Equal(0, tag.Code)
	suite.Equal("SECTION", tag.Value.ToString())
}

func (suite *TaggerTestSuite) TestAllTags() {
	next := Tagger(strings.NewReader(regularDXF))
	suite.Equal(expectedRegularDxfTags, AllTags(next))
}

func (suite *TaggerTestSuite) TestAllTagsWithoutEof() {
	next := Tagger(strings.NewReader(noEOFDXF))

	expected := []*Tag{
		NewTag(0, NewStringValue("SECTION")),
		NewTag(2, NewStringValue("HEADER")),
		NewTag(9, NewStringValue("$ACADVER")),
		NewTag(1, NewStringValue("AC1018")),
		NewTag(9, NewStringValue("$DWGCODEPAGE")),
		NewTag(3, NewStringValue("ANSI_1252")),
		NewTag(0, NewStringValue("ENDSEC")),
	}

	suite.Equal(expected, AllTags(next))
}

func (suite *TaggerTestSuite) TestDXFWithComments() {
	next := Tagger(strings.NewReader(regularDXFComments))

	expected := []*Tag{
		NewTag(999, NewStringValue("Comment0")),
		NewTag(0, NewStringValue("SECTION")),
		NewTag(2, NewStringValue("HEADER")),
		NewTag(9, NewStringValue("$ACADVER")),
		NewTag(999, NewStringValue("Comment1")),
		NewTag(1, NewStringValue("AC1018")),
		NewTag(9, NewStringValue("$DWGCODEPAGE")),
		NewTag(3, NewStringValue("ANSI_1252")),
		NewTag(0, NewStringValue("ENDSEC")),
		NewTag(0, NewStringValue("EOF")),
	}

	suite.Equal(expected, AllTags(next))
}

func (suite *TaggerTestSuite) TestIntAndFloatTags() {
	intAndFloatTags := "  60\n1001\n  20\n15.54"
	next := Tagger(strings.NewReader(intAndFloatTags))

	expected := []*Tag{
		NewTag(60, NewIntegerValue(1001)),
		NewTag(20, NewFloatValue(15.54)),
	}

	suite.Equal(expected, AllTags(next))
}

func TestTaggerTestSuite(t *testing.T) {
	suite.Run(t, new(TaggerTestSuite))
}

type TagSliceTestSuite struct {
	suite.Suite
	tags TagSlice
}

func (suite *TagSliceTestSuite) SetupTest() {
	next := Tagger(strings.NewReader(regularDXF))
	suite.tags = TagSlice(AllTags(next))
}

func (suite *TagSliceTestSuite) TestTagIndex() {
	index := suite.tags.TagIndex(0, 0, len(suite.tags))
	suite.Equal(0, index)

	index = suite.tags.TagIndex(0, index+1, len(suite.tags))
	suite.Equal(6, index)
}

func (suite *TagSliceTestSuite) TestInexistentTagIndex() {
	suite.Equal(-1, suite.tags.TagIndex(50, 0, len(suite.tags)))
}

func (suite *TagSliceTestSuite) TestAllWithCode() {
	expected := []*Tag{
		NewTag(0, NewStringValue("SECTION")),
		NewTag(0, NewStringValue("ENDSEC")),
		NewTag(0, NewStringValue("EOF")),
	}
	suite.Equal(expected, suite.tags.AllWithCode(0))
	suite.Equal([]*Tag{}, suite.tags.AllWithCode(50))
}

func (suite *TagSliceTestSuite) TestRegularTags() {
	next := Tagger(strings.NewReader(regularDXFAppDataAndXData))
	tags := TagSlice(AllTags(next))

	suite.Equal(expectedRegularDxfTags, tags.RegularTags())
}

func (suite *TagSliceTestSuite) TestXDataTags() {
	next := Tagger(strings.NewReader(regularDXFAppDataAndXData))
	tags := TagSlice(AllTags(next))

	expected := []*Tag{
		NewTag(1001, NewStringValue("DXFGRABBER")),
		NewTag(1000, NewStringValue("XDATA_STRING")),
	}

	suite.Equal(expected, tags.XDataTags())
}

func (suite *TagSliceTestSuite) TestAppDataTags() {
	next := Tagger(strings.NewReader(regularDXFAppDataAndXData))
	tags := TagSlice(AllTags(next))

	expected := []*Tag{
		NewTag(102, NewStringValue("{DXFGrabber")),
		NewTag(330, NewStringValue("999")),
		NewTag(102, NewStringValue("}")),
	}
	appData := tags.AppDataTags()

	suite.Equal(expected, appData["{DXFGrabber"])
}

func (suite *TagSliceTestSuite) TestSubclassesTags() {
	next := Tagger(strings.NewReader(dxfEllipse))
	tags := TagSlice(AllTags(next))

	subclasses := tags.SubclassesTags()
	suite.Equal(3, len(subclasses))
	suite.Equal(3, len(subclasses["noname"]))
	suite.Equal(1, len(subclasses["AcDbEntity"]))
	suite.Equal(12, len(subclasses["AcDbEllipse"]))
}

func TestTagSliceTestSuite(t *testing.T) {
	suite.Run(t, new(TagSliceTestSuite))
}

func TestTagGroups(t *testing.T) {
	next := Tagger(strings.NewReader(regularDXF))
	tags := TagSlice(AllTags(next))

	groups := TagGroups(tags[2:len(tags)-2], 9)

	expected := []TagSlice{
		{
			NewTag(9, NewStringValue("$ACADVER")),
			NewTag(1, NewStringValue("AC1018")),
		},
		{
			NewTag(9, NewStringValue("$DWGCODEPAGE")),
			NewTag(3, NewStringValue("ANSI_1252")),
		},
	}

	assert.EqualValues(t, expected, groups)
}
