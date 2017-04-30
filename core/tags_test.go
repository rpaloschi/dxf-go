package core

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
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

	assert.True(t, NewTag(1, str1).Equals(NewTag(1, str1)))
	assert.True(t, NewTag(2, int1).Equals(NewTag(2, int1)))
	assert.True(t, NewTag(3, float1).Equals(NewTag(3, float1)))

	assert.False(t, NewTag(1, str1).Equals(NewTag(1, str2)))
	assert.False(t, NewTag(0, str1).Equals(NewTag(1, str1)))
	assert.False(t, NewTag(0, str1).Equals(NewTag(0, int1)))
	assert.False(t, NewTag(0, str1).Equals(NewTag(0, float1)))

	assert.False(t, NewTag(1, int1).Equals(NewTag(1, int2)))
	assert.False(t, NewTag(0, int1).Equals(NewTag(1, int1)))
	assert.False(t, NewTag(0, int1).Equals(NewTag(0, str1)))
	assert.False(t, NewTag(0, int1).Equals(NewTag(0, float1)))

	assert.False(t, NewTag(1, float1).Equals(NewTag(1, float2)))
	assert.False(t, NewTag(0, float1).Equals(NewTag(1, float1)))
	assert.False(t, NewTag(0, float1).Equals(NewTag(0, str1)))
	assert.False(t, NewTag(0, float1).Equals(NewTag(0, int1)))

	assert.False(t, NewTag(0, float1).Equals(TagSlice{}))
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

func (suite *TaggerTestSuite) TestEmptyValueIsValid() {
	expected := Tag{Code: 1, Value: NewStringValue("")}

	next := Tagger(strings.NewReader("  1\n\n"))
	tag, err := next()

	suite.Equal(nil, err)
	suite.True(expected.Equals(tag))
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

type MockReader struct {
	Data []string
	Done []bool
	Err  []error
}

func (r *MockReader) Read(p []byte) (int, error) {
	if len(r.Data) == 0 {
		return 0, io.EOF
	}

	data := r.Data[0]
	r.Data = r.Data[1:]
	done := r.Done[0]
	r.Done = r.Done[1:]
	err := r.Err[0]
	r.Err = r.Err[1:]

	copy(p, []byte(data))
	if done {
		return 0, err
	}
	return len([]byte(data)), nil
}

func (suite *TaggerTestSuite) TestNextTagFunctionCodeError() {
	reader := &MockReader{
		Data: []string{""},
		Done: []bool{true},
		Err:  []error{io.ErrUnexpectedEOF}}
	next := Tagger(reader)

	tag, err := next()

	suite.Equal(NoneTag, *tag)
	suite.Equal(io.ErrUnexpectedEOF, err)
}

func (suite *TaggerTestSuite) TestNextTagFunctionValueError() {
	reader := &MockReader{
		Data: []string{"10\n", ""},
		Done: []bool{false, true},
		Err:  []error{nil, io.ErrUnexpectedEOF}}
	next := Tagger(reader)

	tag, err := next()

	suite.Equal(NoneTag, *tag)
	suite.Equal(io.ErrUnexpectedEOF, err)
}

func (suite *TaggerTestSuite) TestNextTagFunctionInvalidCodeError() {
	reader := &MockReader{
		Data: []string{"INVALID\nVALUE", ""},
		Done: []bool{false, true},
		Err:  []error{nil, io.EOF}}
	next := Tagger(reader)

	tag, err := next()

	suite.Equal(NoneTag, *tag)
	suite.Equal("strconv.Atoi: parsing \"INVALID\": invalid syntax", err.Error())
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

func (suite *TagSliceTestSuite) TestTagSliceEquality() {
	slice1 := TagSlice{NewTag(0, NewStringValue("SECTION"))}
	slice2 := TagSlice{NewTag(0, NewStringValue("SECTION"))}
	slice3 := TagSlice{
		NewTag(10, NewFloatValue(20.17)),
		NewTag(20, NewFloatValue(1.1)),
	}
	slice4 := TagSlice{
		NewTag(10, NewFloatValue(20.17)),
		NewTag(60, NewIntegerValue(2017)),
	}

	suite.True(slice1.Equals(slice2))
	suite.False(slice1.Equals(slice4))
	suite.False(slice3.Equals(slice4))
	suite.False(slice2.Equals(NewIntegerValue(1)))
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

	for index, slice := range expected {
		otherSlice := groups[index]

		assert.True(t, slice.Equals(otherSlice))
	}
}
