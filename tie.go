package tie

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func New(w http.ResponseWriter, r *http.Request) Util {
	return Util{w: w, r: r}
}

func (u *Util) Bind(v interface{}) error {
	return json.NewDecoder(u.r.Body).Decode(v)
}

func (u *Util) Query(name string) string {
	return u.r.URL.Query().Get(name)
}

// PathSeg returns the URL path segment by its zero-based id.
func (u *Util) PathSeg(seg int) string {
	p := strings.Trim(u.r.URL.Path, "/")
	if s := strings.Split(p, "/"); seg < len(s) {
		return s[seg]
	}
	return ""
}

func (u *Util) Write(code int, opt ...WriteFunc) error {
	c := WriteConfig{Head: u.w.Header(), Body: nopBody}
	for _, f := range opt {
		f(&c)
	}
	u.w.WriteHeader(code)
	return c.Body(u.w)
}

type Util struct {
	w http.ResponseWriter
	r *http.Request
}

func nopBody(io.Writer) error {
	return nil
}

type WriteFunc func(*WriteConfig)

type WriteConfig struct {
	Head http.Header
	Body func(io.Writer) error
}
