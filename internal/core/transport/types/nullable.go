package core_http_types

import (
	"encoding/json"
	"gopet/internal/core/domain"
)

type Nulladble[T any] struct {
	domain.Nulladble[T]
}

func (n *Nulladble[T]) UnmarshalJSON(b []byte) error {
	n.Set = true

	if string(b) == "null" {
		n.Value = nil

		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.Value = &value

	return nil
}

func (n *Nulladble[T]) ToDomain() domain.Nulladble[T] {
	return domain.Nulladble[T]{
		Value: n.Value,
		Set:   n.Set,
	}
}
