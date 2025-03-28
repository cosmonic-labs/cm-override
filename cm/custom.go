package cm

import (
	"bytes"
	"encoding/json"
	_ "unsafe"
)

var nullJSON = []byte("null")

// Custom Types
type Option[T any] struct {
	_ HostLayout
	option[T]
}

// None returns an [Option] representing the none case,
// equivalent to the zero value.
func None[T any]() Option[T] {
	return Option[T]{}
}

// Some returns an [Option] representing the some case.
func Some[T any](v T) Option[T] {
	return Option[T]{
		option: option[T]{
			isSome: true,
			some:   v,
		},
	}
}

// option represents the internal representation of a Component Model option type.
// The first byte is a bool representing none or some,
// followed by storage for the associated type T.
type option[T any] struct {
	_      HostLayout
	isSome bool
	some   T
}

// None returns true if o represents the none case.
func (o *option[T]) None() bool {
	return !o.isSome
}

// Some returns a non-nil *T if o represents the some case,
// or nil if o represents the none case.
func (o *option[T]) Some() *T {
	if o.isSome {
		return &o.some
	}
	return nil
}

// Value returns T if o represents the some case,
// or the zero value of T if o represents the none case.
// This does not have a pointer receiver, so it can be chained.
func (o option[T]) Value() T {
	if !o.isSome {
		var zero T
		return zero
	}
	return o.some
}

func (o *option[T]) MarshalJSON() ([]byte, error) {
	if !o.isSome {
		return []byte("null"), nil
	}

	return json.Marshal(o.some)
}

func (o *option[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullJSON) {
		*o = option[T]{}
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*o = option[T]{isSome: true, some: v}
	return nil
}
