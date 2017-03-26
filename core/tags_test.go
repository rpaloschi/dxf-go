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
	&Tag{Code: 0, Value: &String{"SECTION"}},
	&Tag{Code: 2, Value: &String{"HEADER"}},
	&Tag{Code: 9, Value: &String{"$ACADVER"}},
	&Tag{Code: 1, Value: &String{"AC1018"}},
	&Tag{Code: 9, Value: &String{"$DWGCODEPAGE"}},
	&Tag{Code: 3, Value: &String{"ANSI_1252"}},
	&Tag{Code: 0, Value: &String{"ENDSEC"}},
	&Tag{Code: 0, Value: &String{"EOF"}},
}

func TestTagEquality(t *testing.T) {
	str1, _ := NewString("TEST")
	str2, _ := NewString("STRING")
	int1, _ := NewInteger("1001")
	int2, _ := NewInteger("9")
	float1, _ := NewFloat("10.01")
	float2, _ := NewFloat("0.33")

	assert.Equal(t, &Tag{Code: 1, Value: str1}, &Tag{Code: 1, Value: str1})
	assert.Equal(t, &Tag{Code: 2, Value: int1}, &Tag{Code: 2, Value: int1})
	assert.Equal(t, &Tag{Code: 3, Value: float1}, &Tag{Code: 3, Value: float1})

	assert.NotEqual(t, &Tag{Code: 1, Value: str1}, &Tag{Code: 1, Value: str2})
	assert.NotEqual(t, &Tag{Code: 0, Value: str1}, &Tag{Code: 1, Value: str1})
	assert.NotEqual(t, &Tag{Code: 0, Value: str1}, &Tag{Code: 0, Value: int1})
	assert.NotEqual(t, &Tag{Code: 0, Value: str1}, &Tag{Code: 0, Value: float1})

	assert.NotEqual(t, &Tag{Code: 1, Value: int1}, &Tag{Code: 1, Value: int2})
	assert.NotEqual(t, &Tag{Code: 0, Value: int1}, &Tag{Code: 1, Value: int1})
	assert.NotEqual(t, &Tag{Code: 0, Value: int1}, &Tag{Code: 0, Value: str1})
	assert.NotEqual(t, &Tag{Code: 0, Value: int1}, &Tag{Code: 0, Value: float1})

	assert.NotEqual(t, &Tag{Code: 1, Value: float1}, &Tag{Code: 1, Value: float2})
	assert.NotEqual(t, &Tag{Code: 0, Value: float1}, &Tag{Code: 1, Value: float1})
	assert.NotEqual(t, &Tag{Code: 0, Value: float1}, &Tag{Code: 0, Value: str1})
	assert.NotEqual(t, &Tag{Code: 0, Value: float1}, &Tag{Code: 0, Value: int1})
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
		&Tag{Code: 0, Value: &String{"SECTION"}},
		&Tag{Code: 2, Value: &String{"HEADER"}},
		&Tag{Code: 9, Value: &String{"$ACADVER"}},
		&Tag{Code: 1, Value: &String{"AC1018"}},
		&Tag{Code: 9, Value: &String{"$DWGCODEPAGE"}},
		&Tag{Code: 3, Value: &String{"ANSI_1252"}},
		&Tag{Code: 0, Value: &String{"ENDSEC"}},
	}

	suite.Equal(expected, AllTags(next))
}

func (suite *TaggerTestSuite) TestDXFWithComments() {
	next := Tagger(strings.NewReader(regularDXFComments))

	expected := []*Tag{
		&Tag{Code: 999, Value: &String{"Comment0"}},
		&Tag{Code: 0, Value: &String{"SECTION"}},
		&Tag{Code: 2, Value: &String{"HEADER"}},
		&Tag{Code: 9, Value: &String{"$ACADVER"}},
		&Tag{Code: 999, Value: &String{"Comment1"}},
		&Tag{Code: 1, Value: &String{"AC1018"}},
		&Tag{Code: 9, Value: &String{"$DWGCODEPAGE"}},
		&Tag{Code: 3, Value: &String{"ANSI_1252"}},
		&Tag{Code: 0, Value: &String{"ENDSEC"}},
		&Tag{Code: 0, Value: &String{"EOF"}},
	}

	suite.Equal(expected, AllTags(next))
}

func (suite *TaggerTestSuite) TestIntAndFloatTags() {
	intAndFloatTags := "  60\n1001\n  20\n15.54"
	next := Tagger(strings.NewReader(intAndFloatTags))

	expected := []*Tag{
		&Tag{Code: 60, Value: &Integer{1001}},
		&Tag{Code: 20, Value: &Float{15.54}},
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
		&Tag{Code: 0, Value: &String{value: "SECTION"}},
		&Tag{Code: 0, Value: &String{value: "ENDSEC"}},
		&Tag{Code: 0, Value: &String{value: "EOF"}},
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
		&Tag{Code: 1001, Value: &String{value: "DXFGRABBER"}},
		&Tag{Code: 1000, Value: &String{value: "XDATA_STRING"}},
	}

	suite.Equal(expected, tags.XDataTags())
}

func (suite *TagSliceTestSuite) TestAppDataTags() {
	next := Tagger(strings.NewReader(regularDXFAppDataAndXData))
	tags := TagSlice(AllTags(next))

	expected := []*Tag{
		&Tag{Code: 102, Value: &String{value: "{DXFGrabber"}},
		&Tag{Code: 330, Value: &String{value: "999"}},
		&Tag{Code: 102, Value: &String{value: "}"}},
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
