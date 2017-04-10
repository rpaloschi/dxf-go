package core

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const dxfWithInfo = `  0
SECTION
  2
HEADER
  9
$ACADVER
  1
AC1018
  9
$DWGCODEPAGE
  3
ANSI_1252
  9
$HANDSEED
  5
1
  0
ENDSEC
  0
EOF
`

func TestDXFInfo(t *testing.T) {
	info, err := GetDXFInfo(strings.NewReader(dxfWithInfo))

	assert.Nil(t, err)
	assert.Equal(t, "R2004", info.Release)
	assert.Equal(t, "AC1018", info.Version)
	assert.Equal(t, "cp1252", info.Encoding)
	assert.Equal(t, "1", info.Handseed)
}

func TestDXFInfoError(t *testing.T) {
	info, err := GetDXFInfo(strings.NewReader("  XXX\nSECTION"))

	assert.Equal(t, Info{}, info)
	assert.NotNil(t, err)
}

const dxfInfoAttributeError = `  0
SECTION
  2
HEADER
  9
$ACADVER
  XXX
AC1018
  0
ENDSEC
  0
EOF
`

func TestDXFInfoAttributeError(t *testing.T) {
	info, err := GetDXFInfo(strings.NewReader(dxfInfoAttributeError))

	assert.Equal(t, Info{}, info)
	assert.NotNil(t, err)
}

const dxfInfoDefaultACADVER = `  0
SECTION
  2
HEADER
  9
$ACADVER
  1
XXX
  0
ENDSEC
  0
EOF
`

func TestDXFInfoDefaultAcadVer(t *testing.T) {
	info, _ := GetDXFInfo(strings.NewReader(dxfInfoDefaultACADVER))

	assert.Equal(t, "R12", info.Release)
}
