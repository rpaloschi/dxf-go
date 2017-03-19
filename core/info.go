package core

import "io"

// Info struct. Minimal Information from a DXF File.
type Info struct {
	Release  string
	Version  string
	Encoding string
	Handseed string
}

type setInfoDataFromTag func(Tag, *Info)

var infoTagMapper map[string]setInfoDataFromTag

// GetDXFInfo returns an Info record extracted from the DXF at stream.
func GetDXFInfo(stream io.Reader) (Info, error) {
	info := Info{}
	next := Tagger(stream)

	for {
		tag, err := next()
		if err != nil {
			return info, err
		}

		if *tag == NoneTag || tag.Value.ToString() == "ENDSEC" {
			break
		}

		if tag.Code != 9 {
			continue
		}

		name := tag.Value.ToString()[1:]
		if setter, ok := infoTagMapper[name]; ok {
			tag, err := next()
			if err != nil {
				return info, err
			}

			setter(*tag, &info)
		}
	}
	return info, nil
}

var acadRelease = map[string]string{
	"AC1009": "R12",
	"AC1012": "R13",
	"AC1014": "R14",
	"AC1015": "R2000",
	"AC1018": "R2004",
	"AC1021": "R2007",
	"AC1024": "R2010",
}

func init() {
	infoTagMapper = make(map[string]setInfoDataFromTag)

	infoTagMapper["DWGCODEPAGE"] = func(tag Tag, info *Info) {
		value, _ := AsString(tag.Value)
		info.Encoding = toEncoding(value)
	}

	infoTagMapper["ACADVER"] = func(tag Tag, info *Info) {
		info.Version, _ = AsString(tag.Value)
		release, ok := acadRelease[info.Version]
		if !ok {
			release = "R12"
		}
		info.Release = release
	}

	infoTagMapper["HANDSEED"] = func(tag Tag, info *Info) {
		info.Handseed, _ = AsString(tag.Value)
	}
}
