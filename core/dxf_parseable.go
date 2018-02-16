package core

import (
	"fmt"
)

// TypeParser parses a specific DataType, returning an error or nil if the parsing was successful.
type TypeParser interface {
	Parse(d DataType) error
}

// SetStringFunc is a function that sets a string value.
type SetStringFunc func(string)

// SetIntFunc is a function that sets an integer value.
type SetIntFunc func(int64)

// SetFloatFunc is a function that sets a floating point value.
type SetFloatFunc func(float64)

// StringTypeParser is a TypeParser implementation that Parses String types and sets the value
// using the setter function.
type StringTypeParser struct {
	setter SetStringFunc
}

// Parse parses the DataType expecting it to be a String type.
func (parser StringTypeParser) Parse(d DataType) error {
	if value, ok := AsString(d); ok {
		parser.setter(value)
	} else {
		return fmt.Errorf("Error parsing type of %#v as a String", d)
	}
	return nil
}

// NewStringTypeParser creates a new StringTypeParser with the setter as passed.
func NewStringTypeParser(setter SetStringFunc) *StringTypeParser {
	parser := new(StringTypeParser)
	parser.setter = setter
	return parser
}

// NewStringTypeParserToVar creates a new StringTypeParser that sets the parsed
// value to the value of the passed string pointer.
func NewStringTypeParserToVar(variable *string) *StringTypeParser {
	parser := new(StringTypeParser)
	parser.setter = func(value string) {
		*variable = value
	}
	return parser
}

// IntTypeParser is a TypeParser implementation that Parses Integer types and sets the value
// using the setter function.
type IntTypeParser struct {
	setter SetIntFunc
}

// Parse parses the DataType expecting it to be an Integer type.
func (parser IntTypeParser) Parse(d DataType) error {
	if value, ok := AsInt(d); ok {
		parser.setter(value)
	} else {
		return fmt.Errorf("Error parsing type of %#v as an Integer", d)
	}
	return nil
}

// NewIntTypeParser creates a new IntTypeParser with the setter as passed.
func NewIntTypeParser(setter SetIntFunc) *IntTypeParser {
	parser := new(IntTypeParser)
	parser.setter = setter
	return parser
}

// NewIntTypeParserToVar creates a new IntTypeParser that sets the parsed
// value to the value of the passed int pointer.
func NewIntTypeParserToVar(variable *int64) *IntTypeParser {
	parser := new(IntTypeParser)
	parser.setter = func(value int64) {
		*variable = value
	}
	return parser
}

// FloatTypeParser is a TypeParser implementation that Parses Float types and sets the value
// using the setter function.
type FloatTypeParser struct {
	setter SetFloatFunc
}

// Parse parses the DataType expecting it to be a Float type.
func (parser FloatTypeParser) Parse(d DataType) error {
	if value, ok := AsFloat(d); ok {
		parser.setter(value)
	} else {
		return fmt.Errorf("Error parsing type of %#v as a Float", d)
	}
	return nil
}

// NewFloatTypeParser creates a new FloatTypeParser with the setter as passed.
func NewFloatTypeParser(setter SetFloatFunc) *FloatTypeParser {
	parser := new(FloatTypeParser)
	parser.setter = setter
	return parser
}

// NewFloatTypeParserToVar creates a new FloatTypeParser that sets the parsed
// value to the value of the passed float64 pointer.
func NewFloatTypeParserToVar(variable *float64) *FloatTypeParser {
	parser := new(FloatTypeParser)
	parser.setter = func(value float64) {
		*variable = value
	}
	return parser
}

// DxfParseable is the base abstraction for any element in a DXF file that is composed by tags.
// It defines the basic boilerplate to support parsing and error handling of a slice of tags that
// composes the element.
type DxfParseable struct {
	tagParsers map[int]TypeParser
}

// Init initializes the DxfParseable's parser map so that it can be used by the Parse method.
func (element *DxfParseable) Init(parsers map[int]TypeParser) {
	element.tagParsers = parsers
}

// Update the tagParsers with the content on parsers.
func (element *DxfParseable) Update(parsers map[int]TypeParser) {
	if len(element.tagParsers) == 0 {
		element.tagParsers = parsers
	} else {
		for key, value := range parsers {
			element.tagParsers[key] = value
		}
	}
}

// Parse parses the slice of tags using the configured parser map.
// Returns an error if any error happens during the process, otherwise it returns nil.
func (element *DxfParseable) Parse(tags TagSlice) error {
	for _, tag := range tags.RegularTags() {
		if parser, ok := element.tagParsers[tag.Code]; ok {
			err := parser.Parse(tag.Value)
			if err != nil {
				return err
			}
		} else {
			Log.Printf("Discarding tag: %+v\n", tag.ToString())
		}
	}
	return nil
}
