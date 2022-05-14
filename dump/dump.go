package dump

import (
	"net/http"
	"net/http/httputil"
)

func Res(r *http.Response) string {
	b, err := httputil.DumpResponse(r, true)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
