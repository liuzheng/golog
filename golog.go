package golog

import (
	"github.com/op/go-logging"
	"os"
	"path/filepath"
	"flag"
)

var (
	Log = logging.MustGetLogger("example")

	backendformat = logging.MustStringFormatter(
		`%{time:2006-01-02 15:04:05.000} %{level} %{message}`,
	)

	frontendformat = logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} %{level} %{message}%{color:reset}`,
	)
	debug = false
	logLevel = flag.String("loglevel", "INFO", "set the console log level")
	logPath = flag.String("logpath", "", "set the logfile path")
	logSelector = flag.String("logselector", "*", "Using select string to filter the log")
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func Initial() {
	Logs(*logPath, *logLevel, "INFO")
}
func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func Logs(logpath, frontend, backend string) (*logging.Logger, error) {
	if frontend == "DEBUG" {
		frontendformat = logging.MustStringFormatter(
			`%{color}%{time:2006-01-02 15:04:05.000} %{callpath} %{id:03x}%{color:reset} %{message}%{color:reset}`,
		)
		debug = true
	}
	if backend == "DEBUG" {
		backendformat = logging.MustStringFormatter(
			`%{time:2006-01-02 15:04:05.000} %{callpath} %{id:03x} %{message}`,
		)
		debug = true
	}
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
	if *logSelector == "*" || *logSelector == selector {
		Log.Debugf("[%v] " + format, append([]interface{}{selector}, v...)...)
	}
}

func Info(selector string, format string, v ...interface{}) {
	if *logSelector == "*" || *logSelector == selector {
		if debug {
			Log.Infof("[%v] " + format, append([]interface{}{selector}, v...)...)
		} else {
			Log.Infof(format, v...)
		}
	}
}

func Notice(selector string, format string, v ...interface{}) {
	if *logSelector == "*" || *logSelector == selector {
		if debug {
			Log.Noticef("[%v] " + format, append([]interface{}{selector}, v...)...)
		} else {
			Log.Noticef(format, v...)
		}
	}
}

func Warn(selector string, format string, v ...interface{}) {
	if *logSelector == "*" || *logSelector == selector {
		if debug {
			Log.Warningf("[%v] " + format, append([]interface{}{selector}, v...)...)
		} else {
			Log.Warningf(format, v...)
		}
	}
}

func Error(selector string, format string, v ...interface{}) {
	if *logSelector == "*" || *logSelector == selector {
		if debug {
			Log.Errorf("[%v] " + format, append([]interface{}{selector}, v...)...)
		} else {
			Log.Errorf(format, v...)
		}
	}
}

func Critical(selector string, format string, v ...interface{}) {
	if *logSelector == "*" || *logSelector == selector {
		if debug {
			Log.Criticalf("[%v] " + format, append([]interface{}{selector}, v...)...)
		} else {
			Log.Criticalf(format, v...)
		}
	}
}

func Panic(selector string, format string, v ...interface{}) {
	if *logSelector == "*" || *logSelector == selector {
		if debug {
			Log.Panicf("[%v] " + format, append([]interface{}{selector}, v...)...)
		} else {
			Log.Panicf(format, v...)
		}
	}
}

func Fatal(selector string, format string, v ...interface{}) {
	if *logSelector == "*" || *logSelector == selector {
		if debug {
			Log.Fatalf("[%v] " + format, append([]interface{}{selector}, v...)...)
		} else {
			Log.Fatalf(format, v...)
		}
	}
}
