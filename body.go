package util

import (
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
}

func (body Body) New(r *http.Request) *Body {
	body.Path = strings.Split(r.URL.Path[1:], "/")
	body.Path[0] = isEql(body.Path[0], "", "app.html", body.Path[0])
	body.Path[0] = isEql(body.Path[0], "favicon.ico", "app.png", body.Path[0])
	body.Path[0] = "html/" + body.Path[0]
	dots := strings.Split(body.Path[len(body.Path)-1], ".")
	body.Mime = dots[len(dots)-1]
	if body.Mime == "" {
		body.Mime = "text"
	}
	body.Mime = mime.TypeByExtension("." + body.Mime)
	body.Maps = make(map[string]string)
	body.Case = body.Path[0]
	body.Body, body.Error = io.ReadAll(r.Body)
	return &body
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

func (body *Body) Send(w http.ResponseWriter, r *http.Request) *Body {
	if body.Error == nil {
		w.Header().Set("Content-Type", body.Mime)
		_, body.Error = w.Write(body.Body)
	}
	return body
}

func (body *Body) End(w http.ResponseWriter, r *http.Request) *Body {
	if body.Error == nil {
		body.File(strings.Join(body.Path, "/")).Send(w, r)
	}
	body.Err()
	return body
}
