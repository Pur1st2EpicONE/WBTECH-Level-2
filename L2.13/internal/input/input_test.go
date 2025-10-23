package input

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/spf13/pflag"
)

func TestParseArgsNoFields(t *testing.T) {

	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)

	os.Args = []string{"aboba"}
	data, err := ParseArgs()

	if err == nil {
		t.Fatalf("expected error, got nil, data=%+v", data)
	}

	if !errors.Is(err, ErrNoFieldsToShow) {
		t.Fatalf("expected ErrNoFieldsToShow, got %v", err)
	}

}

func TestParseArgsValid(t *testing.T) {

	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)

	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)

	os.Args = []string{"abuba", "-f", "1,3-5", "file1.txt", "file2.txt"}
	data, err := ParseArgs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedFields := []int{1, 3, 4, 5}
	if !reflect.DeepEqual(data.Fields, expectedFields) {
		t.Fatalf("fields mismatch: want %v, got %v", expectedFields, data.Fields)
	}

	if !reflect.DeepEqual(data.FileNames, []string{"file1.txt", "file2.txt"}) {
		t.Fatalf("filenames mismatch: want %v, got %v", []string{"file1.txt", "file2.txt"}, data.FileNames)
	}

	if len(data.Ranges) != 0 {
		t.Fatalf("expected no open ranges, got %v", data.Ranges)
	}

}

func TestParseArgsOpenRange(t *testing.T) {

	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)

	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)

	os.Args = []string{"qweqweqweqwe", "-f", "2-"}
	data, err := ParseArgs()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(data.Ranges) != 1 {
		t.Fatalf("expected 1 range, got %d", len(data.Ranges))
	}

	r := data.Ranges[0]
	if r.Start != 2 || r.End != -1 {
		t.Fatalf("unexpected range content: want start=2,end=-1 got %+v", r)
	}

}

func TestEdgeCases(t *testing.T) {

	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	t.Run("non-numeric", func(t *testing.T) {

		pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)

		data := new(Data)
		fieldsMap := make(map[int]struct{})
		err := saveField("qwe", data, fieldsMap)

		if !errors.Is(err, ErrInvalidFieldValue) {
			t.Fatalf("expected ErrInvalidFieldValue, got %v", err)
		}

		if data.Flags.Flagf != "qwe" {
			t.Fatalf("expected 'qwe', got %q", data.Flags.Flagf)
		}

	})

	t.Run("zero-field", func(t *testing.T) {

		pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
		data := new(Data)
		fieldsMap := make(map[int]struct{})

		err := saveField("0", data, fieldsMap)

		if !errors.Is(err, ErrNotNumberedFromOne) {
			t.Fatalf("expected ErrNotNumberedFromOne, got %v", err)
		}

	})

	t.Run("double-dash-data-fields", func(t *testing.T) {

		origArgs := os.Args
		defer func() { os.Args = origArgs }()

		os.Args = []string{"quehfahslfk", "-f", "1--3"}

		_, err := ParseArgs()

		if !errors.Is(err, ErrInvalidFieldRange) {
			t.Fatalf("expected ErrInvalidFieldRange for '1--3', got %v", err)
		}

	})

	t.Run("decreasing-range", func(t *testing.T) {

		pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
		data := new(Data)
		fieldsMap := make(map[int]struct{})

		err := saveRange("5-3", data, fieldsMap)

		if !errors.Is(err, ErrDecreasingRange) {
			t.Fatalf("expected ErrDecreasingRange, got %v", err)
		}

	})

	t.Run("getRangeStart-negative", func(t *testing.T) {

		pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
		data := new(Data)

		_, err := getRangeStart([]string{"-2"}, data)

		if !errors.Is(err, ErrInvalidFieldRange) {
			t.Fatalf("expected ErrInvalidFieldRange for -2, got %v", err)
		}

	})

	t.Run("getRangeEnd-invalid", func(t *testing.T) {

		pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
		data := new(Data)

		_, err := getRangeEnd([]string{"2", "a"}, 2, data)

		if !errors.Is(err, ErrInvalidFieldRange) {
			t.Fatalf("expected ErrInvalidFieldRange for invalid end, got %v", err)
		}

	})
}

func TestSortFields(t *testing.T) {

	pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	fieldsMap := map[int]struct{}{
		0:  {},
		-1: {},
		5:  {},
		2:  {},
	}

	var fields []int
	sortFields(&fields, fieldsMap)

	want := []int{2, 5}

	if !reflect.DeepEqual(fields, want) {
		t.Fatalf("sortFields result mismatch: want %v, got %v", want, fields)
	}

}
