package head

import (
	"net/http"

	"github.com/ofabricio/tie/opt"
)

func With(kv ...string) opt.WriteFunc {
	return func(h http.Header) opt.BodyFunc {
		for i := 0; i < len(kv); i += 2 {
			k, v := kv[i+0], kv[i+1]
			h.Set(k, v)
		}
		return nil
	}
}
