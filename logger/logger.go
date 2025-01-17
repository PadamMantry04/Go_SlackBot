// create a wrapper function that returns a logger instance for better logging
package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func InitLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return zerolog.New(os.Stderr).With().Timestamp().Logger()
}
