package gormlogformatter

import (
	"fmt"
	"regexp"

	"github.com/artiomgiza/go-color-256"

	"github.com/jinzhu/gorm"
)

///////////////////////////////////////////////////////////
// Usage //////////////////////////////////////////////////
//
// Default:
// ...DB.SetLogger(gormlog.DefaultLogger())
//
// Custom:
// o := DefaultFormatOptions
// o.LogLevelColor = 123
// o.DontShowDate = false
// ...DB.SetLogger(gormlog.CustomLogger(o))
//
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

type GormLogger interface {
	Print(v ...interface{})
}

func DefaultLogger() GormLogger {
	return gormLog{
		formatOptions: DefaultFormatOptions,
	}
}

func CustomLogger(formatOptions FormatOptions) GormLogger {
	return gormLog{
		formatOptions: formatOptions,
	}
}

type FormatOptions struct {
	SourceFilePathDepth int
	DontShowDate        bool
	LogLevelName        string

	EnableColors         bool
	ErrorColor           int
	TimeColor            int
	DurationColor        int
	LogLevelColor        int
	RowsNoAffectedColor  int
	RowsYesAffectedColor int
	FilePathColor        int
	QueryColor           int
}

var DefaultFormatOptions = FormatOptions{
	SourceFilePathDepth: 3,
	DontShowDate:        true,
	LogLevelName:        "SQL",

	EnableColors:         true,
	ErrorColor:           196, // red
	TimeColor:            23,
	DurationColor:        23,
	LogLevelColor:        23,
	RowsNoAffectedColor:  23,
	RowsYesAffectedColor: 35,
	FilePathColor:        23,
	QueryColor:           60,
}

//////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////

type gormLog struct {
	formatOptions FormatOptions
}

func (l gormLog) Print(values ...interface{}) {

	ok, logData := retrieveData(l.formatOptions, values...)

	if !ok {
		fmt.Println("gorm:", values)
		return
	}

	readyLog := prepareLog(logData, l.formatOptions)

	line := fmt.Sprintf("%v %v %v %v %v %v",
		readyLog.time,
		readyLog.duration,
		readyLog.logLevel,
		readyLog.rowsAffected,
		readyLog.sourceFilePath,
		readyLog.query,
	)

	fmt.Println(line)
}

func prepareLog(data logData, formatOptions FormatOptions) logData {

	res := logData{
		time:           data.time,
		duration:       data.duration,
		logLevel:       data.logLevel,
		rowsAffected:   affectedRowsPart(data.rowsAffected),
		sourceFilePath: sourceFilePathPart(data.sourceFilePath),
		query:          data.query,
		error:          data.error,
	}

	if formatOptions.EnableColors {
		rowsAffectedColor := formatOptions.RowsNoAffectedColor
		if res.rowsAffected != affectedRowsPart("0") {
			rowsAffectedColor = formatOptions.RowsYesAffectedColor
		}

		timeColor := formatOptions.TimeColor
		durationColor := formatOptions.DurationColor
		logLevelColor := formatOptions.LogLevelColor
		queryColor := formatOptions.QueryColor
		if res.error != nil {
			timeColor = formatOptions.ErrorColor
			durationColor = formatOptions.ErrorColor
			logLevelColor = formatOptions.ErrorColor
			queryColor = formatOptions.ErrorColor
		}

		res = logData{
			time:           color.AddColor(res.time, timeColor),
			duration:       color.AddColor(res.duration, durationColor),
			logLevel:       color.AddColor(res.logLevel, logLevelColor),
			rowsAffected:   color.AddColor(res.rowsAffected, rowsAffectedColor),
			sourceFilePath: color.AddColor(res.sourceFilePath, formatOptions.FilePathColor),
			query:          color.AddColor(res.query, queryColor),
		}
	}

	return res
}

func affectedRowsPart(rowsAffected string) string {
	const affectedRowsPattern = "[%v rows]"
	return fmt.Sprintf(affectedRowsPattern, rowsAffected)
}

func sourceFilePathPart(sourceFilePath string) string {
	const sourceFilePattern = "...%v |"
	return fmt.Sprintf(sourceFilePattern, sourceFilePath)
}

