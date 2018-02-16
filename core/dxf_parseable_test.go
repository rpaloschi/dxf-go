package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestStringTypeParser(t *testing.T) {
	var returned string

	parser := NewStringTypeParser(func(value string) {
		returned = value
	})

	expected := "STRING"
	err := parser.Parse(NewStringValue(expected))

	assert.Nil(t, err)
	assert.Equal(t, expected, returned)
}

func TestStringTypeParserInvalidType(t *testing.T) {
	parser := NewStringTypeParser(func(value string) {})

	expected := int64(1000)
	err := parser.Parse(NewIntegerValue(expected))

	assert.Equal(t,
		"Error parsing type of &core.Integer{value:1000} as a String",
		err.Error())
}

func TestStringTypeParserToVar(t *testing.T) {
	var returned string
	expected := "STRING"
	err := NewStringTypeParserToVar(&returned).Parse(NewStringValue(expected))

	assert.Nil(t, err)
	assert.Equal(t, expected, returned)
}

func TestIntTypeParser(t *testing.T) {
	var returned int64

	parser := NewIntTypeParser(func(value int64) {
		returned = value
	})

	expected := int64(1000)
	err := parser.Parse(NewIntegerValue(expected))

	assert.Nil(t, err)
	assert.Equal(t, expected, returned)
}

func TestIntTypeParserInvalidType(t *testing.T) {
	parser := NewIntTypeParser(func(value int64) {})

	expected := 10.50
	err := parser.Parse(NewFloatValue(expected))

	assert.Equal(t,
		"Error parsing type of &core.Float{value:10.5} as an Integer",
		err.Error())
}

func TestIntTypeParserToVar(t *testing.T) {
	var returned int64
	expected := int64(1234)
	err := NewIntTypeParserToVar(&returned).Parse(NewIntegerValue(expected))

	assert.Nil(t, err)
	assert.Equal(t, expected, returned)
}

func TestFloatTypeParser(t *testing.T) {
	var returned float64

	parser := NewFloatTypeParser(func(value float64) {
		returned = value
	})

	expected := 1.1
	err := parser.Parse(NewFloatValue(expected))

	assert.Nil(t, err)
	assert.Equal(t, expected, returned)
}

func TestFloatTypeParserInvalidType(t *testing.T) {
	parser := NewFloatTypeParser(func(value float64) {})

	expected := "INVALID"
	err := parser.Parse(NewStringValue(expected))

	assert.Equal(t,
		"Error parsing type of &core.String{value:\"INVALID\"} as a Float",
		err.Error())
}

func TestFloatTypeParserToVar(t *testing.T) {
	var returned float64
	expected := 1.234
	err := NewFloatTypeParserToVar(&returned).Parse(NewFloatValue(expected))

	assert.Nil(t, err)
	assert.Equal(t, expected, returned)
}

type DxfElementTestSuite struct {
	suite.Suite
	element     *DxfParseable
	stringValue string
	intValue    int64
	floatValue  float64
}

func (suite *DxfElementTestSuite) SetupTest() {
	suite.element = new(DxfParseable)
	suite.element.Init(map[int]TypeParser{
		2:  NewStringTypeParserToVar(&suite.stringValue),
		60: NewIntTypeParserToVar(&suite.intValue),
		20: NewFloatTypeParserToVar(&suite.floatValue),
	})
}

func (suite *DxfElementTestSuite) TestValidTags() {
	tags := TagSlice{
		NewTag(20, NewFloatValue(1.5)),
		NewTag(60, NewIntegerValue(15)),
		NewTag(2, NewStringValue("Fifteen")),
	}

	err := suite.element.Parse(tags)
	suite.Nil(err)
	suite.Equal("Fifteen", suite.stringValue)
	suite.Equal(1.5, suite.floatValue)
	suite.Equal(int64(15), suite.intValue)
}

func (suite *DxfElementTestSuite) TestUpdate() {
	suite.Equal(3, len(suite.element.tagParsers))

	suite.element.Update(map[int]TypeParser{
		9: NewStringTypeParserToVar(&suite.stringValue),
	})

	suite.Equal(4, len(suite.element.tagParsers))

	empty := new(DxfParseable)
	empty.Update(map[int]TypeParser{
		9: NewStringTypeParserToVar(&suite.stringValue),
	})

	suite.Equal(1, len(empty.tagParsers))
}

func (suite *DxfElementTestSuite) TestInvalidValidTagType() {
	tags := TagSlice{
		NewTag(20, NewFloatValue(1.5)),
		NewTag(2, NewIntegerValue(15)),
	}

	err := suite.element.Parse(tags)
	suite.Equal(
		"Error parsing type of &core.Integer{value:15} as a String",
		err.Error())
}

func (suite *DxfElementTestSuite) TestUnregisteredTagIsIgnored() {
	tags := TagSlice{
		NewTag(60, NewIntegerValue(15)),
		NewTag(3, NewStringValue("Fifteen")),
	}

	err := suite.element.Parse(tags)
	suite.Nil(err)
	suite.Equal(int64(15), suite.intValue)
}

func TestDxfElementTestSuite(t *testing.T) {
	suite.Run(t, new(DxfElementTestSuite))
}
