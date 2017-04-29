package entities

import "github.com/rpaloschi/dxf-go/core"

// SeqEnd Entity representation
type SeqEnd struct {
	BaseEntity
}

// Equals tests equality against another SeqEnd.
func (c SeqEnd) Equals(other core.DxfElement) bool {
	if otherSeqEnd, ok := other.(*SeqEnd); ok {
		return c.BaseEntity.Equals(otherSeqEnd.BaseEntity)
	}
	return false
}

// IsSeqEnd should return true only for this type.
func (c SeqEnd) IsSeqEnd() bool {
	return true
}

// NewSeqEnd builds a new SeqEnd from a slice of Tags.
func NewSeqEnd(tags core.TagSlice) (*SeqEnd, error) {
	point := new(SeqEnd)

	point.InitBaseEntityParser()

	err := point.Parse(tags)
	return point, err
}