func getByRegex(val interface{}, r string) string {
	raw, ok := val.(string)
	if !ok {
		return fmt.Sprintf("%v", val)
	}

	re := regexp.MustCompile(r)
	match := re.FindStringSubmatch(raw)
	if match == nil || len(match) < 2 {
		return raw
	}
	return match[1]
}

func retrieveData(formatOptions FormatOptions, values ...interface{}) (knownFormat bool, data logData) {

	////////////////////////////////////////////////////////////////////////////////////////////////////////
	// values: /////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	// []interface {}{
	//  "sql",
	//  "/Users/user/go/src/github.com/repo_name/services/lib/db_store.go:36",
	//  5234688,
	//  "TRUNCATE TABLE wave_dispatching_state CONTINUE IDENTITY CASCADE;",
	//  []interface {}{},
	//  0,
	// }

	vals := gorm.LogFormatter(values...)

	////////////////////////////////////////////////////////////////////////////////////////////////////////
	// vals: /////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	// []interface {}{
	// 	   "\x1b[35m(/Users/user/go/src/github.com/repo_name/services/lib/db_store.go:36)\x1b[0m",
	// 	   "\n\x1b[33m[2019-02-26 00:19:25]\x1b[0m",
	// 	   " \x1b[36;1m[2.70ms]\x1b[0m ",
	// 	   "DELETE from auth_sessions where EXTRACT(EPOCH FROM (current_timestamp - session_expiry)) > '14400'",
	// 	   " \n\x1b[36;31m[0 rowsAffected affected or returned ]\x1b[0m ",
	//  }

	// OR

	//[]interface {}{
	//	"\x1b[35m(/Users/gizatullinartiom/go/src/github.com/gtforge/global_ride_management_service/pkg/partition_manager/partition_manager_dao.go:55)\x1b[0m",
	//	"\n\x1b[33m[2019-03-25 17:46:16]\x1b[0m",
	//	"\x1b[31;1m",
	//	&pq.Error{...},
	//	"\x1b[0m",
	//}

	if len(vals) != 5 {
		return false, logData{}
	}

	// last N parts of the file path. For N=3: xxx(/a1/a2/a3/a4/a5/a6/file.go)xxx -> "a5/a6/file.go"
	path := getByRegex(vals[0], fmt.Sprintf("((\\/(\\w|\\.|\\-)+){%v}\\.go:\\d+)", formatOptions.SourceFilePathDepth))

	// time from "date time": xxx[2019-02-26 00:19:25]xxx -> "00:19:25" or "2019-02-26 00:19:25"
	var time string
	if formatOptions.DontShowDate {
		time = getByRegex(vals[1], "(\\d{2}:\\d{2}:\\d{2})")
	} else {
		time = getByRegex(vals[1], "(\\d{4}-\\d{2}-\\d{2}.\\d{2}:\\d{2}:\\d{2})")
	}

	if err, ok := vals[3].(error); ok {
		return true, logData{
			error:          err,
			duration:       "FAILED",
			time:           time,
			logLevel:       formatOptions.LogLevelName,
			rowsAffected:   "0",
			sourceFilePath: path,
			query:          fmt.Sprint(err),
		}
	}

	// duration as is without surroundings
	duration := getByRegex(vals[2], "\\[(\\d+\\.\\d*\\w+)\\]")

	// query as is:
	query, ok := vals[3].(string)
	if !ok {
		query = fmt.Sprint(vals[3])
	}

	// only number of affected rows: xxx[777 rowsAffected affected or returned]xxx -> "777"
	rowsAffected := getByRegex(vals[4], "(\\d+)\\srows affected")

	return true, logData{
		error:          nil,
		time:           time,
		duration:       duration,
		logLevel:       formatOptions.LogLevelName,
		rowsAffected:   rowsAffected,
		sourceFilePath: path,
		query:          query,
	}

}

type logData struct {
	error          error // used to determine failure in query execution
	time           string
	duration       string
	logLevel       string
	rowsAffected   string
	sourceFilePath string
	query          string
}
