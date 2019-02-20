package errors

import (
	"io"
	"reflect"
	"testing"
)

type nilError struct{}

func (nilError) Error() string { return "nil error" }

func TestGenesisCauseAllCauseError()(t *testing.T) {
	x := New("error")
	tests := []struct {
		err       error
		want      error
		allcauses []error
		strwant   string
	}{{
		// nil error is nil
		err:       nil,
		want:      nil,
		allcauses: nil,
		strwant:   "",
	}, {
		// explicit nil error is nil
		err:       (error)(nil),
		want:      nil,
		allcauses: nil,
		strwant:   "",
	}, {
		// typed nil is nil
		err:       (*nilError)(nil),
		want:      (*nilError)(nil),
		allcauses: []error{},
		strwant:   "",
	}, {
		// uncaused error is unaffected
		err:       io.EOF,
		want:      io.EOF,
		allcauses: []error{},
		strwant:   "EOF",
	}, {
		// caused error returns cause
		err:       Wrap(io.EOF, "ignored"),
		want:      io.EOF,
		allcauses: []error{io.EOF},
		strwant:   "ignored: EOF",
	}, { // 6
		err:       x, // return from errors.New
		want:      x,
		allcauses: []error{},
		strwant:   "error",
	}, {
		WithMessage(nil, "whoops"),
		nil,
		nil,
		"",
	}, {
		WithMessage(io.EOF, "whoops"),
		io.EOF,
		[]error{io.EOF},
		"whoops: EOF",
	}, {
		WithStack(nil),
		nil,
		nil,
		"",
	}, {
		WithStack(io.EOF),
		io.EOF,
		// it has no causing error, just itself
		[]error{},
		"EOF",
	}, { // 11
		WithCauses("foo"),
		nil,
		[]error{},
		"foo",
	}, {
		WithCauses("foo", io.EOF),
		io.EOF,
		[]error{io.EOF},
		"foo: EOF",
	}, {
		WithCauses("foo", io.EOF, x),
		io.EOF,
		[]error{io.EOF, x},
		"foo: [EOF, error]",
	}}

	for i, tt := range tests {
		got := Cause(tt.err)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("test %d: got %#v, want %#v", i+1, got, tt.want)
		}

		allgot := AllCauses(tt.err)
		if !reflect.DeepEqual(allgot, tt.allcauses) {
			t.Errorf("test %d: AllCauses: err: %#v,  got %#v, want %#v", i+1, tt.err, allgot, tt.allcauses)
		}

		if tt.err == nil {
			continue
		}

		// Skip test about nilError. in golang its impossible to check a nil
		// interface:
		//	   var b error = (*X)(nil)
		//	   fmt.Println("b == nil:", b == nil) // false
		_, isnilerr := tt.err.(*nilError)
		if isnilerr {
			continue
		}

		strgot := tt.err.Error()
		if strgot != tt.strwant {
			t.Errorf("test %d: Error(): err: %#v,  got %#v, want %#v", i+1, tt.err, strgot, tt.strwant)
		}
	}
}
