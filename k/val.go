package k

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ErrRequiredKey  = fmt.Errorf("required")
	ErrInvalidValue = fmt.Errorf("not a valid value")
	ErrInvalidTime  = fmt.Errorf("not a valid time layout")
)

type Val string

func (v Val) Str() string {
	return string(v)
}

func (v Val) Split() []string {
	return strings.Split(string(v), ",")
}

func (v Val) Int() (int, error) {
	if v == "" {
		return 0, nil
	}
	i, err := strconv.Atoi(string(v))
	if err != nil {
		return i, ErrInvalidValue
	}
	return i, nil
}

func (v Val) Time(layout string) (time.Time, error) {
	if v == "" {
		return time.Time{}, nil
	}
	t, err := time.Parse(layout, string(v))
	if err != nil {
		return t, ErrInvalidTime
	}
	return t, nil
}

func (v Val) required() error {
	if v == "" {
		return ErrRequiredKey
	}
	return nil
}
