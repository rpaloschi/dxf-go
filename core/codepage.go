package core

import (
	"strings"
)

var codePages = map[string]string{
	"874":  "cp874",  // Thai,
	"932":  "cp932",  // Japanese
	"936":  "gbk",    // UnifiedChinese
	"949":  "cp949",  // Korean
	"950":  "cp950",  // TradChinese
	"1250": "cp1250", // CentralEurope
	"1251": "cp1251", // Cyrillic
	"1252": "cp1252", // WesternEurope
	"1253": "cp1253", // Greek
	"1254": "cp1254", // Turkish
	"1255": "cp1255", // Hebrew
	"1256": "cp1256", // Arabic
	"1257": "cp1257", // Baltic
	"1258": "cp1258", // Vietnam
}

func toEncoding(dxfCodePage string) string {
	for codePage, encoding := range codePages {
		if strings.HasSuffix(dxfCodePage, codePage) {
			return encoding
		}
	}
	return "cp1252"
}
