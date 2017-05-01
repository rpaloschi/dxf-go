package sections

import "github.com/rpaloschi/dxf-go/core"

const tagACADVER = "$ACADVER"
const tagDWGCODEPAGE = "$DWGCODEPAGE"

// HeaderSection representation.
type HeaderSection struct {
	Values map[string]core.TagSlice
}

// Equals Compares two HeaderSections for equality.
// If other cannot be casted to a HeaderSection, returns false.
func (section HeaderSection) Equals(other core.DxfElement) bool {
	if otherSection, ok := other.(*HeaderSection); ok {
		if len(section.Values) != len(otherSection.Values) {
			return false
		}

		for key, slice := range section.Values {
			if otherSlice, ok := otherSection.Values[key]; ok {
				if !slice.Equals(otherSlice) {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}
	return false
}

// NewHeaderSection creates a new *HeaderSection from a core.TagSlice.
func NewHeaderSection(tags core.TagSlice) *HeaderSection {
	header := new(HeaderSection)
	header.Values = make(map[string]core.TagSlice)

	if len(tags) > 3 {
		groups := core.TagGroups(tags[2:len(tags)-1], 9)
		for _, group := range groups {
			var groupTags core.TagSlice
			headerKey := group[0].Value.ToString()

			if keyTags, ok := header.Values[headerKey]; ok {
				groupTags = keyTags
			} else {
				groupTags = make(core.TagSlice, 0)
			}

			groupTags = append(groupTags, group[1:]...)
			header.Values[headerKey] = groupTags
		}
	}

	// default values
	if _, ok := header.Values[tagACADVER]; !ok {
		header.Values[tagACADVER] = core.TagSlice{
			core.NewTag(1, core.NewStringValue("AC1009")),
		}
	}
	if _, ok := header.Values[tagDWGCODEPAGE]; !ok {
		header.Values[tagDWGCODEPAGE] = core.TagSlice{
			core.NewTag(3, core.NewStringValue("ANSI_1252")),
		}
	}

	return header
}

// Get a slice of core.Tags by its key on the HeaderSection.
func (section *HeaderSection) Get(key string) core.TagSlice {
	if keyTags, ok := section.Values[key]; ok {
		return keyTags
	}
	return core.TagSlice{}
}
