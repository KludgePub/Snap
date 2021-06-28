package logger

import "log"

const debugTag = "|DEBUG|"

// Logger for debug output
type Logger struct {
	IsDebug bool
}

// LogDebug message
func (l *Logger) LogDebug(msg string) {
	log.Printf("%s %s\n", debugTag, msg)
}

// LogDebugWithObject message and any object
func (l *Logger) LogDebugWithObject(msg string, object interface{}) {
	log.Printf("%s %s object - %v\n", debugTag, msg, object)
}
