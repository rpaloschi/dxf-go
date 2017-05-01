// Package dxf_go is a library to Read and Write (not yet) DXF files in Go.
//
// dxf_go contains the following packages:
//
// The core package provides the basic abstractions and utility functions to deal with DXF files..
//
// The document package provides the library's entry point and the DxfDocument representation.
//
// The entities package provides all the abstraction and code for DXF entities.
//
// The sections package provides all the abstraction and code for DXF section.
package dxf_go

// blank imports help docs.
import (
	// core package
	_ "github.com/rpaloschi/dxf-go/core"
	// document package
	_ "github.com/rpaloschi/dxf-go/document"
	// entities package
	_ "github.com/rpaloschi/dxf-go/entities"
	// sections package
	_ "github.com/rpaloschi/dxf-go/sections"
)
