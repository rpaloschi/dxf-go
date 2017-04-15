package sections

import (
	"errors"
	"github.com/rpaloschi/dxf-go/core"
)

// TableEntryTags splits a slice of tags that contains the TABLES entry of a DXF file, validates
// its basic structure and returns only the tags to be parsed as tables. (Removes the
// initial and final markup tags for table).
func TableEntryTags(tags core.TagSlice) ([]core.TagSlice, error) {
	groups := core.TagGroups(tags, 0)
	lastIndex := len(groups) - 1
	first := groups[0][0].Value.ToString()
	last := groups[lastIndex][0].Value.ToString()

	if first != "TABLE" || last != "ENDTAB" {
		return []core.TagSlice{},
			errors.New("Invalid table. Missing TABLE AND/OR ENDTAB tags.")
	}

	return groups[1:lastIndex], nil
}

// SplitTagChunks splits a TagSlice into a series of TagSlices delimited by chunkDelimiter tags.
// The Iteration ends at stopTag.
func SplitTagChunks(tags core.TagSlice, stopTag *core.Tag, chunkDelimiter *core.Tag) []core.TagSlice {
	chunks := make([]core.TagSlice, 0)

	tagIndex := 0
	for tagIndex < len(tags) {
		if tags[tagIndex].Equals(stopTag) {
			break
		}

		chunk := make([]*core.Tag, 1)
		chunk[0] = tags[tagIndex]
		tagIndex++

		foundStop := false

		for {
			if tags[tagIndex].Equals(chunkDelimiter) {
				chunk = append(chunk, tags[tagIndex])
				tagIndex++
				break
			}
			if tags[tagIndex].Equals(stopTag) {
				foundStop = true
				tagIndex++
				break
			} else {
				chunk = append(chunk, tags[tagIndex])
				tagIndex++
			}
		}
		chunks = append(chunks, chunk)

		if foundStop {
			break
		}
	}

	return chunks
}
