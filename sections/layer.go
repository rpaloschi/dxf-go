package sections

import (
	"github.com/rpaloschi/dxf-go/core"
)

const lockBit = 0x4
const frozenBit = 0x1

// Layer representation.
type Layer struct {
	core.DxfElement
	Name     string
	Color    int
	LineType string
	Locked   bool
	Frozen   bool
	On       bool
}

// NewLayer builds a new Layer from a tag slice.
func NewLayer(tags core.TagSlice) (*Layer, error) {
	layer := new(Layer)

	layer.On = true
	layer.Color = 7

	layer.Init(map[int]core.TypeParser{
		2: core.NewStringTypeParserToVar(&layer.Name),
		70: core.NewIntTypeParser(func(flags int) {
			layer.Frozen = flags&frozenBit != 0
			layer.Locked = flags&lockBit != 0
		}),
		62: core.NewIntTypeParser(func(color int) {
			if color < 0 {
				layer.On = false
				layer.Color = -color
			} else {
				layer.Color = color
			}
		}),
		6: core.NewStringTypeParserToVar(&layer.LineType),
	})

	err := layer.Parse(tags)
	return layer, err
}

// NewLayerTable parses the slice of tags into a table that maps the layer name to
// the parsed Layer object.
func NewLayerTable(tags core.TagSlice) (map[string]*Layer, error) {
	table := make(map[string]*Layer)

	tableSlices, err := TableEntryTags(tags)
	if err != nil {
		return table, err
	}

	for _, slice := range tableSlices {
		layer, err := NewLayer(slice)
		if err != nil {
			return nil, err
		}
		table[layer.Name] = layer
	}

	return table, nil
}

// TODO:
// 290 Plotting flag. If set to 0, do not plot this layer
// 370 Lineweight enum value
