package logger

import (
	"io"
	"log"
	"os"
)

var (
	DEBUG   = 0
	INFO    = 1
	WARNING = 2
	ERROR   = 3
	FATAL   = 4

	logLevel int
)

func InitLogger(level int) {
	logLevel = level
	log.SetFlags(log.LstdFlags)
}

func Writer() io.Writer {
	return os.Stdout
}

func Debug(format string, v ...interface{}) {
	if logLevel <= DEBUG {
		log.Printf("[DEBUG] "+format, v...)
	}
}

func Info(format string, v ...interface{}) {
	if logLevel <= INFO {
		log.Printf("[INFO] "+format, v...)
	}
}

func Warning(format string, v ...interface{}) {
	if logLevel <= WARNING {
		log.Printf("[WARNING] "+format, v...)
	}
}

func Error(format string, v ...interface{}) {
	if logLevel <= ERROR {
		log.Printf("[ERROR] "+format, v...)
	}
}

func Fatal(format string, v ...interface{}) {
	if logLevel <= FATAL {
		log.Fatalf("[FATAL] "+format, v...)
	}
}
