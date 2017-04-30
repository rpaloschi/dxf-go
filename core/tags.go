package core

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Tag from a DXF file.
type Tag struct {
	Code  int
	Value DataType
}

// ToString returns a string representation of this Tag.
func (tag Tag) ToString() string {
	return fmt.Sprintf("{ Code: %v; Value: %v }", tag.Code, tag.Value.ToString())
}

// Equals tests equality against another Tag.
func (tag Tag) Equals(other DxfElement) bool {
	if otherTag, ok := other.(*Tag); ok {
		return tag.Code == otherTag.Code && tag.Value.Equals(otherTag.Value)
	}
	return false
}

// NewTag creates a new Tag with code and value
func NewTag(code int, value DataType) *Tag {
	tag := new(Tag)
	tag.Code = code
	tag.Value = value
	return tag
}

// NoneTag a constant that represents a nul tag.
var NoneTag = Tag{999999, &String{"NONE"}}

const appDataMarker = 102
const subclassMarker = 100

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
		if len(code) > 0 {
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

// TagSlice a slice specialization for tag pointers.
type TagSlice []*Tag

// Equals tests equality against another TagSlice.
func (slice TagSlice) Equals(other DxfElement) bool {
	if otherSlice, ok := other.(TagSlice); ok {
		if len(slice) != len(otherSlice) {
			return false
		}

		for index, tag := range slice {
			otherTag := otherSlice[index]

			if !tag.Equals(otherTag) {
				return false
			}
		}
		return true
	}
	return false
}

// TagIndex returns the index of the first occurrence of a tag with tagCode between
// the interval [startingIndex, endIndex).
// If no tag is found, it returns -1
func (slice TagSlice) TagIndex(tagCode int, startingIndex int, endIndex int) int {
	for index := startingIndex; index < endIndex; index++ {
		if slice[index].Code == tagCode {
			return index
		}
	}
	return -1
}

// AllWithCode returns a slice of tags that have the code tagCode.
func (slice TagSlice) AllWithCode(tagCode int) []*Tag {
	tags := make([]*Tag, 0)

	for _, tag := range slice {
		if tag.Code == tagCode {
			tags = append(tags, tag)
		}
	}

	return tags
}

// RegularTags returns a slice of Tags. It will return all tags
// that are not XDATA, APP_DATA or SUBCLASS.
func (slice TagSlice) RegularTags() []*Tag {
	tags := make([]*Tag, 0)

	inAppDataRange := false
	for _, tag := range slice {
		if tag.Code >= 1000 {
			continue
		}

		if tag.Code == appDataMarker {
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

// XDataTags returns a slice of Tags that contains code >= 1000.
func (slice TagSlice) XDataTags() []*Tag {
	tags := make([]*Tag, 0)

	for _, tag := range slice {
		if tag.Code > 999 {
			tags = append(tags, tag)
		}
	}

	return tags
}

// AppDataTags returns a slice of tags containing all App Data Tags.
func (slice TagSlice) AppDataTags() map[string][]*Tag {
	appData := make(map[string][]*Tag)
	appTags := make([]*Tag, 0)

	for _, tag := range slice {
		if tag.Code == appDataMarker {
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

// SubclassesTags returns a slice of tags containing all Subclass Tags.
func (slice TagSlice) SubclassesTags() map[string][]*Tag {
	classes := make(map[string][]*Tag)
	tags := make([]*Tag, 0)
	name := "noname"

	for _, tag := range slice.RegularTags() {
		if tag.Code == subclassMarker {
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

// TagGroups splits a TagSlice into Groups of TagSlices starting with a Split Tag
// and ending before the next Split Tag.
// A Split Tag is a tag with Code == splitCode, like (0, 'SECTION') for splitCode = 0.
func TagGroups(tags TagSlice, splitCode int) []TagSlice {
	groups := make([]TagSlice, 0)

	group := make(TagSlice, 0)
	for _, tag := range tags {
		if tag.Code == splitCode {
			if len(group) > 0 {
				groups = append(groups, group)
				group = make(TagSlice, 0)
			}
			group = append(group, tag)
		} else if len(group) > 0 {
			group = append(group, tag)
		}
	}

	if len(group) > 0 && group[0].Code == splitCode {
		groups = append(groups, group)
	}

	return groups
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
