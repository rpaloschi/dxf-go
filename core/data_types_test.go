package core

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DataTypesTestSuite struct {
	suite.Suite
	strType   DataType
	intType   DataType
	floatType DataType
}

func (suite *DataTypesTestSuite) SetupTest() {
	suite.strType, _ = NewString("DXF")
	suite.intType, _ = NewInteger("2017")
	suite.floatType, _ = NewFloat("20.17")
}

func (suite *DataTypesTestSuite) TestToString() {
	suite.Equal("DXF", suite.strType.ToString())
	suite.Equal("2017", suite.intType.ToString())
	suite.Equal("20.17", suite.floatType.ToString())
}

func (suite *DataTypesTestSuite) TestAsString() {
	value, ok := AsString(suite.strType)
	suite.True(ok)
	suite.Equal("DXF", value)

	_, ok = AsString(suite.intType)
	suite.False(ok)

	_, ok = AsString(suite.floatType)
	suite.False(ok)
}

func (suite *DataTypesTestSuite) TestAsInteger() {
	value, ok := AsInt(suite.intType)
	suite.True(ok)
	suite.Equal(2017, value)

	_, ok = AsInt(suite.strType)
	suite.False(ok)

	_, ok = AsInt(suite.floatType)
	suite.False(ok)
}

func (suite *DataTypesTestSuite) TestAsFloat() {
	value, ok := AsFloat(suite.floatType)
	suite.True(ok)
	suite.Equal(20.17, value)

	_, ok = AsFloat(suite.strType)
	suite.False(ok)

	_, ok = AsFloat(suite.intType)
	suite.False(ok)
}

func TestDataTypesTestSuite(t *testing.T) {
	suite.Run(t, new(DataTypesTestSuite))
}
