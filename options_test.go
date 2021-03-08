package heidi_test

import (
	"errors"
	"testing"

	"github.com/alexeyco/heidi"
)

func TestOptions_Raw(t *testing.T) {
	t.Parallel()

	options := heidi.Options{
		"foo":  "bar",
		"fizz": 123,
	}

	t.Run("SuccessCauseOptionExist", func(t *testing.T) {
		t.Parallel()

		v, err := options.Raw("foo")
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "bar" {
			t.Fatalf(`Option should be "bar", "%v" given`, v)
		}

		v, err = options.Raw("fizz")
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != 123 {
			t.Fatalf(`Option should be "123", "%v" given`, v)
		}
	})

	t.Run("ErrorCauseOptionDoesNotExist", func(t *testing.T) {
		t.Parallel()

		v, err := options.Raw("buzz")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrOptionNotFound) {
			t.Fatal(`Error should be heidi.ErrOptionNotFound`)
		}

		if v != nil {
			t.Fatalf(`Option should be nil, "%v" given`, v)
		}
	})
}

func TestOptions_String(t *testing.T) {
	t.Parallel()

	options := heidi.Options{
		"foo":  "bar",
		"fizz": 123,
	}

	t.Run("SuccessCauseOptionIsString", func(t *testing.T) {
		t.Parallel()

		v, err := options.String("foo")
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "bar" {
			t.Fatalf(`Option should be "bar", "%s" given`, v)
		}
	})

	t.Run("ErrorCauseOptionHasWrongType", func(t *testing.T) {
		t.Parallel()

		v, err := options.String("fizz")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrOptionWrongType) {
			t.Fatal(`Error should be heidi.ErrOptionWrongType`)
		}

		if v != "" {
			t.Fatalf(`Option should be empty string, "%s" given`, v)
		}
	})

	t.Run("ErrorCauseOptionDoesNotExist", func(t *testing.T) {
		t.Parallel()

		v, err := options.String("buzz")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrOptionNotFound) {
			t.Fatal(`Error should be heidi.ErrOptionNotFound`)
		}

		if v != "" {
			t.Fatalf(`Option should be empty string, "%s" given`, v)
		}
	})
}

func TestOptions_Int(t *testing.T) {
	t.Parallel()

	options := heidi.Options{
		"foo":  "bar",
		"fizz": 123,
	}

	t.Run("SuccessCauseOptionIsInt", func(t *testing.T) {
		t.Parallel()

		v, err := options.Int("fizz")
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != 123 {
			t.Fatalf(`Option should be 123, %d given`, v)
		}
	})

	t.Run("ErrorCauseOptionHasWrongType", func(t *testing.T) {
		t.Parallel()

		v, err := options.Int("foo")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrOptionWrongType) {
			t.Fatal(`Error should be heidi.ErrOptionWrongType`)
		}

		if v != 0 {
			t.Fatalf(`Option should be 0, %d given`, v)
		}
	})

	t.Run("ErrorCauseOptionDoesNotExist", func(t *testing.T) {
		t.Parallel()

		v, err := options.Int("buzz")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrOptionNotFound) {
			t.Fatal(`Error should be heidi.ErrOptionNotFound`)
		}

		if v != 0 {
			t.Fatalf(`Option should be 0, %d given`, v)
		}
	})
}
