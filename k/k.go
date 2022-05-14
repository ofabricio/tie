package k

import (
	"fmt"
	"time"
)

type Param struct {
	Get func(string) string
	Err *Error
}

func (q *Param) Val(name string) Val {
	return Val(q.Str(name))
}

func (q *Param) Str(name string) string {
	return q.Get(name)
}

func (q *Param) Int(name string) int {
	v, err := q.Val(name).Int()
	q.setErr(name, err)
	return v
}

func (q *Param) Time(name, layout string) time.Time {
	v, err := q.Val(name).Time(layout)
	q.setErr(name, err)
	return v
}

func (q *Param) Split(name string) *split {
	return &split{q.Val(name).Split(), q}
}

func (q *Param) Required(name string) *required {
	v := q.Val(name)
	q.setErr(name, v.required())
	return &required{name, string(v), q}
}

func (q *Param) setErr(name string, err error) {
	if err != nil && q.Err == nil {
		q.Err = &Error{Key: name, Err: err} //fmt.Errorf("key %s %v", name, err)
	}
}

type split struct {
	v []string
	*Param
}

func (q *split) Values() []string {
	return q.v
}

func (q *split) Int(name string) (v []int) {
	for _, s := range q.v {
		i, err := Val(s).Int()
		q.setErr(name, err)
		v = append(v, i)
	}
	return
}

func (q *split) Time(name, layout string) (v []time.Time) {
	for _, s := range q.v {
		t, err := Val(s).Time(layout)
		q.setErr(name, err)
		v = append(v, t)
	}
	return
}

type required struct {
	k string
	v string
	*Param
}

func (q *required) Val() Val {
	return Val(q.v)
}

func (q *required) Str() string {
	return q.v
}

func (q *required) Int() int {
	i, err := q.Val().Int()
	q.setErr(q.k, err)
	return i
}

func (q *required) Time(layout string) time.Time {
	t, err := q.Val().Time(layout)
	q.setErr(q.k, err)
	return t
}

type Error struct {
	Key string
	Val string
	Err error
}

func (e *Error) Error() string {
	return fmt.Sprintf("key %s is %v", e.Key, e.Err)
}
