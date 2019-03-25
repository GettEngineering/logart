package logrushumanformatter

type FormatOptions struct {
	TimeLayout     string
	LogLevelLength int
	LogIDProvider  func() string

	FirstEverPrintedFields OrderedFields
	LastEverPrintedFields  OrderedFields
}

var DefaultFormatOptions = FormatOptions{
	TimeLayout:     "15:04:05.000", // could be with date: "2006-01-02 15:04:05.000"
	LogLevelLength: 3,              // DEB / INF / WAR / ERR / ...
	LogIDProvider:  func() string { return "" },

	FirstEverPrintedFields: OrderedFields{},
	LastEverPrintedFields:  OrderedFields{},
}

func buildLog(l rawLogData) readyLogData {
	return readyLogData{
		logTime:      l.logTime,
		logID:        l.logID,
		logLevel:     l.logLevel,
		logMsg:       l.logMsg,
		logFieldsStr: buildFields(l.logFields),
	}
}
