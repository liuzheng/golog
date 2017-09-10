package golog

import (
	"github.com/op/go-logging"
	"os"
	"path/filepath"
	"flag"
)

var Log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var backendformat = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05.000} %{callpath} %{id:03x} %{message}`,
)
var frontendformat = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02 15:04:05.000} %{callpath} %{id:03x}%{color:reset} %{message}`,
)
var (
	LogLevel = flag.String("logleve", "INFO", "set the console log level")
	LogPath = flag.String("logpath", "", "set the logfile path")
	LogSelector = flag.String("logselector", "*", "Using select string to filter the log")
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func init() {
	flag.Parse()
	Logs(*LogPath, *LogLevel, "INFO")
}
func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func Logs(logpath, frontend, backend string) (*logging.Logger, error) {
	Frontend := logging.NewLogBackend(os.Stderr, "", 0)
	FrontendFormatter := logging.NewBackendFormatter(Frontend, frontendformat)
	FrontendLeveled := logging.AddModuleLevel(FrontendFormatter)
	level, _ := logging.LogLevel(frontend)
	FrontendLeveled.SetLevel(level, "")
	if logpath != "" {
		if _, err := os.Stat(logpath); os.IsNotExist(err) {
			// path/to/whatever does not exist
			filePathDir := filepath.Dir(logpath)
			if _, err = os.Stat(filePathDir); os.IsNotExist(err) {
				err = os.MkdirAll(filePathDir, 0755)
				if err != nil {
					return nil, err
				}
			}

		}
		f, err := os.OpenFile(logpath, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		Backend := logging.NewLogBackend(f, "", 0)
		BackendFormatter := logging.NewBackendFormatter(Backend, backendformat)
		BackendLeveled := logging.AddModuleLevel(BackendFormatter)
		level, _ := logging.LogLevel(backend)
		BackendLeveled.SetLevel(level, "")
		logging.SetBackend(BackendLeveled)
		logging.SetBackend(FrontendLeveled, BackendLeveled)
	} else {
		logging.SetBackend(FrontendLeveled)
	}
	return Log, nil
}

func Debug(selector string, format string, v ...interface{}) {
	if *LogSelector == "*" || *LogSelector == selector {
		Log.Errorf("[%v] " + format, []interface{}{selector, v}...)
	}
}
func Info(selector string, format string, v ...interface{}) {
	if *LogSelector == "*" || *LogSelector == selector {
		Log.Infof("[%v] " + format, []interface{}{selector, v}...)
	}
}

func Warn(selector string, format string, v ...interface{}) {
	if *LogSelector == "*" || *LogSelector == selector {
		Log.Warningf("[%v] " + format, []interface{}{selector, v}...)
	}
}

func Error(selector string, format string, v ...interface{}) {
	if *LogSelector == "*" || *LogSelector == selector {
		Log.Errorf("[%v] " + format, []interface{}{selector, v}...)
	}
}

func Critical(selector string, format string, v ...interface{}) {
	if *LogSelector == "*" || *LogSelector == selector {
		Log.Criticalf("[%v] " + format, []interface{}{selector, v}...)
	}
}
func Notice(selector string, format string, v ...interface{}) {
	if *LogSelector == "*" || *LogSelector == selector {
		Log.Noticef("[%v] " + format, []interface{}{selector, v}...)
	}
}