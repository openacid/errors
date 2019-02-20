package errors

import (
	"io"
	"reflect"
	"testing"
)

func TestGenesisCauseAllCauseError(t *testing.T) {
	x := New("error")
	basic := CausedBy("basic-error")
	tests := []struct {
		err       error
		want      error
		allcauses []error
		strwant   string
	}{{
		// nil error is nil
		nil,
		nil,
		nil,
		"",
	}, {
		// explicit nil error is nil
		(error)(nil),
		nil,
		nil,
		"",
	}, {
		// typed nil is nil
		(*nilError)(nil),
		(*nilError)(nil),
		[]error{},
		"",
	}, {
		// uncaused error is unaffected
		io.EOF,
		io.EOF,
		[]error{},
		"EOF",
	}, {
		// caused error returns cause
		Wrap(io.EOF, "ignored"),
		io.EOF,
		[]error{io.EOF},
		"ignored: EOF",
	}, { // 6
		x, // return from errors.New
		x,
		[]error{},
		"error",
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
		basic,
		basic,
		[]error{},
		"basic-error",
	}, {
		CausedBy("foo", io.EOF),
		io.EOF,
		[]error{io.EOF},
		"foo: EOF",
	}, {
		CausedBy("foo", io.EOF, x),
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
