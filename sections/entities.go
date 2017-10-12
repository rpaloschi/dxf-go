package sections

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/rpaloschi/dxf-go/entities"
)

// EntitiesSection representation.
type EntitiesSection struct {
	Entities entities.EntitySlice
}

// Equals Compare two EntitiesSection for equality
func (e EntitiesSection) Equals(other core.DxfElement) bool {
	if otherSection, ok := other.(*EntitiesSection); ok {
		for i, entity := range e.Entities {
			otherEntity := otherSection.Entities[i]
			if !entity.Equals(otherEntity) {
				return false
			}
		}

		return true
	}

	return false
}

// NewEntitiesSection parses the EntitiesSection from a slice of tags.
func NewEntitiesSection(tags core.TagSlice) (*EntitiesSection, error) {
	section := new(EntitiesSection)

	if len(tags) == 3 {
		return section, nil
	}

	entities, err := NewEntityList(core.TagGroups(tags[2:len(tags)-1], 0))
	if err != nil {
		return nil, err
	}
	section.Entities = entities

	return section, nil
}

// NewEntityList Parses a list of tag slices into entities. returns an EntitySlice.
func NewEntityList(tags []core.TagSlice) (entities.EntitySlice, error) {
	entityList := make(entities.EntitySlice, 0)

	var accumulator *entityAccumulator
	for _, group := range tags {
		entityType := group[0].Value.ToString()

		if factory, ok := entityFactory[entityType]; ok {
			entity, err := factory(group)
			if err != nil {
				return nil, err
			}

			if accumulator != nil {
				if entity.IsSeqEnd() {
					accumulator.Stop()
					entityList = append(entityList, accumulator.parent)
					accumulator = nil
				} else {
					accumulator.entities = append(accumulator.entities, entity)
				}
			} else if entity.HasNestedEntities() {
				accumulator = newEntityAccumulator(entity)
			} else {
				entityList = append(entityList, entity)
			}
		} else {
			core.Log.Printf("Unsupported Entity Type: %v", entityType)
		}
	}

	return entityList, nil
}

type entityAccumulator struct {
	parent   entities.Entity
	entities entities.EntitySlice
}

func (e entityAccumulator) Stop() {
	e.parent.AddNestedEntities(e.entities)
}

func newEntityAccumulator(parent entities.Entity) *entityAccumulator {
	accumulator := new(entityAccumulator)

	accumulator.parent = parent
	accumulator.entities = make(entities.EntitySlice, 0)

	return accumulator
}

type entityFactoryFunc func(tags core.TagSlice) (entities.Entity, error)

var entityFactory map[string]entityFactoryFunc

func init() {
	entityFactory = map[string]entityFactoryFunc{
		"LINE": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewLine(tags)
		},
		"POINT": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewPoint(tags)
		},
		"CIRCLE": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewCircle(tags)
		},
		"ARC": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewArc(tags)
		},
		"TEXT": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewText(tags)
		},
		"INSERT": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewInsert(tags)
		},
		"SEQEND": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewSeqEnd(tags)
		},
		"POLYLINE": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewPolyline(tags)
		},
		"VERTEX": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewVertex(tags)
		},
		"LWPOLYLINE": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewLWPolyline(tags)
		},
		"ELLIPSE": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewEllipse(tags)
		},
		"SPLINE": func(tags core.TagSlice) (entities.Entity, error) {
			return entities.NewSpline(tags)
		},
	}
}
