package tie

import (
	"encoding/json"
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) util {
	return util{w, r}
}

type util struct {
	w http.ResponseWriter
	r *http.Request
}

func (t *util) Json(code int, v interface{}) error {
	t.w.WriteHeader(code)
	t.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(t.w).Encode(v)
}

func (t *util) Bind(v interface{}) error {
	return json.NewDecoder(t.r.Body).Decode(v)
}

func (t *util) QueryParam(s string) string {
	return t.r.URL.Query().Get(s)
}
