package sections

import (
	"fmt"
	"github.com/rpaloschi/dxf-go/core"
)

const lockBit = 0x4
const frozenBit = 0x1

// Layer representation.
type Layer struct {
	Name     string
	Color    int
	LineType string
	Locked   bool
	Frozen   bool
	On       bool
}

// NewLayer builds a new Layer from a tag slice.
func NewLayer(tags core.TagSlice) *Layer {
	layer := new(Layer)

	layer.On = true
	layer.Color = 7

	for _, tag := range tags.RegularTags() {
		switch tag.Code {
		case 2:
			layer.Name, _ = core.AsString(tag.Value)
		case 70:
			flags, _ := core.AsInt(tag.Value)
			layer.Frozen = flags&frozenBit != 0
			layer.Locked = flags&lockBit != 0

		case 62:
			color, _ := core.AsInt(tag.Value)
			if color < 0 {
				layer.On = false
				layer.Color = -color
			} else {
				layer.Color = color
			}

		case 6:
			layer.LineType, _ = core.AsString(tag.Value)

		default:
			fmt.Printf("Discarding tag for Layer: %+v\n", tag.ToString())
		}
	}

	return layer
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
		layer := NewLayer(slice)
		table[layer.Name] = layer
	}

	return table, nil
}

// TODO:
// 290 Plotting flag. If set to 0, do not plot this layer
// 370 Lineweight enum value
