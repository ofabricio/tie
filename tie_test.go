package tie_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/ofabricio/tie"
)

func Example() {

	var payload struct {
		Name string `json:"name"`
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/users/123?type=ADM", strings.NewReader(`{ "name": "Mary" }`))
	r.Header.Set("X-Version", "1")

	u := tie.New(w, r)

	fmt.Println("Query:", u.Query("type"))
	fmt.Println("Header:", u.Head("X-Version"))
	fmt.Println("PathSeg:", u.PathSeg(0))
	u.Bind(&payload)

	fmt.Println("Body:", payload.Name)

	// Output:
	// Query: ADM
	// Header: 1
	// PathSeg: users
	// Body: Mary
}
