package logrusformatter

import (
	"strings"

	"github.com/artiomgiza/go-color-256"
)

type ColorOptions struct {
	ColorsEnabled func() bool

	LogMessageColor  int
	LogFieldKeyColor int
	LogFieldValColor int

	ColorizeErrField bool

	// log levels:
	LevelTraceColor   int
	LevelDebugColor   int
	LevelInfoColor    int
	LevelWarningColor int
	LevelErrorColor   int
	LevelFatalColor   int
	LevelPanicColor   int

	OverrideLogColor func(m string) (bool, int)
}

// Color schema:
// https://github.com/artiomgiza/go-color-256/blob/master/colors.png
// Readme for more...

var DefaultColorOptions = ColorOptions{
	ColorsEnabled: func() bool { return true },

	LogMessageColor:  34,  // sort of green
	LogFieldKeyColor: 62,  // dark purple
	LogFieldValColor: 244, // dark gray
	ColorizeErrField: true,

	// Header color (different per log logLevel)
	LevelTraceColor:   146, // almost white
	LevelDebugColor:   25,  // king of blue
	LevelInfoColor:    32,  // kind of green
	LevelWarningColor: 184, // king of yellow
	LevelErrorColor:   196, // red
	LevelFatalColor:   196, // red
	LevelPanicColor:   196, // red

	OverrideLogColor: func(m string) (bool, int) { return false, 0 },
}

func buildColorLog(l rawLogData, colorOptions ColorOptions) readyLogData {

	headerColor := headerColor(l.logLevel, colorOptions)
	msgColor := colorOptions.LogMessageColor
	fieldKeyColor := colorOptions.LogFieldKeyColor
	fieldValColor := colorOptions.LogFieldValColor

	// if should override - apply given override color to whole log line
	if shouldOverride, overrideColor := colorOptions.OverrideLogColor(l.logMsg); shouldOverride {
		headerColor = overrideColor
		msgColor = overrideColor
		fieldKeyColor = overrideColor
		fieldValColor = overrideColor
	}

	errFieldColor := 0
	if colorOptions.ColorizeErrField {
		errFieldColor = colorOptions.LevelErrorColor
	}

	return readyLogData{
		logTime:      color.AddColor(l.logTime, headerColor),
		logID:        color.AddColor(l.logID, headerColor),
		logLevel:     color.AddColor(l.logLevel, headerColor),
		logMsg:       color.AddColor(l.logMsg, msgColor),
		logFieldsStr: colorizeFields(l.logFields, fieldKeyColor, fieldValColor, errFieldColor),
	}
}

func headerColor(logLevel string, co ColorOptions) int {
	headerColorMap := map[string]int{
		"TRA": co.LevelTraceColor,
		"DEB": co.LevelDebugColor,
		"INF": co.LevelInfoColor,
		"WAR": co.LevelWarningColor,
		"ERR": co.LevelErrorColor,
		"FAT": co.LevelFatalColor,
		"PAN": co.LevelPanicColor,
	}

	headerColor, ok := headerColorMap[strings.ToUpper(logLevel)[:3]]
	if !ok {
		headerColor = 254 // almost white
	}
	return headerColor
}

func colorizeFields(pairs []field, kCol, vCol int, errFieldColor int) string {
	if len(pairs) == 0 {
		return ""
	}

	var coloredPairs []string
	for _, p := range pairs {
		var sb strings.Builder
		if p.key == "error" && errFieldColor != 0 {
			// override error fields with error color
			sb.WriteString(color.AddColor(p.key, errFieldColor))
			sb.WriteString(color.AddColor(": ", vCol))
			sb.WriteString(color.AddColor(p.val, errFieldColor))
		} else {
			sb.WriteString(color.AddColor(p.key, kCol))
			sb.WriteString(color.AddColor(": ", vCol))
			sb.WriteString(color.AddColor(p.val, vCol))
		}
		coloredPairs = append(coloredPairs, sb.String())
	}

	var sb strings.Builder
	sb.WriteString(color.AddColor("{", vCol))
	sb.WriteString(strings.Join(coloredPairs, color.AddColor(", ", vCol)))
	sb.WriteString(color.AddColor("}", vCol))
	return sb.String()
}
