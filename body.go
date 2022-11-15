package util

import (
	"encoding/json"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"
)

type Body struct {
	Body  []byte
	Error error
	Case  string
	Mime  string
	Path  []string
	Maps  map[string]string
	R     *http.Request
	W     http.ResponseWriter
}

func (body Body) New(w http.ResponseWriter, r *http.Request) *Body {
	body.R = r
	body.W = w
	body.Path = strings.Split(r.URL.Path[1:], "/")
	body.Maps = make(map[string]string)
	body.Case = body.Path[0]
	body.Body, body.Error = io.ReadAll(r.Body)
	body.Unmarshal()
	return &body
}

func (body *Body) Marshal() *Body {
	if body.Error == nil {
		body.Body, body.Error = json.Marshal(body.Maps)
		if body.Error == nil {
			body.Mime = "application/json; charset=utf-8"
		}
	}
	return body
}

func (body *Body) Unmarshal() *Body {
	var keys []string
	var maps [][]string
	body.Error = json.Unmarshal(body.Body, &keys)
	if body.Error == nil {
		for _, key := range keys {
			body.Maps[key] = ""
		}
	} else {
		body.Error = nil
		json.Unmarshal(body.Body, &maps)
	}
	return body
}

func (body *Body) Err() {
	if body.Error != nil {
		log.Println(body.Error)
	}
}

func (body *Body) File(file string) *Body {
	if body.Error == nil {
		body.Body, body.Error = os.ReadFile(strings.Join(body.Path, "/"))
	}
	return body
}

func (body *Body) Send() *Body {
	dots := strings.Split(body.Path[len(body.Path)-1], ".")
	body.Mime = dots[len(dots)-1]
	if body.Mime == "" {
		body.Mime = "text"
	}
	body.Mime = mime.TypeByExtension("." + body.Mime)
	if body.Error == nil {
		body.W.Header().Set("Content-Type", body.Mime)
		_, body.Error = body.W.Write(body.Body)
	}
	return body
}

func (body *Body) End() *Body {
	if body.Error == nil {
		body.File(strings.Join(body.Path, "/")).Send()
	}
	body.Err()
	return body
}
