package Default

import (
	"fmt"
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/Logger"
	"github.com/ian-kent/go-log/appenders"
	"github.com/ian-kent/go-log/layout"
	"github.com/ian-kent/go-log/levels"
	"github.com/ian-kent/go-log/log"
	"github.com/ian-kent/go-log/logger"
	"time"
)

type defaultLogger struct {
	logger logger.Logger
}

func (d *defaultLogger) formatNowTime() string {
	now := time.Now()
	_, offset := now.Zone()
	var timezoneSign string
	if offset >= 0 {
		timezoneSign = "+"
	} else {
		timezoneSign = "-"
	}
	return fmt.Sprintf("%s %s%d", now.Local().Format("2006-01-02 15:04:05"), timezoneSign, offset/(60*60))
}

func (d *defaultLogger) combineParams(msg string, params ...interface{}) []interface{} {
	combined := []interface{}{
		fmt.Sprintf("{%s} %s", d.formatNowTime(), msg),
	}
	combined = append(combined, params...)
	return combined
}

func (d *defaultLogger) Trace(msg string, params ...interface{}) {
	d.logger.Trace(d.combineParams(msg, params...)...)
}

func (d *defaultLogger) Debug(msg string, params ...interface{}) {
	d.logger.Debug(d.combineParams(msg, params...)...)
}

func (d *defaultLogger) Info(msg string, params ...interface{}) {
	d.logger.Info(d.combineParams(msg, params...)...)
}

func (d *defaultLogger) Warn(msg string, params ...interface{}) {
	d.logger.Warn(d.combineParams(msg, params...)...)
}

func (d *defaultLogger) Error(msg string, params ...interface{}) {
	d.logger.Error(d.combineParams(msg, params...)...)
}

func (d *defaultLogger) Fatal(msg string, params ...interface{}) {
	d.logger.Fatal(d.combineParams(msg, params...)...)
}

func getPrefixWithSpace(prefix string) string {
	if prefix == "" {
		return ""
	} else {
		return prefix + " "
	}
}

func New(logFileName, prefix string, isDevMode bool) Logger {
	logger := log.Logger()

	if isDevMode {
		logger.SetLevel(levels.TRACE)
	} else {
		logger.SetLevel(levels.INFO)
	}

	layoutToUse := layout.Pattern(getPrefixWithSpace(prefix) + "[%p] %m") //level/priority, message

	rollingFileAppender := appenders.RollingFile(logFileName, true)
	rollingFileAppender.MaxBackupIndex = 5
	rollingFileAppender.MaxFileSize = 20 * 1024 * 1024 // 20 MB
	rollingFileAppender.SetLayout(layoutToUse)

	consoleAppender := appenders.Console()
	consoleAppender.SetLayout(layoutToUse)
	logger.SetAppender(
		Multiple( //appenders.Multiple( ONCE PULL REQUEST OF ABOVE IS IN
			layoutToUse,
			rollingFileAppender,
			consoleAppender,
		))

	return &defaultLogger{logger}
}
