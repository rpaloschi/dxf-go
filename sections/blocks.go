package sections

import "github.com/rpaloschi/dxf-go/core"

// Block representation.
type Block struct {
	core.DxfParseable
	Name            string
	Handle          int
	LayerName       string
	SecondLayerName string
	BasePoint       core.Point
	XrefPathName    string
	Description     string
}
