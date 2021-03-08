package heidi

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

var ErrFilterType = errors.New("there are only following types supported")

type Filter func(s Dataset, n int, f string, v interface{}, options Options) (interface{}, error)

const suppressDefaultSymbol = "*"

func Suppress(_ Dataset, _ int, _ string, v interface{}, options Options) (interface{}, error) {
	symbol, err := options.String("symbol")
	if err != nil {
		if !errors.Is(err, ErrOptionNotFound) {
			return nil, err
		}

		symbol = suppressDefaultSymbol
	}

	return symbol, nil
}

func Censor(_ Dataset, _ int, _ string, v interface{}, options Options) (interface{}, error) {
	symbol, err := options.String("symbol")
	if err != nil {
		if !errors.Is(err, ErrOptionNotFound) {
			return nil, err
		}

		symbol = suppressDefaultSymbol
	}

	var s string
	switch v.(type) {
	case string:
		s = v.(string)
		break
	case int:
		s = fmt.Sprintf("%d", v.(int))
		break
	case int64:
		s = fmt.Sprintf("%d", v.(int64))
		break
	default:
		return nil, fmt.Errorf(`%w: string, int, int64`, ErrFilterType)
	}

	return strings.Repeat(symbol, utf8.RuneCountInString(s)), nil
}
