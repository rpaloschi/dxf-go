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

const APP_DATA_MARKER = 102
const SUBCLASS_MARKER = 100

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

// AllTags iterates until next finishes and returns all returned tags as a slice.
func AllTags(next NextTagFunction) []*Tag {
	tags := make([]*Tag, 0)

	tag, _ := next()
	for *tag != NoneTag {
		tags = append(tags, tag)
		tag, _ = next()
	}

	return tags
}

type TagSlice []*Tag

func (slice TagSlice) TagIndex(tagCode int, startingIndex int, endIndex int) int {
	for index := startingIndex; index < endIndex; index++ {
		if slice[index].Code == tagCode {
			return index
		}
	}
	return -1
}

func (slice TagSlice) AllWithCode(tagCode int) []*Tag {
	tags := make([]*Tag, 0)

	for _, tag := range slice {
		if tag.Code == tagCode {
			tags = append(tags, tag)
		}
	}

	return tags
}

func (slice TagSlice) RegularTags() []*Tag {
	tags := make([]*Tag, 0)

	inAppDataRange := false
	for _, tag := range slice {
		if tag.Code >= 1000 {
			continue
		}

		if tag.Code == APP_DATA_MARKER {
			if inAppDataRange {
				inAppDataRange = false
			} else {
				inAppDataRange = true
			}
			continue
		}

		if !inAppDataRange {
			tags = append(tags, tag)
		}
	}

	return tags
}

func (slice TagSlice) XDataTags() []*Tag {
	tags := make([]*Tag, 0)

	for _, tag := range slice {
		if tag.Code > 999 {
			tags = append(tags, tag)
		}
	}

	return tags
}

func (slice TagSlice) AppDataTags() map[string][]*Tag {
	appData := make(map[string][]*Tag)
	appTags := make([]*Tag, 0)

	for _, tag := range slice {
		if tag.Code == APP_DATA_MARKER {
			if tag.Value.ToString() == "}" {
				appTags = append(appTags, tag)
				appData[appTags[0].Value.ToString()] = appTags
				appTags = appTags[:0]
			} else {
				appTags = appTags[:0]
				appTags = append(appTags, tag)
			}
		} else {
			if len(appTags) > 0 {
				appTags = append(appTags, tag)
			}
		}
	}

	return appData
}

func (slice TagSlice) SubclassesTags() map[string][]*Tag {
	classes := make(map[string][]*Tag)
	tags := make([]*Tag, 0)
	name := "noname"

	for _, tag := range slice.RegularTags() {
		if tag.Code == SUBCLASS_MARKER {
			classes[name] = tags
			tags = tags[:0]
			name = tag.Value.ToString()
		} else {
			tags = append(tags, tag)
		}
	}
	classes[name] = tags
	return classes
}

// dataTypeFactory a factory function that receives a string and return an instance
// of a DataType. The string should contain the DataType value.
type dataTypeFactory func(string) (DataType, error)

// groupCodeTypes maps DXF group codes to DataTypeFactory functions. See the init
// functions for the known group codes.
var groupCodeTypes map[int]dataTypeFactory

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
