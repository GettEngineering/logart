package logrushumanformatter

import (
	"fmt"

	"strings"

	"github.com/sirupsen/logrus"
)

func Set() {
	logrus.SetFormatter(formatter{
		formatOptions: DefaultFormatOptions,
		colorOptions:  DefaultColorOptions,
	})
}

func SetWithLogIDProvider(logIDProvider func() string) {
	formatOptions := DefaultFormatOptions
	formatOptions.LogIDProvider = logIDProvider

	logrus.SetFormatter(formatter{
		formatOptions: formatOptions,
		colorOptions:  DefaultColorOptions,
	})
}

func SetCustomized(formatOptions FormatOptions, colorOptions ColorOptions) {
	logrus.SetFormatter(formatter{
		formatOptions: formatOptions,
		colorOptions:  colorOptions,
	})
}

func (f formatter) Format(entry *logrus.Entry) ([]byte, error) {
	rawLogData := rawLogData{
		logTime:   logTime(entry, f.formatOptions),
		logID:     logID(entry, f.formatOptions),
		logLevel:  logLevel(entry, f.formatOptions),
		logMsg:    logMsg(entry, f.formatOptions),
		logFields: logFields(entry, f.formatOptions),
	}

	var readyLogData readyLogData
	if f.colorOptions.ColorsEnabled() {
		readyLogData = buildColorLog(rawLogData, f.colorOptions)
	} else {
		readyLogData = buildLog(rawLogData)
	}

	logLine := fmt.Sprintf("%v %v %v %v %v\n", readyLogData.logTime, readyLogData.logID, readyLogData.logLevel, readyLogData.logMsg, readyLogData.logFieldsStr)
	return []byte(logLine), nil
}

type formatter struct {
	formatOptions FormatOptions
	colorOptions  ColorOptions
}

func logTime(e *logrus.Entry, o FormatOptions) string {
	return e.Time.UTC().Format(o.TimeLayout)
}

func logLevel(e *logrus.Entry, o FormatOptions) string {
	level := strings.ToUpper(e.Level.String())
	for len(level) < o.LogLevelLength {
		level += "  " // padding with spaces if needed
	}
	if o.LogLevelLength < 3 {
		o.LogLevelLength = 3
	}
	return level[:o.LogLevelLength]
}

func logID(_ *logrus.Entry, o FormatOptions) string {
	return o.LogIDProvider()
}

func logMsg(e *logrus.Entry, o FormatOptions) string {
	return e.Message
}

// formatted, but not ready to be used
// need organize fields and add color (if needed)
type rawLogData struct {
	logTime   string
	logID     string
	logLevel  string
	logMsg    string
	logFields []field
}

// data with well formatted fields and colors
type readyLogData struct {
	logTime      string
	logID        string
	logLevel     string
	logMsg       string
	logFieldsStr string
}
