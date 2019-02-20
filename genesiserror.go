package errors

import (
	"fmt"
	"io"
	"strings"
)

type directCauser interface {
	DirectCauses() []error
}

type genesisError struct {
	code         string
	directCauses []error
	*stack
}

var (
	// TODO test it doc it
	Stack = &fundamental{msg: "flag to add stack"}
	NoStack = &fundamental{msg: "flag NOT to add stack"}
)


// TODO
func CausedBy(message string, errs ...error) error {

	var withstack bool
	
	if len(errs) == 0 {
		// There is no causing error, add stack by default.
		withstack = true
		
	} else if errs[0] == Stack {
		errs = errs[1:]
		withstack = true

	} else if errs[0] == NoStack {
		errs = errs[1:]
		withstack = false

	} else {
		withstack := true
		for _, err := range errs {
			_, ok = err.(genesisError)
			if ok {
				withstack = false
				break
			}
		}
	}

	if withstack {
		return &genesisError{
			code:         message,
			directCauses: errs,
			stack:        callers(),
		}
	} else {
		return &genesisError{
			code:         message,
			directCauses: errs,
		}
	}


}

func (w *genesisError) Cause() error {
	if len(w.directCauses) == 0 {
		return nil
	}
	return w.directCauses[0]
}

func (w *genesisError) Error() string {

	s := w.code
	causes := w.DirectCauses()
	switch len(causes) {
	case 0:
		return s
	case 1:
		return s + ": " + causes[0].Error()
	default:
		s = s + ": ["
		for i, c := range causes {
			s += c.Error()
			if i < len(causes)-1 {
				s += ", "
			}
		}
		s += "]"
		return s
	}
}

// TODO
func (w *genesisError) AddCause(err error) {
	w.directCauses = append(w.directCauses, err)
}

// TODO
func (w *genesisError) DirectCauses() []error {
	// []error(nil) is not []error{}
	// Although they have the same observable behaviors.
	// Read more: https://stackoverflow.com/questions/44305170/nil-slices-vs-non-nil-slices-vs-empty-slices-in-go-language
	//
	// To let tests be happy, convert nil slice to empty slice.
	// Only when required, we allocate a new slice.
	if w.directCauses == nil {
		return []error{}
	}
	return w.directCauses[:]
}

func printCauseTree(s fmt.State, err error, indent int) {

	indstr := strings.Repeat(" ", indent*4)
	io.WriteString(s, indstr)

	switch g := err.(type) {
	case *genesisError:

		io.WriteString(s, g.code)
		io.WriteString(s, "\n")

		causes := g.DirectCauses()
		for _, c := range causes {
			printCauseTree(s, c, indent+1)
		}
	default:
		io.WriteString(s, err.Error())
		io.WriteString(s, "\n")
	}
}

func (g *genesisError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			printCauseTree(s, g, 0)
			g.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, g.Error())
	case 'q':
		fmt.Fprintf(s, "%q", g.Error())
	}
}

func AllCauses(err error) []error {

	if err == nil {
		return nil
	}

	for err != nil {
		switch ww := err.(type) {
		case withStack:
			err = ww.Cause()
			continue
		case *withStack:
			err = ww.Cause()
			continue
		}

		ac, ok := err.(directCauser)
		if ok {
			return ac.DirectCauses()
		}

		c, ok := err.(causer)
		if ok {
			return []error{c.Cause()}
		}

		return []error{}
	}

	return nil

}
