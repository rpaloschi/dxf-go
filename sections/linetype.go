package sections

import (
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
	ShapeNumber      int64
	Scale            float64
	RotationAngle    float64
	XOffset          float64
	YOffset          float64
	Text             string
}

// Equals compares two LineElement objects for equality.
func (e LineElement) Equals(other LineElement) bool {
	return core.FloatEquals(e.Length, other.Length) &&
		e.AbsoluteRotation == other.AbsoluteRotation &&
		e.IsTextString == other.IsTextString &&
		e.IsShape == other.IsShape &&
		e.ShapeNumber == other.ShapeNumber &&
		core.FloatEquals(e.Scale, other.Scale) &&
		core.FloatEquals(e.RotationAngle, other.RotationAngle) &&
		core.FloatEquals(e.XOffset, other.XOffset) &&
		core.FloatEquals(e.YOffset, other.YOffset) &&
		e.Text == other.Text
}

// LineType representation
type LineType struct {
	core.DxfParseable
	Name        string
	Description string
	Length      float64
	Pattern     []*LineElement
}

// Equals compares two LineType objects for equality.
func (ltype LineType) Equals(other core.DxfElement) bool {
	if otherLtype, ok := other.(*LineType); ok {
		if ltype.Name != otherLtype.Name ||
			ltype.Description != otherLtype.Description ||
			!core.FloatEquals(ltype.Length, otherLtype.Length) ||
			len(ltype.Pattern) != len(otherLtype.Pattern) {

			return false
		}

		for i, pattern1 := range ltype.Pattern {
			pattern2 := otherLtype.Pattern[i]

			if !pattern1.Equals(*pattern2) {
				return false
			}
		}
		return true
	}
	return false
}

// NewLineType creates a new LineType object from a slice of tags.
func NewLineType(tags core.TagSlice) (*LineType, error) {
	ltype := new(LineType)
	ltype.Pattern = make([]*LineElement, 0)

	var flags74 int64
	var lineElement *LineElement

	ltype.Init(map[int]core.TypeParser{
		2:  core.NewStringTypeParserToVar(&ltype.Name),
		3:  core.NewStringTypeParserToVar(&ltype.Description),
		40: core.NewFloatTypeParserToVar(&ltype.Length),
		49: core.NewFloatTypeParser(func(length float64) {
			if lineElement != nil {
				ltype.Pattern = append(ltype.Pattern, lineElement)
			}
			lineElement = new(LineElement)
			lineElement.Scale = 1.0
			lineElement.Length = length
		}),
		74: core.NewIntTypeParser(func(flags int64) {
			flags74 = flags
			if flags74 > 0 {
				lineElement.AbsoluteRotation = flags74&absRotationBit > 0
				lineElement.IsTextString = flags74&textStringBit > 0
				lineElement.IsShape = flags74&elementShapeBit > 0
			}
		}),
		75: core.NewIntTypeParser(func(flags int64) {
			if flags74 == 0 {
				core.Log.Print("WARNING! there should be no 75 Code tag if 74 value is 0\n")
			} else if lineElement.IsTextString && flags != 0 {
				core.Log.Print("WARNING! Tag 75 should be 0 if 74 is a TextString\n")
			} else if lineElement.IsShape {
				lineElement.ShapeNumber = flags
			}
		}),
		46: core.NewFloatTypeParser(func(scale float64) {
			lineElement.Scale = scale
		}),
		50: core.NewFloatTypeParser(func(angle float64) {
			lineElement.RotationAngle = angle
		}),
		44: core.NewFloatTypeParser(func(xOffset float64) {
			lineElement.XOffset = xOffset
		}),
		45: core.NewFloatTypeParser(func(yOffset float64) {
			lineElement.YOffset = yOffset
		}),
		9: core.NewStringTypeParser(func(text string) {
			lineElement.Text = text
		}),
	})

	err := ltype.Parse(tags)

	if lineElement != nil {
		ltype.Pattern = append(ltype.Pattern, lineElement)
	}

	return ltype, err
}

// NewLineTypeTable parses the slice of tags into a table that maps the LineType name to
// the parsed LineType object.
func NewLineTypeTable(tags core.TagSlice) (Table, error) {
	table := make(Table)

	tableSlices, err := TableEntryTags(tags)
	if err != nil {
		return table, err
	}

	for _, slice := range tableSlices {
		ltype, err := NewLineType(slice)
		if err != nil {
			return nil, err
		}
		table[ltype.Name] = ltype
	}

	return table, nil
}
