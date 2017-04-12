package core

// DxfElement is the common interface for all Dxf elements.
type DxfElement interface {
	Equals(other DxfElement) bool
}
