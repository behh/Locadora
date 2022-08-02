package logger

import (
	"io"
	"log"
	"os"

	"github.com/kardianos/osext"
	"gopkg.in/natefinch/lumberjack.v2"
)

//Logger aponta para os Logs da Aplicação
type Logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

type Service struct {
	Logger Logger
}

func (s *Service) LogError(e error) {
	s.Logger.Error.Println(e)
}

func (s *Service) LogInfo(str string) {
	s.Logger.Info.Println(s)
}

func (s *Service) LogWarning(str string) {
	s.Logger.Warning.Println(s)
}

func NewService() *Service {
	log := Logger{
		Info:    &log.Logger{},
		Warning: &log.Logger{},
		Error:   &log.Logger{},
	}
	inicializaLogs(&log)
	return &Service{
		Logger: log,
	}
}

func inicializaLogs(logger *Logger) {
	pwd, err := osext.ExecutableFolder()
	if err != nil {
		log.Panic(err)
	}

	errorLogHandler := io.MultiWriter(&lumberjack.Logger{
		Filename:   pwd + "/Error.log",
		MaxSize:    200,
		MaxBackups: 5,
		MaxAge:     0,
		LocalTime:  true,
	}, os.Stdout)

	infoLogHandler := io.MultiWriter(&lumberjack.Logger{
		Filename:   pwd + "/Info.log",
		MaxSize:    200,
		MaxBackups: 5,
		MaxAge:     0,
		LocalTime:  true,
	}, os.Stdout)

	warningLogHandler := io.MultiWriter(&lumberjack.Logger{
		Filename:   pwd + "/Warning.log",
		MaxSize:    200,
		MaxBackups: 5,
		MaxAge:     0,
		LocalTime:  true,
	}, os.Stdout)

	logger.Info, logger.Warning, logger.Error = formataLogs(infoLogHandler, warningLogHandler, errorLogHandler)
}

func formataLogs(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) (*log.Logger, *log.Logger, *log.Logger) {

	info := log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning := log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	errorlog := log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	return info, warning, errorlog
}
