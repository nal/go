package logger

import (
	"errors"
	"log"
	"os"

	zap "go.uber.org/zap"
)

// This interface is for pure educational purpose.
// Implement proxy module for supporting both log and zap loggers.
type LoggerIf interface {
	Init(backend string) (logger *Logger, err error)
	Info(message string, v ...interface{})
	Defer()
}

// Logger struct
type Logger struct {
	// specify actual backend
	backend string
	// we can define here only one logger obj with reflection
	// but reflection should be avoided if possible
	// so we explicitly define supported loggers
	logLogger *log.Logger
	zapLogger *zap.Logger
}

// Init inits logger with supported backends.
// Current available backends are: log, zap.
// Pass empty backend if you don't care about backend.
// By default zap is used.
func Init(backend string) (logger *Logger, err error) {
	logger = new(Logger) // or logger = &Logger{}

	// Set default value
	if backend == "" {
		backend = "zap"
	}

	if backend == "zap" {
		logger.backend = "zap"
		logger.zapLogger, _ = zap.NewProduction()
		// logger.zapLogger = logger.zapLogger.Sugar()
	} else if backend == "log" {
		logger.backend = "log"
		logger.logLogger = log.New(os.Stdout, "", log.Lshortfile)
	} else {
		return nil, errors.New("unknown backend!")
	}

	return
}

func (l *Logger) Info(message string, v ...interface{}) {
	if l.backend == "zap" {
		l.zapLogger.Sugar().Infow(message, v...)
	} else if l.backend == "log" {
		if len(v) > 0 {
			// Merge message and v
			var v2 []interface{} = make([]interface{}, len(v)+1)
			v2[0] = message
			for index, value := range v {
				v2[index+1] = value
			}

			l.logLogger.Println(v2...)
		} else {
			l.logLogger.Println(message)
		}
	}
}

func (l *Logger) Defer() {
	l.Info("Called logger.Defer()")

	if l.backend == "zap" {
		l.zapLogger.Sync() // flushes buffer, if any
	} else if l.backend == "log" {
		// Defer() for log backend is no-op
	}
}
