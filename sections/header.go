package sections

import "github.com/rpaloschi/dxf-go/core"

const TAG_ACADVER = "$ACADVER"
const TAG_DWGCODEPAGE = "$DWGCODEPAGE"

type HeaderSection struct {
	values map[string][]*core.Tag
}

func NewHeaderSection(tags core.TagSlice) *HeaderSection {
	header := new(HeaderSection)
	header.values = make(map[string][]*core.Tag)

	if len(tags) > 3 {
		groups := core.TagGroups(tags[2:len(tags)-1], 9)
		for _, group := range groups {
			var groupTags []*core.Tag
			headerKey := group[0].Value.ToString()

			if keyTags, ok := header.values[headerKey]; ok {
				groupTags = keyTags
			} else {
				groupTags = make([]*core.Tag, 0)
			}

			groupTags = append(groupTags, group[1:]...)
			header.values[headerKey] = groupTags
		}
	}

	// default values
	if _, ok := header.values[TAG_ACADVER]; !ok {
		header.values[TAG_ACADVER] = []*core.Tag{
			core.NewTag(1, core.NewStringValue("AC1009")),
		}
	}
	if _, ok := header.values[TAG_DWGCODEPAGE]; !ok {
		header.values[TAG_DWGCODEPAGE] = []*core.Tag{
			core.NewTag(3, core.NewStringValue("ANSI_1252")),
		}
	}

	return header
}

func (section *HeaderSection) Get(key string) []*core.Tag {
	if keyTags, ok := section.values[key]; ok {
		return keyTags
	}
	return []*core.Tag{}
}
