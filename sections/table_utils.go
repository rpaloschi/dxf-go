package sections

import (
	"errors"
	"github.com/rpaloschi/dxf-go/core"
)

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

func SplitTagChunks(tags core.TagSlice, stopTag *core.Tag, chunkDelimiter *core.Tag) []core.TagSlice {
	chunks := make([]core.TagSlice, 0)

	tagIndex := 0
	for tagIndex < len(tags) {
		if tags[tagIndex].Equals(*stopTag) {
			break
		}

		chunk := make([]*core.Tag, 1)
		chunk[0] = tags[tagIndex]
		tagIndex++

		foundStop := false

		for {
			if tags[tagIndex].Equals(*chunkDelimiter) {
				chunk = append(chunk, tags[tagIndex])
				tagIndex++
				break
			}
			if tags[tagIndex].Equals(*stopTag) {
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
