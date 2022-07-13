package tie

import (
	"net/http"

	"github.com/ofabricio/tie/json"
	"github.com/ofabricio/tie/opt"
	"github.com/ofabricio/tie/u"
)

func New(w http.ResponseWriter, r *http.Request) Tied {
	return Tied{w: w, r: r}
}

func (t *Tied) Query(name string) string {
	return u.Query(t.r).Name(name).Str()
}

func (t *Tied) Head(name string) string {
	return u.Head(t.r).Name(name).Str()
}

// PathSeg returns the URL path segment by its zero-based id.
func (t *Tied) PathSeg(seg int) string {
	return u.PathSeg(t.r, seg).Str()
}

func (t *Tied) Bind(v any) error {
	return t.Read(json.Bind(v))
}

func (t *Tied) Read(opts ...opt.ReadFunc) error {
	return u.Read(t.r, opts...)
}

func (t *Tied) Write(code int, opts ...opt.WriteFunc) error {
	return u.Write(t.w, code, opts...)
}

type Tied struct {
	w http.ResponseWriter
	r *http.Request
}
