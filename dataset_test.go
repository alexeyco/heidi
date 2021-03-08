package heidi_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexeyco/heidi"
)

func TestDataset_Pluck(t *testing.T) {
	t.Parallel()

	dataset := heidi.Dataset{
		heidi.Row{
			"string": "foo",
			"int":    123,
			"int64":  int64(123),
		},
		heidi.Row{
			"string": "bar",
			"int":    456,
			"int64":  int64(456),
		},
	}

	t.Run("SuccessOnString", func(t *testing.T) {
		t.Parallel()

		expected := []interface{}{
			"foo",
			"bar",
		}

		v, err := dataset.Pluck("string")
		if err != nil {
			t.Fatalf(`Error should be nil, %s given`, err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Fatalf(`Plucked dataset (%v) should be equal to expected slice (%v)`, v, expected)
		}
	})

	t.Run("SuccessOnInt", func(t *testing.T) {
		t.Parallel()

		expected := []interface{}{
			123,
			456,
		}

		v, err := dataset.Pluck("int")
		if err != nil {
			t.Fatalf(`Error should be nil, %s given`, err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Fatalf(`Plucked dataset (%v) should be equal to expected slice (%v)`, v, expected)
		}
	})

	t.Run("SuccessOnInt64", func(t *testing.T) {
		t.Parallel()

		expected := []interface{}{
			int64(123),
			int64(456),
		}

		v, err := dataset.Pluck("int64")
		if err != nil {
			t.Fatalf(`Error should be nil, %s given`, err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Fatalf(`Plucked dataset (%v) should be equal to expected slice (%v)`, v, expected)
		}
	})

	t.Run("ErrorCauseFieldValueDoesNotExist", func(t *testing.T) {
		t.Parallel()

		v, err := dataset.Pluck("wrong")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrValueNotFound) {
			t.Fatal(`Error should be heidi.ErrValueNotFound`)
		}

		if v != nil {
			t.Fatalf(`Plucked value should be nil, "%v" given`, v)
		}
	})
}

func TestRow_Raw(t *testing.T) {
	t.Parallel()

	row := heidi.Row{
		"foo":  "bar",
		"fizz": 123,
	}

	t.Run("SuccessCauseOptionExist", func(t *testing.T) {
		t.Parallel()

		v, err := row.Raw("foo")
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "bar" {
			t.Fatalf(`Row value should be "bar", "%v" given`, v)
		}

		v, err = row.Raw("fizz")
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != 123 {
			t.Fatalf(`Row value should be "123", "%v" given`, v)
		}
	})

	t.Run("ErrorCauseOptionDoesNotExist", func(t *testing.T) {
		t.Parallel()

		v, err := row.Raw("buzz")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrValueNotFound) {
			t.Fatal(`Error should be heidi.ErrValueNotFound`)
		}

		if v != nil {
			t.Fatalf(`Row value should be nil, "%v" given`, v)
		}
	})
}

func TestRow_String(t *testing.T) {
	t.Parallel()

	row := heidi.Row{
		"foo":  "bar",
		"fizz": 123,
	}

	t.Run("SuccessCauseValueIsString", func(t *testing.T) {
		t.Parallel()

		v, err := row.String("foo")
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != "bar" {
			t.Fatalf(`Row value should be "bar", "%s" given`, v)
		}
	})

	t.Run("ErrorCauseValueHasWrongType", func(t *testing.T) {
		t.Parallel()

		v, err := row.String("fizz")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrValueWrongType) {
			t.Fatal(`Error should be heidi.ErrValueWrongType`)
		}

		if v != "" {
			t.Fatalf(`Row value should be empty string, "%s" given`, v)
		}
	})

	t.Run("ErrorCauseValueDoesNotExist", func(t *testing.T) {
		t.Parallel()

		v, err := row.String("buzz")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrValueNotFound) {
			t.Fatal(`Error should be heidi.ErrValueNotFound`)
		}

		if v != "" {
			t.Fatalf(`Row value should be empty string, "%s" given`, v)
		}
	})
}

func TestRow_Int(t *testing.T) {
	t.Parallel()

	row := heidi.Row{
		"foo":  "bar",
		"fizz": 123,
	}

	t.Run("SuccessCauseValueIsInt", func(t *testing.T) {
		t.Parallel()

		v, err := row.Int("fizz")
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if v != 123 {
			t.Fatalf(`Row value should be 123, %d given`, v)
		}
	})

	t.Run("ErrorCauseValueHasWrongType", func(t *testing.T) {
		t.Parallel()

		v, err := row.Int("foo")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrValueWrongType) {
			t.Fatal(`Error should be heidi.ErrValueWrongType`)
		}

		if v != 0 {
			t.Fatalf(`Row value should be 0, %d given`, v)
		}
	})

	t.Run("ErrorCauseValueDoesNotExist", func(t *testing.T) {
		t.Parallel()

		v, err := row.Int("buzz")
		if err == nil {
			t.Fatal(`Error should not be nil`)
		}

		if !errors.Is(err, heidi.ErrValueNotFound) {
			t.Fatal(`Error should be heidi.ErrValueNotFound`)
		}

		if v != 0 {
			t.Fatalf(`Row value should be 0, %d given`, v)
		}
	})
}
