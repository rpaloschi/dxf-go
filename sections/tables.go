package sections

import (
	"github.com/rpaloschi/dxf-go/core"
)

// Table representation.
type Table map[string]core.DxfElement

// Equals Compare two Tables for equality
func (t Table) Equals(other core.DxfElement) bool {
	if otherTable, ok := other.(Table); ok {
		if len(t) != len(otherTable) {
			return false
		}

		for key, element := range t {
			if otherElement, ok := otherTable[key]; ok {
				if !element.Equals(otherElement) {
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

// TablesSection representation
type TablesSection struct {
	Layers    Table
	Styles    Table
	LineTypes Table
}

// Equals Compare two TablesSection for equality
func (t TablesSection) Equals(other core.DxfElement) bool {
	if otherTable, ok := other.(*TablesSection); ok {
		return t.Layers.Equals(otherTable.Layers) &&
			t.Styles.Equals(otherTable.Styles) &&
			t.LineTypes.Equals(otherTable.LineTypes)
	}

	return false
}

// NewTablesSection parses the TablesSection from a slice of tags.
func NewTablesSection(tags core.TagSlice) (*TablesSection, error) {
	tables := new(TablesSection)

	tableParsers := map[string]func(slice core.TagSlice) error{
		"LAYER": func(slice core.TagSlice) error {
			layerTables, err := NewLayerTable(slice)
			tables.Layers = layerTables
			return err

		},
		"STYLE": func(slice core.TagSlice) error {
			styleTables, err := NewStyleTable(slice)
			tables.Styles = styleTables
			return err
		},
		"LTYPE": func(slice core.TagSlice) error {
			lineTypeTables, err := NewLineTypeTable(slice)
			tables.LineTypes = lineTypeTables
			return err
		},
	}

	// skip (0, 'SECTION') and (2, 'TABLES')
	tags = tags[2:]
	stopTag := core.NewTag(0, core.NewStringValue("ENDSEC"))
	endOfChunk := core.NewTag(0, core.NewStringValue("ENDTAB"))
	for _, tableTags := range SplitTagChunks(tags, stopTag, endOfChunk) {
		entryTagsList, err := TableEntryTags(tableTags)
		if err != nil {
			return nil, err
		}

		for _, entryTags := range entryTagsList {
			tableType := entryTags[0].Value.ToString()
			if tableFactory, ok := tableParsers[tableType]; ok {
				if err := tableFactory(tableTags); err != nil {
					return nil, err
				}
			} else {
				core.Log.Printf("Ignoring unknown table type: %+v\n", tableType)
			}
		}
	}

	return tables, nil
}
