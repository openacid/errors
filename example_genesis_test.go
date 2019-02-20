package errors_test

import (
	"fmt"
	"io"

	"github.com/openacid/errors"
)

func ExampleGenesisError() {

	err := errors.WithCauses("foo", // create a error "foo" caused by other 2 errors: EOF and read-stopped
		io.EOF,
		errors.WithCauses("read-stopped", // "read stopped" error caused by other 2 errors, too:
			io.ErrShortBuffer,
			io.EOF))
	fmt.Printf("%s", err.Error())

	// Output: foo: [EOF, read-stopped: [short buffer, EOF]]
}

func ExampleGenesisPrintf() {

	err := errors.WithCauses("foo", // create a error "foo" caused by error "read-stopped"
		errors.WithCauses("read-stopped", // "read stopped" error caused by other 2 errors, too:
			io.ErrShortBuffer,
			io.EOF))

	fmt.Println(err)
	fmt.Printf("%s\n", err)
	fmt.Printf("%q\n", err)
	fmt.Printf("%v\n", err)

	// Output:
	// foo: read-stopped: [short buffer, EOF]
	// foo: read-stopped: [short buffer, EOF]
	// "foo: read-stopped: [short buffer, EOF]"
	// foo: read-stopped: [short buffer, EOF]
}

func ExampleGenesisStack() {

	err := errors.WithCauses("foo", // create a error "foo" caused by error "read-stopped"
		errors.WithCauses("read-stopped", // "read stopped" error caused by other 2 errors, too:
			io.ErrShortBuffer,
			io.EOF))

	fmt.Printf("%+v\n", err)

	// Example output:
	// foo
	//     read-stopped
	//         short buffer
	//         EOF
	//
	// github.com/openacid/errors_test.ExampleGenesisStack
	//         /Users/drdrxp/xp/vcs/go/src/github.com/openacid/errors/example_test.go:241
	// testing.runExample
	//         /usr/local/Cellar/go/1.10.3/libexec/src/testing/example.go:122
	// testing.runExamples
	//         /usr/local/Cellar/go/1.10.3/libexec/src/testing/example.go:46
	// testing.(*M).Run
	//         /usr/local/Cellar/go/1.10.3/libexec/src/testing/testing.go:979
	// main.main
	//         _testmain.go:122
	// runtime.main
	//         /usr/local/Cellar/go/1.10.3/libexec/src/runtime/proc.go:198
	// runtime.goexit
	//         /usr/local/Cellar/go/1.10.3/libexec/src/runtime/asm_amd64.s:2361
}
