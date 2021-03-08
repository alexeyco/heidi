package heidi

import (
	"errors"
	"fmt"
)

var (
	ErrRowNotFound    = errors.New("not found")
	ErrValueNotFound  = errors.New("not found")
	ErrValueWrongType = errors.New("has wrong type")
)

type Dataset []Row

func (s Dataset) Pluck(name string) (v []interface{}, err error) {
	v = make([]interface{}, len(s))
	for n, row := range s {
		if v[n], err = row.Raw(name); err != nil {
			return nil, err
		}
	}

	return v, nil
}

type Row map[string]interface{}

func (r Row) Raw(name string) (interface{}, error) {
	v, ok := r[name]
	if !ok {
		return nil, fmt.Errorf(`row value "%s" %w`, name, ErrValueNotFound)
	}

	return v, nil
}

func (r Row) String(name string) (string, error) {
	raw, err := r.Raw(name)
	if err != nil {
		return "", err
	}

	v, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf(`row value "%s" %w`, name, ErrValueWrongType)
	}

	return v, nil
}

func (r Row) Int(name string) (int, error) {
	raw, err := r.Raw(name)
	if err != nil {
		return 0, err
	}

	v, ok := raw.(int)
	if !ok {
		return 0, fmt.Errorf(`row value "%s" %w`, name, ErrValueWrongType)
	}

	return v, nil
}
