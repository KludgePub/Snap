package logger

import (
	"fmt"
	golog "log"
	"os"
	"sync"
)

const (
	debugTag = "|DEBUG|"
	verboseTag = "|Snap|"
)

var (
	logOnce sync.Once
	logSingleton *Logger
)

// Logger for debug output
type Logger struct {
	IsDebug bool

	impl *golog.Logger
}

func log() *Logger {
	logOnce.Do(func() {
		logSingleton = &Logger{
			impl: golog.New(os.Stderr, "", golog.Ltime),
		}
	})

	return logSingleton
}

// LogDebug message
func (l *Logger) LogDebug(msg string) {
	if l.IsDebug {
		log().impl.Printf("%s%s %s", verboseTag, debugTag, l.makeLog(msg))
	}
}

// LogDebugWithObject with message and object
func (l *Logger) LogDebugWithObject(msg string, object interface{}) {
	if l.IsDebug {
		log().impl.Printf("%s%s %s", verboseTag, debugTag, l.makeLogWithObject(msg, object))
	}
}

// Log message
func (l *Logger) Log(msg string) {
	log().impl.Printf("%s %s", verboseTag, l.makeLog(msg))
}

// LogWithObject with message and object
func (l *Logger) LogWithObject(msg string, object interface{}) {
	log().impl.Printf("%s %s", verboseTag, l.makeLogWithObject(msg, object))
}


func (l *Logger) makeLog(m string) string {
	return fmt.Sprintf("%s\n", m)
}

func (l *Logger) makeLogWithObject(m string, object interface{}) string {
	return fmt.Sprintf("\nMessage: %s\nObject: %v\n", m, object)
}