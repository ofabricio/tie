package k

import (
	"errors"
	"fmt"
	"time"
)

type Param struct {
	Get func(string) string
	Err *Error
	key string
	val Val
}

func (p *Param) Name(name string) *Param {
	p.key = name
	p.val = Val(p.Get(name))
	return p
}

func (p *Param) Default(def string) *Param {
	p.val = p.val.Default(def)
	return p
}

func (p *Param) Required() *Param {
	if p.val == "" {
		p.setErr(ErrRequiredKey, ErrRequiredKey)
	}
	return p
}

func (p *Param) Lower() *Param {
	p.val = p.val.Lower()
	return p
}

func (p *Param) Upper() *Param {
	p.val = p.val.Upper()
	return p
}

func (p *Param) Split() []string {
	return p.val.Split(",")
}

func (p *Param) Str() string {
	return p.val.Str()
}

func (p *Param) Int() int {
	v, err := p.val.Int()
	p.setErr(ErrInvalidValue, err)
	return v
}

func (p *Param) IntClamp(min, max int) int {
	v, err := p.val.IntClamp(min, max)
	p.setErr(ErrInvalidValue, err)
	return v
}

func (p *Param) Time(layout string) time.Time {
	v, err := p.val.Time(layout)
	p.setErr(ErrInvalidTime, err)
	return v
}

func (p *Param) setErr(err, why error) {
	if why != nil && p.Err == nil {
		p.Err = &Error{Key: p.key, Val: p.val.Str(), Err: err, Why: why}
	}
}

type Error struct {
	Key string
	Val string
	Err error
	Why error
}

func (e *Error) Error() string {
	return fmt.Sprintf("key %s is %v", e.Key, e.Err)
}

var (
	ErrRequiredKey  = errors.New("required")
	ErrInvalidValue = errors.New("not a valid value")
	ErrInvalidTime  = errors.New("not a valid time layout")
)
