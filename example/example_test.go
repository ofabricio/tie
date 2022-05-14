package example_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/ofabricio/tie/json"
	"github.com/ofabricio/tie/u"
)

func ExampleND() {

	w, _ := ReqRes("")

	ch := make(chan Data)
	go func() {
		ch <- Data{Name: "Mary"}
		ch <- Data{Name: "John"}
		close(ch)
	}()

	u.Write(w, http.StatusCreated, json.ND(ch))

	fmt.Println(w.Code)
	fmt.Println(w.Body.String())

	// Output:
	// 201
	// {"name":"Mary"}
	// {"name":"John"}
}

func ExampleBindND() {

	_, r := ReqRes("{\"name\":\"Mary\"}\n{\"name\":\"John\"}")

	ch := make(chan Data)

	go u.Read(r, json.BindND(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Output:
	// {Mary}
	// {John}
}

func ExampleBindND_and_write() {

	w, r := ReqRes("{\"name\":\"Mary\"}\n{\"name\":\"John\"}")

	in := make(chan Data)
	ou := make(chan Data)

	go u.Read(r, json.BindND(in))

	go func() {
		for v := range in {
			v.Name = strings.ToUpper(v.Name)
			ou <- v
		}
		close(ou)
	}()

	u.Write(w, http.StatusCreated, json.ND(ou))

	fmt.Println(w.Code)
	fmt.Println(w.Body.String())

	// Output:
	// 201
	// {"name":"MARY"}
	// {"name":"JOHN"}
}

func ReqRes(body string) (*httptest.ResponseRecorder, *http.Request) {
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(body))
	return httptest.NewRecorder(), req
}

type Data struct {
	Name string `json:"name"`
}
