package logclient

import (
	"io"
	"os"
)

// ExampleLogClient creates a LogClient, injecting a logging object
// that will print each log entry in triplicate.
//
func ExampleLogClient() {

	// Create an object to Write() to Stdout.

	o := io.Writer(os.Stdout)

	// That's not fancy enough, so create an object to
	// triple-write to Stdout.

	tri := io.MultiWriter(o, o, o)

	// Create a LogClient, which will log messages to m.
	// Note that the c knows nothing about m except that it implements
	// the loginfo.Writer interface.  Also, m knows nothing about
	// the loginfo.Writer interface: it is sufficient that it
	// implements a Write([]byte)(int,err) method.

	c := New(tri)

	// Cause c to write "Something\n" to its log.
	c.LogSomething()

	// Output:
	// Something
	// Something
	// Something
}
