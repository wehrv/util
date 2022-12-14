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
	Body   []byte
	Error  error
	Mime   string
	Path   []string
	Maps   map[string]string
	Reader *http.Request
	Writer http.ResponseWriter
}

func (body Body) New(w http.ResponseWriter, r *http.Request) *Body {
	body.Reader = r
	body.Writer = w
	body.Path = strings.Split(r.URL.Path[1:], "/")
	body.Maps = make(map[string]string)
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
	if body.Error == nil {
		dots := strings.Split(body.Path[len(body.Path)-1], ".")
		switch dots[len(dots)-1] {
		case "":
			body.Mime = "text/plain; charset=utf-8"
		case "application/json":
			body.Mime = "application/json; charset=utf-8"
		case "webmanifest":
			body.Mime = "application/manifest+json; charset=utf-8"
		default:
			body.Mime = mime.TypeByExtension("." + dots[len(dots)-1])
		}
		body.Writer.Header().Set("Content-Type", body.Mime)
		_, body.Error = body.Writer.Write(body.Body)
	}
	return body
}

func (body *Body) End() *Body {
	body.File(strings.Join(body.Path, "/")).Send().Err()
	return body
}
