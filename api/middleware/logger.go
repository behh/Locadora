package middleware

import (
	"io"
	"log"
	"os"

	"github.com/kardianos/osext"
	"github.com/urfave/negroni"
	"gopkg.in/natefinch/lumberjack.v2"
)

func APILogger() negroni.Handler {
	apiLogger := negroni.NewLogger()
	apiLogger.ALogger = formataLogs()

	return apiLogger
}

func formataLogs() negroni.Logger {
	pwd, err := osext.ExecutableFolder()
	if err != nil {
		log.Panic(err)
	}
	apiLogHandler := io.MultiWriter(&lumberjack.Logger{
		Filename:   pwd + "/Api.log",
		MaxSize:    200,
		MaxBackups: 5,
		MaxAge:     0,
		LocalTime:  true,
	}, os.Stdout)

	apiLogger := negroni.Logger{ALogger: log.New(apiLogHandler, "[LocadoraAPI] ", 0)}
	apiLogger.SetFormat(negroni.LoggerDefaultFormat)

	return apiLogger
}
