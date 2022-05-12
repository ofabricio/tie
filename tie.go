package tie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func New(w http.ResponseWriter, r *http.Request) util {
	return util{w: w, r: r}
}

func (u *util) Bind(v interface{}) error {
	return json.NewDecoder(u.r.Body).Decode(v)
}

func (u *util) Query(name string) string {
	return u.r.URL.Query().Get(name)
}

// PathSeg returns the URL path segment by its zero-based id.
func (u *util) PathSeg(seg int) string {
	p := strings.Trim(u.r.URL.Path, "/")
	if s := strings.Split(p, "/"); seg < len(s) {
		return s[seg]
	}
	return ""
}

func (u *util) Write(code int, opt ...WriteFunc) error {
	c := WriteConfig{Head: u.w.Header(), Body: nopWriter}
	for _, f := range opt {
		f(&c)
	}
	u.w.WriteHeader(code)
	return c.Body(u.w)
}

type util struct {
	w http.ResponseWriter
	r *http.Request
}

func Msg(format string, a ...any) WriteFunc {
	return JsonKV("message", fmt.Sprintf(format, a...))
}

func JsonKV(kv ...any) WriteFunc {
	m := make(map[string]any, len(kv)>>1)
	for i := 0; i < len(kv); i += 2 {
		k, v := kv[i+0], kv[i+1]
		m[k.(string)] = v
	}
	return Json(m)
}

func Json(v any) WriteFunc {
	return func(c *WriteConfig) {
		c.Head.Set("Content-Type", "application/json; charset=utf-8")
		c.Body = func(w io.Writer) error {
			return json.NewEncoder(w).Encode(v)
		}
	}
}

func NDJson[T any](ch chan T) WriteFunc {
	return func(c *WriteConfig) {
		c.Head.Set("Content-Type", "application/x-ndjson; charset=utf-8")
		c.Body = func(w io.Writer) error {
			enc := json.NewEncoder(w)
			for v := range ch {
				if err := enc.Encode(v); err != nil {
					return err
				}
			}
			return nil
		}
	}
}

func Copy(r io.Reader) WriteFunc {
	return func(c *WriteConfig) {
		c.Body = func(w io.Writer) error {
			_, err := io.Copy(w, r)
			return err
		}
	}
}

func Header(kv ...string) WriteFunc {
	return func(c *WriteConfig) {
		for i := 0; i < len(kv); i += 2 {
			k, v := kv[i+0], kv[i+1]
			c.Head.Set(k, v)
		}
	}
}

func nopWriter(io.Writer) error {
	return nil
}

type WriteFunc func(*WriteConfig)

type WriteConfig struct {
	Head http.Header
	Body func(io.Writer) error
}
