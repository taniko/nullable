package nullable

import (
	"bytes"
	"encoding/json"
)

type Nullable[T any] struct {
	value T
	valid bool
}

var (
	_ json.Marshaler   = (*Nullable[any])(nil)
	_ json.Unmarshaler = (*Nullable[any])(nil)
)

var nullBytes = []byte("null")

func New[T any](v T) Nullable[T] {
	return Nullable[T]{
		value: v,
		valid: true,
	}
}

func (t Nullable[T]) IsNull() bool {
	return !t.valid
}

func (t Nullable[T]) Value() T {
	return t.value
}

func (t Nullable[T]) MarshalJSON() ([]byte, error) {
	if t.IsNull() {
		return json.Marshal(nil)
	}
	return json.Marshal(t.Value())
}

func (t *Nullable[T]) UnmarshalJSON(b []byte) error {
	if b == nil {
		return nil
	} else if bytes.Equal(b, nullBytes) {
		return nil
	}
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	t.value = v
	t.valid = true
	return nil
}
