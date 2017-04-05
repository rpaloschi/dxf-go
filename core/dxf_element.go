package core

import "fmt"

type TypeParser interface {
	Parse(d DataType) error
}

type SetStringFunc func(string)

type SetIntFunc func(int)

type SetFloatFunc func(float64)

type StringTypeParser struct {
	setter SetStringFunc
}

func (parser StringTypeParser) Parse(d DataType) error {
	if value, ok := AsString(d); ok {
		parser.setter(value)
	} else {
		return fmt.Errorf("Error parsing type of %v as a String", d)
	}
	return nil
}

func NewStringTypeParser(setter SetStringFunc) *StringTypeParser {
	parser := new(StringTypeParser)
	parser.setter = setter
	return parser
}

func NewStringTypeParserToVar(variable *string) *StringTypeParser {
	parser := new(StringTypeParser)
	parser.setter = func(value string) {
		*variable = value
	}
	return parser
}

type IntTypeParser struct {
	setter SetIntFunc
}

func (parser IntTypeParser) Parse(d DataType) error {
	if value, ok := AsInt(d); ok {
		parser.setter(value)
	} else {
		return fmt.Errorf("Error parsing type of %v as a Int", d)
	}
	return nil
}

func NewIntTypeParser(setter SetIntFunc) *IntTypeParser {
	parser := new(IntTypeParser)
	parser.setter = setter
	return parser
}

func NewIntTypeParserToVar(variable *int) *IntTypeParser {
	parser := new(IntTypeParser)
	parser.setter = func(value int) {
		*variable = value
	}
	return parser
}

type FloatTypeParser struct {
	setter SetFloatFunc
}

func (parser FloatTypeParser) Parse(d DataType) error {
	if value, ok := AsFloat(d); ok {
		parser.setter(value)
	} else {
		return fmt.Errorf("Error parsing type of %v as a Float", d)
	}
	return nil
}

func NewFloatTypeParser(setter SetFloatFunc) *FloatTypeParser {
	parser := new(FloatTypeParser)
	parser.setter = setter
	return parser
}

func NewFloatTypeParserToVar(variable *float64) *FloatTypeParser {
	parser := new(FloatTypeParser)
	parser.setter = func(value float64) {
		*variable = value
	}
	return parser
}

type DxfElement struct {
	tagParsers map[int]TypeParser
}

func (element *DxfElement) Init(parsers map[int]TypeParser) {
	element.tagParsers = parsers
}

func (element *DxfElement) Parse(tags TagSlice) error {
	for _, tag := range tags.RegularTags() {
		if parser, ok := element.tagParsers[tag.Code]; ok {
			err := parser.Parse(tag.Value)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("Discarding tag for Layer: %+v\n", tag.ToString())
		}
	}
	return nil
}
