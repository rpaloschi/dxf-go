[![Build Status](https://travis-ci.org/rpaloschi/dxf-go.svg?branch=master)](https://travis-ci.org/rpaloschi/dxf-go)
# DXF-go

A DXF implementation for Golang.

It was born from my personal need of a DXF parser for Go and the fact that none exists at the moment.

It was heavily influenced by the great work done at [dxfgrabber](https://github.com/mozman/dxfgrabber).

It currently parses a great part of a DXF file - 2014 compatible. 

It currently doesn't import the Object Section and doesn't generate the files. But it wil :)

There is a lot to be done and help is appreciated.

## Getting Started

A sample usage:

```
        file, err := os.Open(dxfPath) 
	if err != nil {
		log.Fatal(err)
	}

	doc, err := document.DxfDocumentFromStream(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, block := range doc.Blocks {
		for _, entity := range block.Entities {
			if polyline, ok := entity.(*entities.Polyline); ok {
				// process polyline here...
			} else if lwpolyline, ok := entity.(*entities.LWPolyline); ok {
				// process lwpolyline here...
			}
      //...
		}
	}
```

### Prerequisites

 * Go (1.8+)

### Installing

```
$ go get github.com/rpaloschi/dxf-go
``` 

## Authors

* **Ronald Paloschi** - [rpaloschi](https://github.com/rpaloschi)

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE.md) file for details

## Acknowledgments

TODO
