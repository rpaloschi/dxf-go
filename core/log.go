package core

import (
	"log"
	"os"
)

var Log *log.Logger = log.New(os.Stderr, "dxf-go", 0)
