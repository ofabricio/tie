package u

import (
	"io"
	"net/http"
	"strings"

	"github.com/ofabricio/tie/k"
	"github.com/ofabricio/tie/opt"
)

func Read(r *http.Request, opts ...opt.ReadFunc) error {
	for _, f := range opts {
		if err := f(r); err != nil {
			return err
		}
	}
	return nil
}

func Write(w http.ResponseWriter, code int, opts ...opt.WriteFunc) error {
	h := w.Header()
	WriteBody := nopWriter
	for _, f := range opts {
		if body := f(h); body != nil {
			WriteBody = body
		}
	}
	w.WriteHeader(code)
	return WriteBody(w)
}

func Query(r *http.Request) *k.Param {
	return &k.Param{Get: r.URL.Query().Get}
}

func Head(r *http.Request) *k.Param {
	return &k.Param{Get: r.Header.Get}
}

func PathSeg(r *http.Request, seg int) k.Val {
	p := strings.Trim(r.URL.Path, "/")
	if s := strings.Split(p, "/"); seg < len(s) {
		return k.Val(s[seg])
	}
	return k.Val("")
}

func nopWriter(io.Writer) error {
	return nil
}
