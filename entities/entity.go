package entities

import (
	"github.com/rpaloschi/dxf-go/core"
)

// Entity all entities should implement this interface.
type Entity interface {
	core.DxfElement
	IsSeqEnd() bool
	HasNestedEntities() bool
	AddNestedEntities(entities EntitySlice)
}

// RegularEntity most of the Entities will return the same values for
// those APIs, this creates a default implementation that will be added to
// the BaseEntity class so that only that classes that need a different
// implementation need to overwrite it.
type RegularEntity struct{}

// IsSeqEnd a RegularEntity is not a SeqEnd entity.
func (r RegularEntity) IsSeqEnd() bool {
	return false
}

// HasNestedEntities a RegularEntity has no NestedEntities.
func (r RegularEntity) HasNestedEntities() bool {
	return false
}

// AddNestedEntities a default empty implementation just to implement the interface.
func (r RegularEntity) AddNestedEntities(entities EntitySlice) {}

// EntitySlice a slice of Entity objects.
type EntitySlice []Entity

// Equals compares the EntitySlice to the other for equality.
func (e EntitySlice) Equals(other EntitySlice) bool {
	if len(e) != len(other) {
		return false
	}

	for i, entity := range e {
		otherEntity := other[i]

		if !entity.Equals(otherEntity) {
			return false
		}
	}

	return true
}

// Space the Entity Space
type Space int

const (
	MODEL Space = iota
	PAPER
)

// ShadowMode for the Entity. How should shadows be handled for this Entity.
type ShadowMode int

const (
	CASTS_AND_RECEIVE ShadowMode = iota
	CASTS
	RECEIVES
	IGNORES
)

// BaseEntity holds the common part for all Entity types.
// New Entity types should be composed by it.
type BaseEntity struct {
	core.DxfParseable
	RegularEntity
	Handle        string
	Owner         string
	Space         Space
	LayoutTabName string
	LayerName     string
	LineTypeName  string
	On            bool
	Color         int
	LineWeight    int
	LineTypeScale float64
	Visible       bool
	TrueColor     core.TrueColor
	ColorName     string
	Transparency  int
	ShadowMode    ShadowMode
}

// Equals compare two BaseEntity objects for equality.
// It does not implements DxfElement by design, meaning that the composed
// Entity structs should do.
func (entity BaseEntity) Equals(other BaseEntity) bool {
	return entity.Handle == other.Handle &&
		entity.Owner == other.Owner &&
		entity.Space == other.Space &&
		entity.LayoutTabName == other.LayoutTabName &&
		entity.LayerName == other.LayerName &&
		entity.LineTypeName == other.LineTypeName &&
		entity.On == other.On &&
		entity.Color == other.Color &&
		entity.LineWeight == other.LineWeight &&
		core.FloatEquals(entity.LineTypeScale, other.LineTypeScale) &&
		entity.Visible == other.Visible &&
		entity.TrueColor == other.TrueColor &&
		entity.ColorName == other.ColorName &&
		entity.Transparency == other.Transparency &&
		entity.ShadowMode == other.ShadowMode
}

// InitBaseEntityParser Inits the EntityParsers for the BaseEntity attributes.
func (entity *BaseEntity) InitBaseEntityParser() {
	entity.On = true
	entity.Visible = true
	entity.Init(map[int]core.TypeParser{
		5:  core.NewStringTypeParserToVar(&entity.Handle),
		6:  core.NewStringTypeParserToVar(&entity.LineTypeName),
		8:  core.NewStringTypeParserToVar(&entity.LayerName),
		48: core.NewFloatTypeParserToVar(&entity.LineTypeScale),
		60: core.NewIntTypeParser(func(value int) {
			entity.Visible = value == 0
		}),
		62: core.NewIntTypeParser(func(value int) {
			if value < 0 {
				entity.On = false
				value = -value
			}
			entity.Color = value
		}),
		67: core.NewIntTypeParser(func(value int) {
			entity.Space = Space(value)
		}),
		284: core.NewIntTypeParser(func(value int) {
			entity.ShadowMode = ShadowMode(value)
		}),
		330: core.NewStringTypeParserToVar(&entity.Owner),
		370: core.NewIntTypeParserToVar(&entity.LineWeight),
		410: core.NewStringTypeParserToVar(&entity.LayoutTabName),

		420: core.NewIntTypeParser(func(value int) {
			entity.TrueColor = core.TrueColor(value)
		}),
		430: core.NewStringTypeParserToVar(&entity.ColorName),
		440: core.NewIntTypeParserToVar(&entity.Transparency),
	})
}
