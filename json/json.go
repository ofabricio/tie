package json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ofabricio/tie"
)

func Msg(format string, a ...any) tie.WriteFunc {
	return With("message", fmt.Sprintf(format, a...))
}

func With(kv ...any) tie.WriteFunc {
	m := make(map[string]any, len(kv)>>1)
	for i := 0; i < len(kv); i += 2 {
		k, v := kv[i+0].(string), kv[i+1]
		m[k] = v
	}
	return Body(m)
}

func Body(v any) tie.WriteFunc {
	return func(c *tie.WriteConfig) {
		c.Head.Set("Content-Type", "application/json; charset=utf-8")
		c.Body = func(w io.Writer) error {
			return json.NewEncoder(w).Encode(v)
		}
	}
}

func ND[T any](ch <-chan T) tie.WriteFunc {
	return func(c *tie.WriteConfig) {
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
