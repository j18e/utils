package utils

// Logger is used for writing logs.
type Logger interface {
	Info(...interface{})
	Errorf(string, ...interface{})
}
