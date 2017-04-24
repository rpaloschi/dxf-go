package entities

import "github.com/rpaloschi/dxf-go/core"

// HorizontalTextJustification Horizontal Text Justification type
type HorizontalTextJustification int

const (
	HTEXT_LEFT HorizontalTextJustification = iota
	HTEXT_CENTER
	HTEXT_RIGHT
	HTEXT_ALIGNED
	HTEXT_MIDDLE
	HTEXT_FIT
)

// VerticalTextJustification Vertical Text Justification type
type VerticalTextJustification int

const (
	VTEXT_BASELINE VerticalTextJustification = iota
	VTEXT_BOTTOM
	VTEXT_MIDDLE
	VTEXT_TOP
)

// Text Entity representation
type Text struct {
	BaseEntity
	Thickness               float64
	FirstAlignmentPoint     core.Point
	Height                  float64
	Value                   string
	Rotation                float64
	RelativeXScale          float64
	ObliqueAngle            float64
	StyleName               string
	MirroredX               bool
	MirroredY               bool
	HorizontalJustification HorizontalTextJustification
	SecondAlignmentPoint    core.Point
	ExtrusionDirection      core.Point
	VerticalJustification   VerticalTextJustification
}

// Equals tests equality against another Text.
func (e Text) Equals(other core.DxfElement) bool {
	if otherText, ok := other.(*Text); ok {
		return e.BaseEntity.Equals(otherText.BaseEntity) &&
			core.FloatEquals(e.Thickness, otherText.Thickness) &&
			e.FirstAlignmentPoint.Equals(otherText.FirstAlignmentPoint) &&
			core.FloatEquals(e.Height, otherText.Height) &&
			e.Value == otherText.Value &&
			core.FloatEquals(e.Rotation, otherText.Rotation) &&
			core.FloatEquals(e.RelativeXScale, otherText.RelativeXScale) &&
			core.FloatEquals(e.ObliqueAngle, otherText.ObliqueAngle) &&
			e.StyleName == otherText.StyleName &&
			e.MirroredX == otherText.MirroredX &&
			e.MirroredY == otherText.MirroredY &&
			e.HorizontalJustification == otherText.HorizontalJustification &&
			e.SecondAlignmentPoint.Equals(otherText.SecondAlignmentPoint) &&
			e.ExtrusionDirection.Equals(otherText.ExtrusionDirection) &&
			e.VerticalJustification == otherText.VerticalJustification
	}
	return false
}

const backwardTextBit = 0x2
const upsideDownTextBit = 0x4

// NewText builds a new Text from a slice of Tags.
func NewText(tags core.TagSlice) (*Text, error) {
	text := new(Text)

	// set default
	text.RelativeXScale = 1.0
	text.StyleName = "STANDARD"
	text.ExtrusionDirection = core.Point{X: 0.0, Y: 0.0, Z: 1.0}

	text.InitBaseEntityParser()
	text.Update(map[int]core.TypeParser{
		1:  core.NewStringTypeParserToVar(&text.Value),
		7:  core.NewStringTypeParserToVar(&text.StyleName),
		10: core.NewFloatTypeParserToVar(&text.FirstAlignmentPoint.X),
		20: core.NewFloatTypeParserToVar(&text.FirstAlignmentPoint.Y),
		30: core.NewFloatTypeParserToVar(&text.FirstAlignmentPoint.Z),
		39: core.NewFloatTypeParserToVar(&text.Thickness),
		40: core.NewFloatTypeParserToVar(&text.Height),
		41: core.NewFloatTypeParserToVar(&text.RelativeXScale),
		50: core.NewFloatTypeParserToVar(&text.Rotation),
		51: core.NewFloatTypeParserToVar(&text.ObliqueAngle),
		71: core.NewIntTypeParser(func(flags int) {
			text.MirroredX = flags&backwardTextBit != 0
			text.MirroredY = flags&upsideDownTextBit != 0
		}),
		72: core.NewIntTypeParser(func(value int) {
			text.HorizontalJustification = HorizontalTextJustification(value)
		}),
		73: core.NewIntTypeParser(func(value int) {
			text.VerticalJustification = VerticalTextJustification(value)
		}),
		11:  core.NewFloatTypeParserToVar(&text.SecondAlignmentPoint.X),
		21:  core.NewFloatTypeParserToVar(&text.SecondAlignmentPoint.Y),
		31:  core.NewFloatTypeParserToVar(&text.SecondAlignmentPoint.Z),
		210: core.NewFloatTypeParserToVar(&text.ExtrusionDirection.X),
		220: core.NewFloatTypeParserToVar(&text.ExtrusionDirection.Y),
		230: core.NewFloatTypeParserToVar(&text.ExtrusionDirection.Z),
	})

	err := text.Parse(tags)
	return text, err
}
