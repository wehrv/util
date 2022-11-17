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
	fields := strings.Fields(strings.TrimSpace(strings.Replace(str, "_", " ", -1)))
	for field := range fields {
		fields[field] = cases.Title(language.AmericanEnglish).String(fields[field])
	}
	header := strings.Join(fields, "")
	return header
}
