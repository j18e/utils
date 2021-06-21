package utils_test

type logger struct{}

func (l *logger) Info(...interface{})           {}
func (l *logger) Errorf(string, ...interface{}) {}
