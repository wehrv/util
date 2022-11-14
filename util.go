package util // github.com/wehrv/util

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func IsEql(a, b, c, d string) string {
	if a == b {
		return c
	}
	return d
}

func Fetch(url string) ([]byte, error) {
	val, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer val.Body.Close()
	return io.ReadAll(val.Body)
}

func UnGZ(data []byte) ([]byte, error) {
	datb, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return data, err
	}
	return io.ReadAll(datb)
}

func SnakeToCamel(str string) string {
	//	fmt.Println(str)
	fields := strings.Fields(strings.TrimSpace(strings.Replace(str, "_", " ", -1)))
	//	fmt.Println(fields)
	for field := range fields {
		caser := cases.Title(language.AmericanEnglish)
		fields[field] = caser.String(fields[field])
	}
	//	fmt.Println(fields)
	header := strings.Join(fields, "")
	//	fmt.Println(header)
	return header
}
