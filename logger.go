package utils

// Logger is used for writing logs.
type Logger interface {
	Info(...interface{})
	Errorf(string, ...interface{})
}

// NewTestLogger does nothing and is only used for testing.
func NewTestLogger() Logger {
	return &testLogger{}
}

type testLogger struct{}

func (l *testLogger) Info(...interface{})           {}
func (l *testLogger) Errorf(string, ...interface{}) {}
