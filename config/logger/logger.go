package logger

import (
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

type Logger struct {
	Writter zerolog.Logger
}

var Log Logger

func CreateLog(appName string) {
	logger := httplog.NewLogger(appName, httplog.Options{
		JSON: true,
	})
	Log.Writter = logger
}
