package heidi

import (
	"errors"
	"fmt"
)

var (
	ErrOptionNotFound  = errors.New("not found")
	ErrOptionWrongType = errors.New("has wrong type")
)

type Options map[string]interface{}

func (o Options) Raw(name string) (interface{}, error) {
	v, ok := o[name]
	if !ok {
		return nil, fmt.Errorf(`option "%s" %w`, name, ErrOptionNotFound)
	}

	return v, nil
}

func (o Options) String(name string) (string, error) {
	raw, err := o.Raw(name)
	if err != nil {
		return "", err
	}

	v, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf(`option "%s" %w`, name, ErrOptionWrongType)
	}

	return v, nil
}

func (o Options) Int(name string) (int, error) {
	raw, err := o.Raw(name)
	if err != nil {
		return 0, err
	}

	v, ok := raw.(int)
	if !ok {
		return 0, fmt.Errorf(`option "%s" %w`, name, ErrOptionWrongType)
	}

	return v, nil
}
