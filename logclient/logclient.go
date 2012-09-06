// The logclient package is a simple example of dependency injection.
//
// Dependency Injection (DI) is the practice of passing dependencies
// (objects and interfaces required by an object) to the constructor
// of the object.  Dependency Injection aids testing by allowing test
// code to directly control all dependencies, and aids modularity by
// not hard-coding dependencies into the constructor.
//
package logclient

// The Writer interface is used by this package to write to the log.
// In general, packages may define any interfaces required.  The
// implementation that satisfies the interface may be unaware of the
// interface; to satisfy the interface it need only supply the
// corresponding methods.
//
type Writer interface {
	Write([]byte) (int, error)
}

// Type LogClient records a single dependency: The log that the client
// will write to.  In general, any number of dependencies may be
// recorded.
//
type LogClient struct {
	log Writer // A Recorded Dependency
}

// New constructs the log client, recording the single injected
// dependency: the log written.  In general, multiple dependencies may
// be injected, and any subset of these dependencies may be recorded
// for use after the constructor returns.
//
func New(log Writer) *LogClient {
	// Return a new LogClient that records the injected
	// dependency for later use:
	return &LogClient{log}
}

// LogSomething writes "Something\n" to the injected log.
//
func (c *LogClient) LogSomething() {
	c.log.Write([]byte("Something\n"))
}
