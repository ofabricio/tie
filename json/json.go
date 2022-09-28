package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ofabricio/tie/opt"
)

func Msg(format string, a ...any) opt.WriteFunc {
	return With("message", fmt.Sprintf(format, a...))
}

func With(kv ...any) opt.WriteFunc {
	m := make(map[string]any, len(kv)>>1)
	for i := 0; i < len(kv); i += 2 {
		k, v := kv[i+0].(string), kv[i+1]
		m[k] = v
	}
	return Body(m)
}

func Body(v any) opt.WriteFunc {
	return func(h http.Header) opt.BodyFunc {
		h.Set("Content-Type", "application/json")
		return func(w io.Writer) error {
			return json.NewEncoder(w).Encode(v)
		}
	}
}

func ND[T any](ch <-chan T) opt.WriteFunc {
	return func(h http.Header) opt.BodyFunc {
		h.Set("Content-Type", "application/x-ndjson")
		return func(w io.Writer) error {
			enc := json.NewEncoder(w)
			for v := range ch {
				if err := enc.Encode(v); err != nil {
					return err
				}
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
			}
			return nil
		}
	}
}

func Bind(v any) opt.ReadFunc {
	return func(r *http.Request) error {
		return json.NewDecoder(r.Body).Decode(v)
	}
}

func BindND[T any](ch chan<- T) opt.ReadFunc {
	return func(r *http.Request) error {
		defer close(ch)
		enc := json.NewDecoder(r.Body)
		for {
			var v T
			if err := enc.Decode(&v); err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			ch <- v
		}
		return nil
	}
}
