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
func NewLayer(tags core.TagSlice) (*Layer, error) {
	layer := new(Layer)

	layer.On = true
	layer.Color = 7

	for _, tag := range tags {
		switch tag.Code {
		case 2:
			if name, ok := core.AsString(tag.Value); ok {
				layer.Name = name
			} else {
				return nil, fmt.Errorf("Error converting layer name tag: %v", tag.ToString())
			}

		case 70:
			if flags, ok := core.AsInt(tag.Value); ok {
				layer.Frozen = flags&frozenBit != 0
				layer.Locked = flags&lockBit != 0
			} else {
				return nil, fmt.Errorf("Error converting layer flags tag: %v", tag.ToString())
			}

		case 62:
			if color, ok := core.AsInt(tag.Value); ok {
				if color < 0 {
					layer.On = false
					layer.Color = -color
				} else {
					layer.Color = color
				}
			} else {
				return nil, fmt.Errorf("Error converting layer color tag: %v", tag.ToString())
			}

		case 6:
			if lineType, ok := core.AsString(tag.Value); ok {
				layer.LineType = lineType
			} else {
				return nil, fmt.Errorf("Error converting layer Line Type tag: %v",
					tag.ToString())
			}

		default:
			fmt.Printf("Discarding tag for Layer: %+v\n", tag.ToString())
		}
	}

	return layer, nil
}
