package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func New(levelValue, format string) (zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(strings.ToLower(levelValue))
	if err != nil {
		return zerolog.Logger{}, fmt.Errorf("parse log level %q: %w", levelValue, err)
	}

	var output io.Writer
	switch strings.ToLower(format) {
	case "json":
		output = os.Stdout
	case "console":
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	default:
		return zerolog.Logger{}, fmt.Errorf("unsupported log format %q: use json or console", format)
	}

	log := zerolog.New(output).
		Level(level).
		With().
		Timestamp().
		Logger()

	return log, nil
}
