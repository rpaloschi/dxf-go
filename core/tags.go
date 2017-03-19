package core

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// Tag from a DXF file.
type Tag struct {
	Code  int
	Value DataType
}

// NoneTag a constant that represents a nul tag.
var NoneTag = Tag{999999, &String{"NONE"}}

// dataTypeFactory a factory function that receives a string and return an instance
// of a DataType. The string should contain the DataType value.
type dataTypeFactory func(string) (DataType, error)

// groupCodeTypes maps DXF group codes to DataTypeFactory functions. See the init
// functions for the known group codes.
var groupCodeTypes map[int]dataTypeFactory

// NextTagFunction is the prototype of a function that returns the next Tag in a stream.
type NextTagFunction func() (*Tag, error)

// Tagger function. Returns a NextTagFunction that, in turn, returns the tags
// from the stream sequentially each time it is called. It finishes when it returns
// an error or a NoneTag.
func Tagger(stream io.Reader) NextTagFunction {
	counter := 0
	scanner := bufio.NewScanner(stream)

	readLine := func() (string, error) {
		if scanner.Scan() {
			return scanner.Text(), nil
		} else if err := scanner.Err(); err != nil {
			return "", err
		}

		return "", nil
	}

	return func() (*Tag, error) {
		code, err := readLine()
		if err != nil {
			return &NoneTag, err
		}
		value, err := readLine()
		if err != nil {
			return &NoneTag, err
		}

		charsToTrim := " \r\n"
		counter += 2
		if len(code) > 0 && len(value) > 0 {
			intCode, err := strconv.Atoi(strings.Trim(code, charsToTrim))
			if err != nil {
				return &NoneTag, err
			}
			valueType, _ := groupCodeTypes[intCode](strings.Trim(value, charsToTrim))
			tag := new(Tag)
			tag.Code = intCode
			tag.Value = valueType
			return tag, nil
		}

		// EOF
		return &NoneTag, nil
	}
}

func init() {
	groupCodeTypes = make(map[int]dataTypeFactory)

	for code := 0; code < 10; code++ {
		groupCodeTypes[code] = NewString
	}

	for code := 10; code < 20; code++ {
		groupCodeTypes[code] = NewFloat
	}

	for code := 20; code < 60; code++ {
		groupCodeTypes[code] = NewFloat
	}

	for code := 60; code < 100; code++ {
		groupCodeTypes[code] = NewInteger
	}

	for code := 100; code < 106; code++ {
		groupCodeTypes[code] = NewString
	}

	for code := 110; code < 113; code++ {
		groupCodeTypes[code] = NewFloat
	}

	for code := 113; code < 150; code++ {
		groupCodeTypes[code] = NewFloat
	}

	for code := 170; code < 180; code++ {
		groupCodeTypes[code] = NewInteger
	}

	groupCodeTypes[210] = NewFloat

	for code := 211; code < 240; code++ {
		groupCodeTypes[code] = NewFloat
	}

	for code := 270; code < 290; code++ {
		groupCodeTypes[code] = NewInteger
	}

	for code := 290; code < 300; code++ {
		groupCodeTypes[code] = NewInteger
	}

	for code := 300; code < 370; code++ {
		groupCodeTypes[code] = NewString
	}

	for code := 370; code < 390; code++ {
		groupCodeTypes[code] = NewInteger
	}

	for code := 390; code < 400; code++ {
		groupCodeTypes[code] = NewString
	}

	for code := 400; code < 410; code++ {
		groupCodeTypes[code] = NewInteger
	}

	for code := 410; code < 420; code++ {
		groupCodeTypes[code] = NewString
	}

	for code := 420; code < 430; code++ {
		groupCodeTypes[code] = NewInteger
	}

	for code := 430; code < 440; code++ {
		groupCodeTypes[code] = NewString
	}

	for code := 440; code < 460; code++ {
		groupCodeTypes[code] = NewInteger
	}

	for code := 460; code < 470; code++ {
		groupCodeTypes[code] = NewFloat
	}

	for code := 470; code < 480; code++ {
		groupCodeTypes[code] = NewString
	}

	for code := 480; code < 482; code++ {
		groupCodeTypes[code] = NewString
	}

	for code := 999; code < 1010; code++ {
		groupCodeTypes[code] = NewString
	}

	for code := 1010; code < 1020; code++ {
		groupCodeTypes[code] = NewFloat
	}

	for code := 1020; code < 1060; code++ {
		groupCodeTypes[code] = NewFloat
	}

	for code := 1060; code < 1072; code++ {
		groupCodeTypes[code] = NewInteger
	}
}
