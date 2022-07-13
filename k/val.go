package k

import (
	"strconv"
	"strings"
	"time"
)

type Val string

func (v Val) Str() string {
	return string(v)
}

func (v Val) Split(sep string) []string {
	if v.Empty() {
		return nil
	}
	return strings.Split(string(v), sep)
}

func (v Val) Int() (int, error) {
	if v.Empty() {
		return 0, nil
	}
	return strconv.Atoi(string(v))
}

func (v Val) IntClamp(min, max int) (int, error) {
	i, err := v.Int()
	if i < min {
		return min, err
	}
	if i > max {
		return max, err
	}
	return i, err
}

func (v Val) Time(layout string) (time.Time, error) {
	if v.Empty() {
		return time.Time{}, nil
	}
	return time.Parse(layout, string(v))
}

func (v Val) Default(def string) Val {
	if v.Empty() {
		return Val(def)
	}
	return v
}

func (v Val) Empty() bool {
	return v == ""
}
