package sections

import (
	"fmt"
	"github.com/rpaloschi/dxf-go/core"
)

const absRotationBit = 0x1
const textStringBit = 0x2
const elementShapeBit = 0x4

// LineElement represents a single element in a LineType.
type LineElement struct {
	Length           float64
	AbsoluteRotation bool
	IsTextString     bool
	IsShape          bool
	ShapeNumber      int
	Scale            float64
	RotationAngle    float64
	XOffset          float64
	YOffset          float64
	Text             string
}

// LineType representation
type LineType struct {
	Name        string
	Description string
	Length      float64
	Pattern     []*LineElement
}

// NewLineType creates a new LineType object from a slice of tags.
func NewLineType(tags core.TagSlice) *LineType {
	ltype := new(LineType)
	ltype.Pattern = make([]*LineElement, 0)

	flags74 := 0
	var lineElement *LineElement
	for _, tag := range tags.RegularTags() {
		switch tag.Code {
		case 2:
			ltype.Name, _ = core.AsString(tag.Value)

		case 3:
			ltype.Description, _ = core.AsString(tag.Value)

		case 40:
			ltype.Length, _ = core.AsFloat(tag.Value)

		case 49:
			if lineElement != nil {
				ltype.Pattern = append(ltype.Pattern, lineElement)
			}
			lineElement = new(LineElement)
			lineElement.Scale = 1.0
			lineElement.Length, _ = core.AsFloat(tag.Value)

		case 74:
			flags74, _ = core.AsInt(tag.Value)
			if flags74 > 0 {
				lineElement.AbsoluteRotation = flags74&absRotationBit > 0
				lineElement.IsTextString = flags74&textStringBit > 0
				lineElement.IsShape = flags74&elementShapeBit > 0
			}

		case 75:
			flags, _ := core.AsInt(tag.Value)
			if flags74 == 0 {
				fmt.Print("WARNING! there should be no 75 Code tag if 74 value is 0\n")
			} else if lineElement.IsTextString && flags != 0 {
				fmt.Print("WARNING! Tag 75 should be 0 if 74 is a TextString\n")
			} else if lineElement.IsShape {
				lineElement.ShapeNumber = flags
			}

		case 46:
			lineElement.Scale, _ = core.AsFloat(tag.Value)

		case 50:
			lineElement.RotationAngle, _ = core.AsFloat(tag.Value)

		case 44:
			lineElement.XOffset, _ = core.AsFloat(tag.Value)

		case 45:
			lineElement.YOffset, _ = core.AsFloat(tag.Value)

		case 9:
			lineElement.Text, _ = core.AsString(tag.Value)

		default:
			fmt.Printf("Discarding tag for Style: %+v\n", tag.ToString())
		}
	}

	if lineElement != nil {
		ltype.Pattern = append(ltype.Pattern, lineElement)
	}

	return ltype
}

func NewLineTypeTable(tags core.TagSlice) (map[string]*LineType, error) {
	table := make(map[string]*LineType)

	tableSlices, err := TableEntryTags(tags)
	if err != nil {
		return table, err
	}

	for _, slice := range tableSlices {
		ltype := NewLineType(slice)
		table[ltype.Name] = ltype
	}

	return table, nil
}
