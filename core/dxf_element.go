package core

// DxfElement is the common interface for all Dxf elements.
// Common operations should be part of this interface.
type DxfElement interface {
	Equals(other DxfElement) bool
}
