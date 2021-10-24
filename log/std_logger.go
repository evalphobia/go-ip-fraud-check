package log

import "log"

// StdLogger use standard log package.
type StdLogger struct{}

// Infof logging information.
func (*StdLogger) Infof(format string, v ...interface{}) {
	log.Printf("[INFO] "+format, v...)
}

// Errorf logging error information.
func (*StdLogger) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR] "+format, v...)
}
