package heidi_test

import (
	"errors"
	"testing"

	"github.com/alexeyco/heidi"
)

func TestSuppress(t *testing.T) {
	t.Parallel()

	t.Run("SuccessOnString", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Suppress(heidi.Dataset{}, 0, "foo", "bar", heidi.Options{})
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "*" {
			t.Fatalf(`Suppressed value should be "*", "%s" given`, v)
		}
	})

	t.Run("SuccessOnInt", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Suppress(heidi.Dataset{}, 0, "foo", 123, heidi.Options{})
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "*" {
			t.Fatalf(`Suppressed value should be "*", "%s" given`, v)
		}
	})

	t.Run("SuccessWithCustomSymbol", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Suppress(heidi.Dataset{}, 0, "foo", "bar", heidi.Options{
			"symbol": "-",
		})

		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "-" {
			t.Fatalf(`Suppressed value should be "-", "%s" given`, v)
		}
	})

	t.Run("ErrorCauseWrongSymbolOptionType", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Suppress(heidi.Dataset{}, 0, "foo", "bar", heidi.Options{
			"symbol": 1,
		})

		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrOptionWrongType) {
			t.Fatal(`Error should be heidi.ErrOptionWrongType`)
		}

		if v != nil {
			t.Fatalf(`Suppressed value should be nil, "%v" given`, v)
		}
	})
}

func TestCensor(t *testing.T) {
	t.Parallel()

	t.Run("SuccessOnString", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Censor(heidi.Dataset{}, 0, "foo", "bar", heidi.Options{})

		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "***" {
			t.Fatalf(`Censored value should be "***", "%s" given`, v)
		}
	})

	t.Run("SuccessOnInt", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Censor(heidi.Dataset{}, 0, "foo", 1234, heidi.Options{})

		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "****" {
			t.Fatalf(`Censored value should be "****", "%s" given`, v)
		}
	})

	t.Run("SuccessOnInt64", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Censor(heidi.Dataset{}, 0, "foo", int64(123456789), heidi.Options{})

		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "*********" {
			t.Fatalf(`Censored value should be "*********", "%s" given`, v)
		}
	})

	t.Run("SuccessWithCustomSymbol", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Censor(heidi.Dataset{}, 0, "foo", "bar", heidi.Options{
			"symbol": "-",
		})

		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "---" {
			t.Fatalf(`Censored value should be "---", "%s" given`, v)
		}
	})

	t.Run("ErrorCauseWrongSymbolOptionType", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Censor(heidi.Dataset{}, 0, "foo", "bar", heidi.Options{
			"symbol": 1,
		})

		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrOptionWrongType) {
			t.Fatal(`Error should be heidi.ErrOptionWrongType`)
		}

		if v != nil {
			t.Fatalf(`Censored value should be nil, "%v" given`, v)
		}
	})

	t.Run("ErrorCauseWrongFilterType", func(t *testing.T) {
		t.Parallel()

		v, err := heidi.Censor(heidi.Dataset{}, 0, "foo", []string{}, heidi.Options{})

		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrFilterType) {
			t.Fatalf(`Error should be heidi.ErrFilterType`)
		}

		if v != nil {
			t.Fatalf(`Censored value should be nil, "%v" given`, v)
		}
	})
}
