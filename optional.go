package optional

import (
	"encoding/json"
)

type Optional[T any] struct {
	present bool
	value   T
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) IsNull() bool {
	return !o.present
}

func (o Optional[T]) Present() bool {
	return o.present
}

func (o Optional[T]) Value() T {
	return o.value
}

func (o Optional[T]) ValueOr(d T) T {
	if o.present {
		return o.value
	}

	return d
}

func (o Optional[T]) OrElse(fn func() T) T {
	if o.present {
		return o.value
	}
	return fn()
}

func (o Optional[T]) AndThen(fn func(v T) Optional[T]) Optional[T] {
	if o.present {
		return fn(o.value)
	}
	return Empty[T]()
}

func (o *Optional[T]) Reset() {
	o.present = false
}

func (o *Optional[T]) MarshalJSON() ([]byte, error) {
	if o.Present() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if err := json.Unmarshal(data, &o.value); err != nil {
		return err
	}

	o.present = true
	return nil
}
