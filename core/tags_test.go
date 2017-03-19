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

func allTags(next NextTagFunction) []*Tag {
	tags := make([]*Tag, 0)

	tag, _ := next()
	for *tag != NoneTag {
		tags = append(tags, tag)
		tag, _ = next()
	}

	return tags
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

	expected := []*Tag{
		&Tag{Code: 0, Value: &String{"SECTION"}},
		&Tag{Code: 2, Value: &String{"HEADER"}},
		&Tag{Code: 9, Value: &String{"$ACADVER"}},
		&Tag{Code: 1, Value: &String{"AC1018"}},
		&Tag{Code: 9, Value: &String{"$DWGCODEPAGE"}},
		&Tag{Code: 3, Value: &String{"ANSI_1252"}},
		&Tag{Code: 0, Value: &String{"ENDSEC"}},
		&Tag{Code: 0, Value: &String{"EOF"}},
	}

	suite.Equal(expected, allTags(next))
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

	suite.Equal(expected, allTags(next))
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

	suite.Equal(expected, allTags(next))
}

func (suite *TaggerTestSuite) TestIntAndFloatTags() {
	intAndFloatTags := "  60\n1001\n  20\n15.54"
	next := Tagger(strings.NewReader(intAndFloatTags))

	expected := []*Tag{
		&Tag{Code: 60, Value: &Integer{1001}},
		&Tag{Code: 20, Value: &Float{15.54}},
	}

	suite.Equal(expected, allTags(next))
}

func TestTaggerTestSuite(t *testing.T) {
	suite.Run(t, new(TaggerTestSuite))
}
