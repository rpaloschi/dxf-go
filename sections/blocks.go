package sections

import (
	"github.com/rpaloschi/dxf-go/core"
	"github.com/rpaloschi/dxf-go/entities"
)

// Block representation.
type Block struct {
	core.DxfParseable
	Name         string
	Handle       string
	LayerName    string
	SecondName   string
	BasePoint    core.Point
	XrefPathName string
	Description  string
	Entities     entities.EntitySlice
}

// Equals compare with the other block for equality.
func (b Block) Equals(other core.DxfElement) bool {
	if otherBlock, ok := other.(*Block); ok {
		return b.Name == otherBlock.Name &&
			b.Handle == otherBlock.Handle &&
			b.LayerName == otherBlock.LayerName &&
			b.SecondName == otherBlock.SecondName &&
			b.BasePoint.Equals(otherBlock.BasePoint) &&
			b.XrefPathName == otherBlock.XrefPathName &&
			b.Description == otherBlock.Description &&
			b.Entities.Equals(otherBlock.Entities)
	}
	return false
}

// NewBlock builds a new Block from a slice of Tags.
func NewBlock(tags core.TagSlice) (*Block, error) {
	block := new(Block)

	block.Init(map[int]core.TypeParser{
		1:  core.NewStringTypeParserToVar(&block.XrefPathName),
		2:  core.NewStringTypeParserToVar(&block.Name),
		3:  core.NewStringTypeParserToVar(&block.SecondName),
		4:  core.NewStringTypeParserToVar(&block.Description),
		5:  core.NewStringTypeParserToVar(&block.Handle),
		8:  core.NewStringTypeParserToVar(&block.LayerName),
		10: core.NewFloatTypeParserToVar(&block.BasePoint.X),
		20: core.NewFloatTypeParserToVar(&block.BasePoint.Y),
		30: core.NewFloatTypeParserToVar(&block.BasePoint.Z),
	})

	err := block.Parse(tags)
	return block, err
}

// BlocksSection BLOCKS section representation.
type BlocksSection map[string]*Block

// Equals Compare with the other BlocksSection for equality.
func (b BlocksSection) Equals(other BlocksSection) bool {
	if len(b) != len(other) {
		return false
	}

	for i, block := range b {
		otherBlock := other[i]

		if !block.Equals(otherBlock) {
			return false
		}
	}

	return true
}

// NewBlocksSection creates a new BlocksSection from a slice of tags.
func NewBlocksSection(tags core.TagSlice) (BlocksSection, error) {
	blocks := make(BlocksSection)

	if len(tags) > 3 {
		groups := make([]core.TagSlice, 0)
		tagGroups := core.TagGroups(tags[2:len(tags)-1], 0)
		for _, group := range tagGroups {
			if group[0].Value.ToString() == "ENDBLK" {
				block, err := NewBlock(groups[0])
				if err != nil {
					return nil, err
				}

				allEntitites, err := NewEntityList(groups[1:])
				if err != nil {
					return nil, err
				}

				block.Entities = allEntitites
				blocks[block.Name] = block
				groups = make([]core.TagSlice, 0)
			} else {
				groups = append(groups, group)
			}
		}

	}

	return blocks, nil
}
