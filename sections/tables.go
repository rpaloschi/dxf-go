package sections

import (
	"fmt"
	"github.com/rpaloschi/dxf-go/core"
)

type StringMappedTable interface {
	Keys() []string
	Get(key string) (core.DxfElement, bool)
}

type LayerTable map[string]*Layer

func (table LayerTable) Keys() []string {
	keys := make([]string, len(table))
	i := 0
	for k := range table {
		keys[i] = k
		i++
	}
	return keys
}

func (table LayerTable) Get(key string) (core.DxfElement, bool) {
	element, ok := table[key]
	return element, ok
}

type StyleTable map[string]*Style

func (table StyleTable) Keys() []string {
	keys := make([]string, len(table))
	i := 0
	for k := range table {
		keys[i] = k
		i++
	}
	return keys
}

func (table StyleTable) Get(key string) (core.DxfElement, bool) {
	element, ok := table[key]
	return element, ok
}

type LineTypeTable map[string]*LineType

func (table LineTypeTable) Keys() []string {
	keys := make([]string, len(table))
	i := 0
	for k := range table {
		keys[i] = k
		i++
	}
	return keys
}

func (table LineTypeTable) Get(key string) (core.DxfElement, bool) {
	element, ok := table[key]
	return element, ok
}

func StringMappedTablesAreEquals(tableA StringMappedTable, tableB StringMappedTable) bool {
	keysA := tableA.Keys()
	keysB := tableB.Keys()

	if len(keysA) != len(keysB) {
		return false
	}

	for _, key := range keysA {
		elementA, _ := tableA.Get(key)
		if elementB, ok := tableB.Get(key); ok {
			if !elementA.Equals(elementB) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// TablesSection representation
type TablesSection struct {
	Layers    LayerTable
	Styles    StyleTable
	LineTypes LineTypeTable
}

// Equals Compare two TablesSection for equality
func (t TablesSection) Equals(other core.DxfElement) bool {
	if otherTable, ok := other.(*TablesSection); ok {
		return StringMappedTablesAreEquals(t.Layers, otherTable.Layers) &&
			StringMappedTablesAreEquals(t.Styles, otherTable.Styles) &&
			StringMappedTablesAreEquals(t.LineTypes, otherTable.LineTypes)
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
				fmt.Printf("Ignoring unknown table type: %+v\n", tableType)
			}
		}
	}

	return tables, nil
}
