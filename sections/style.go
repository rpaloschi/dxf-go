package sections

import (
	"github.com/rpaloschi/dxf-go/core"
)

const verticalTextBit = 0x4
const shapeBit = 0x1

const backwardsBit = 0x2
const upsideDownBit = 0x4

// Style Table representation
type Style struct {
	core.DxfElement
	Name           string
	Height         float64
	Width          float64
	Oblique        float64
	IsBackwards    bool
	IsUpsideDown   bool
	IsShape        bool
	IsVerticalText bool
	Font           string
	BigFont        string
}

// Equals compares two Style objects for equality.
func (style Style) Equals(other Style) bool {
	return style.Name == other.Name &&
		core.FloatEquals(style.Height, other.Height) &&
		core.FloatEquals(style.Width, other.Width) &&
		core.FloatEquals(style.Oblique, other.Oblique) &&
		style.IsBackwards == other.IsBackwards &&
		style.IsUpsideDown == other.IsUpsideDown &&
		style.IsShape == other.IsShape &&
		style.IsVerticalText == other.IsVerticalText &&
		style.Font == other.Font &&
		style.BigFont == other.BigFont
}

// NewStyle creates a new Style object from a slice of tags.
func NewStyle(tags core.TagSlice) (*Style, error) {
	style := new(Style)

	style.Height = 1.0
	style.Width = 1.0

	style.Init(map[int]core.TypeParser{
		2:  core.NewStringTypeParserToVar(&style.Name),
		3:  core.NewStringTypeParserToVar(&style.Font),
		4:  core.NewStringTypeParserToVar(&style.BigFont),
		40: core.NewFloatTypeParserToVar(&style.Height),
		41: core.NewFloatTypeParserToVar(&style.Width),
		50: core.NewFloatTypeParserToVar(&style.Oblique),
		70: core.NewIntTypeParser(func(flags int) {
			style.IsShape = flags&shapeBit != 0
			style.IsVerticalText = flags&verticalTextBit != 0
		}),
		71: core.NewIntTypeParser(func(flags int) {
			style.IsBackwards = flags&backwardsBit != 0
			style.IsUpsideDown = flags&upsideDownBit != 0
		}),
	})

	err := style.Parse(tags)
	return style, err
}

// NewStyleTable parses the slice of tags into a table that maps the Style name to
// the parsed Style object.
func NewStyleTable(tags core.TagSlice) (map[string]*Style, error) {
	table := make(map[string]*Style)

	tableSlices, err := TableEntryTags(tags)
	if err != nil {
		return table, err
	}

	for _, slice := range tableSlices {
		style, err := NewStyle(slice)
		if err != nil {
			return nil, err
		}
		table[style.Name] = style
	}

	return table, nil
}
