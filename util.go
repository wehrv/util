package util // github.com/wehrv/util

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func isErr(err error) bool {
	return err != nil
}

func noErr(err error) bool {
	return err == nil
}

func abErr(err error) {
	if isErr(err) {
		log.Fatal(err)
	}
}

func ppErr(err error) bool {
	if isErr(err) {
		log.Println(err)
		return true
	}
	return false
}

func npErr(err error) bool {
	return !ppErr(err)
}

func ipErr(err error) bool {
	return ppErr(err)
}

func IsEql(a, b, c, d string) string {
	if a == b {
		return c
	}
	return d
}

func fetch(url string) ([]byte, error) {
	val, err := http.Get(url)
	if ipErr(err) {
		return []byte{}, err
	}
	defer val.Body.Close()
	return io.ReadAll(val.Body)
}

func ungz(data []byte) ([]byte, error) {
	datb, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return data, err
	}
	return io.ReadAll(datb)
}

func snakeToCamel(str string) string {
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
