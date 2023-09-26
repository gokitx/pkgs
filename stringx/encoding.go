package stringx

import (
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func DataToUTF8(s []byte) ([]byte, error) {
	// https://developer.mozilla.org/zh-CN/docs/Web/API/Encoding_API/Encodings
	encoder, name, _ := charset.DetermineEncoding(s, "")

	switch name {
	case "unicode-1-1-utf-8", "utf-8", "utf8":
		return s, nil
	case "gbk", "gb18030", "hz-gb2312":
		return simplifiedchinese.GBK.NewDecoder().Bytes(s)
	default:
		// TODO
		// windows-1252
		if strings.HasPrefix(name, "windows-") {
			return simplifiedchinese.GBK.NewDecoder().Bytes(s)
		}
		// default encoding decoder
		return encoder.NewDecoder().Bytes(s)
	}
}
