package sections

import "github.com/rpaloschi/dxf-go/core"

const tagACADVER = "$ACADVER"
const tagDWGCODEPAGE = "$DWGCODEPAGE"

// HeaderSection representation.
type HeaderSection struct {
	values map[string]core.TagSlice
}

// Equals Compares two HeaderSections for equality.
// If other cannot be casted to a HeaderSection, returns false.
func (section HeaderSection) Equals(other core.DxfElement) bool {
	if otherSection, ok := other.(*HeaderSection); ok {
		if len(section.values) != len(otherSection.values) {
			return false
		}

		for key, slice := range section.values {
			if otherSlice, ok := otherSection.values[key]; ok {
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
	header.values = make(map[string]core.TagSlice)

	if len(tags) > 3 {
		groups := core.TagGroups(tags[2:len(tags)-1], 9)
		for _, group := range groups {
			var groupTags core.TagSlice
			headerKey := group[0].Value.ToString()

			if keyTags, ok := header.values[headerKey]; ok {
				groupTags = keyTags
			} else {
				groupTags = make(core.TagSlice, 0)
			}

			groupTags = append(groupTags, group[1:]...)
			header.values[headerKey] = groupTags
		}
	}

	// default values
	if _, ok := header.values[tagACADVER]; !ok {
		header.values[tagACADVER] = core.TagSlice{
			core.NewTag(1, core.NewStringValue("AC1009")),
		}
	}
	if _, ok := header.values[tagDWGCODEPAGE]; !ok {
		header.values[tagDWGCODEPAGE] = core.TagSlice{
			core.NewTag(3, core.NewStringValue("ANSI_1252")),
		}
	}

	return header
}

// Get a slice of core.Tags by its key on the HeaderSection.
func (section *HeaderSection) Get(key string) core.TagSlice {
	if keyTags, ok := section.values[key]; ok {
		return keyTags
	}
	return core.TagSlice{}
}
