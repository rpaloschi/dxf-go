package entities

import (
	"github.com/rpaloschi/dxf-go/core"
)

// Insert Entity representation
type Insert struct {
	BaseEntity
	BlockName          string
	InsertionPoint     core.Point
	ScaleFactorX       float64
	ScaleFactorY       float64
	ScaleFactorZ       float64
	RotationAngle      float64
	ColumnCount        int
	RowCount           int
	ColumnSpacing      float64
	RowSpacing         float64
	AttributesFollow   bool
	Entities           EntitySlice
	ExtrusionDirection core.Point
}

// Equals tests equality against another Insert.
func (i Insert) Equals(other core.DxfElement) bool {
	if otherInsert, ok := other.(*Insert); ok {
		return i.BaseEntity.Equals(otherInsert.BaseEntity) &&
			i.BlockName == otherInsert.BlockName &&
			i.InsertionPoint.Equals(otherInsert.InsertionPoint) &&
			core.FloatEquals(i.ScaleFactorX, otherInsert.ScaleFactorX) &&
			core.FloatEquals(i.ScaleFactorY, otherInsert.ScaleFactorY) &&
			core.FloatEquals(i.ScaleFactorZ, otherInsert.ScaleFactorZ) &&
			core.FloatEquals(i.RotationAngle, otherInsert.RotationAngle) &&
			i.ColumnCount == otherInsert.ColumnCount &&
			i.RowCount == otherInsert.RowCount &&
			core.FloatEquals(i.ColumnSpacing, otherInsert.ColumnSpacing) &&
			core.FloatEquals(i.RowSpacing, otherInsert.RowSpacing) &&
			i.AttributesFollow == otherInsert.AttributesFollow &&
			i.Entities.Equals(otherInsert.Entities) &&
			i.ExtrusionDirection.Equals(otherInsert.ExtrusionDirection)
	}
	return false
}

// HasNestedEntities a Insert will have nested entities if the AttributesFollow
// value is set.
func (i Insert) HasNestedEntities() bool {
	return i.AttributesFollow
}

// AddNestedEntities Add a slice of Entities as nested Entities.
func (i *Insert) AddNestedEntities(entities EntitySlice) {
	i.Entities = entities
}

// NewInsert builds a new Insert from a slice of Tags.
func NewInsert(tags core.TagSlice) (*Insert, error) {
	insert := new(Insert)

	// set defaults
	insert.ScaleFactorX = 1.0
	insert.ScaleFactorY = 1.0
	insert.ScaleFactorZ = 1.0
	insert.RotationAngle = 0.0
	insert.ColumnCount = 1
	insert.RowCount = 1
	insert.ColumnSpacing = 0.0
	insert.RowSpacing = 0.0
	insert.Entities = make(EntitySlice, 0)
	insert.ExtrusionDirection = core.Point{X: 0.0, Y: 0.0, Z: 1.0}

	insert.InitBaseEntityParser()
	insert.Update(map[int]core.TypeParser{
		2:  core.NewStringTypeParserToVar(&insert.BlockName),
		10: core.NewFloatTypeParserToVar(&insert.InsertionPoint.X),
		20: core.NewFloatTypeParserToVar(&insert.InsertionPoint.Y),
		30: core.NewFloatTypeParserToVar(&insert.InsertionPoint.Z),
		41: core.NewFloatTypeParserToVar(&insert.ScaleFactorX),
		42: core.NewFloatTypeParserToVar(&insert.ScaleFactorY),
		43: core.NewFloatTypeParserToVar(&insert.ScaleFactorZ),
		44: core.NewFloatTypeParserToVar(&insert.ColumnSpacing),
		45: core.NewFloatTypeParserToVar(&insert.RowSpacing),
		50: core.NewFloatTypeParserToVar(&insert.RotationAngle),
		66: core.NewIntTypeParser(func(value int) {
			insert.AttributesFollow = value == 1
		}),
		70:  core.NewIntTypeParserToVar(&insert.ColumnCount),
		71:  core.NewIntTypeParserToVar(&insert.RowCount),
		210: core.NewFloatTypeParserToVar(&insert.ExtrusionDirection.X),
		220: core.NewFloatTypeParserToVar(&insert.ExtrusionDirection.Y),
		230: core.NewFloatTypeParserToVar(&insert.ExtrusionDirection.Z),
	})

	err := insert.Parse(tags)
	return insert, err
}
