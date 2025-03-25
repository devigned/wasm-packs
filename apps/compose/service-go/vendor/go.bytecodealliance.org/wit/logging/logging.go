package logging

import (
	"io"
	"log"
	"math"
)

// Level represents a logging level, identical to [slog.Level].
type Level int

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
	LevelNever Level = math.MaxInt
)

// Logger represents a simple logging interface.
type Logger interface {
	// Level returns the current logging level for this Logger.
	Level() Level

	// Logf logs a message with a logging level.
	Logf(level Level, format string, v ...any)

	// Debugf logs a message with level debug.
	Debugf(format string, v ...any)

	// Infof logs a message with level info.
	Infof(format string, v ...any)

	// Printf is an alias for Infof.
	Printf(format string, v ...any)

	// Warnf logs a message with level warn.
	Warnf(format string, v ...any)

	// Errorf logs a message with level error.
	Errorf(format string, v ...any)
}

// DiscardLogger returns a logger that discards all output.
func DiscardLogger() Logger {
	return &logger{level: LevelNever}
}

// NewLogger returns a new leveled logger that writes to out.
func NewLogger(out io.Writer, level Level) Logger {
	return &logger{
		level:  level,
		logger: log.New(out, "", 0),
	}
}

type logger struct {
	level  Level
	logger *log.Logger
}

func (l *logger) Level() Level {
	return l.level
}

func (l *logger) Logf(level Level, format string, v ...any) {
	if l.level > level || l.logger == nil {
		return
	}
	l.logger.Printf(format, v...)
}

func (l *logger) Debugf(format string, v ...any) {
	l.Logf(LevelDebug, format, v...)
}

func (l *logger) Infof(format string, v ...any) {
	l.Logf(LevelInfo, format, v...)
}

func (l *logger) Printf(format string, v ...any) {
	l.Logf(LevelInfo, format, v...)
}

func (l *logger) Warnf(format string, v ...any) {
	l.Logf(LevelWarn, format, v...)
}

func (l *logger) Errorf(format string, v ...any) {
	l.Logf(LevelError, format, v...)
}
