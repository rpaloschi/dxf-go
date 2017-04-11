package sections

import (
	"fmt"
	"github.com/rpaloschi/dxf-go/core"
)

// TablesSection representation
type TablesSection struct {
	LayerTable    map[string]*Layer
	StyleTable    map[string]*Style
	LineTypeTable map[string]*LineType
}

// Equals Compare two TablesSection for equality
func (t TablesSection) Equals(other TablesSection) bool {
	if len(t.LayerTable) != len(other.LayerTable) ||
		len(t.StyleTable) != len(other.StyleTable) ||
		len(t.LineTypeTable) != len(other.LineTypeTable) {

		return false
	}

	for key, layer := range t.LayerTable {
		if otherLayer, ok := other.LayerTable[key]; ok {
			if !layer.Equals(*otherLayer) {
				return false
			}
		} else {
			return false
		}
	}

	for key, style := range t.StyleTable {
		if otherStyle, ok := other.StyleTable[key]; ok {
			if !style.Equals(*otherStyle) {
				return false
			}
		} else {
			return false
		}
	}

	for key, ltype := range t.LineTypeTable {
		if otherLType, ok := other.LineTypeTable[key]; ok {
			if !ltype.Equals(*otherLType) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// NewTablesSection parses the TablesSection from a slice of tags.
func NewTablesSection(tags core.TagSlice) (*TablesSection, error) {
	tables := new(TablesSection)

	tableParsers := map[string]func(slice core.TagSlice) error{
		"LAYER": func(slice core.TagSlice) error {
			layerTables, err := NewLayerTable(slice)
			tables.LayerTable = layerTables
			return err

		},
		"STYLE": func(slice core.TagSlice) error {
			styleTables, err := NewStyleTable(slice)
			tables.StyleTable = styleTables
			return err
		},
		"LTYPE": func(slice core.TagSlice) error {
			lineTypeTables, err := NewLineTypeTable(slice)
			tables.LineTypeTable = lineTypeTables
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
				fmt.Printf("Ignoring unknown table type: %+v\n", tableType)
			}
		}
	}

	return tables, nil
}
