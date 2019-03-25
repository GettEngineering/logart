# logrus formatter


![alt text](https://github.com/gtforge/logart/blob/master/log-formatters/logrus-human-formatter/readme_files/formatted.png "Example")

### Usage:

`go get github.com/gtforge/logart/log-formatters/logrus-human-formatter`

Default setup:
```
logrushumanformatter.Set()
```

Default setup:
```
logrushumanformatter.SetWithLogIDProvider(logIDProvider)

// when logIDProvider is a function with no input params and string output
// and should provide kind of sessionID (unique number per running scope)
```



Custom setup:
```
formatOptions := DefaultFormatOptions
formatOptions.Option1 = customValue1
formatOptions.Option2 = customValue2
...

colorOptions := DefaultColorOptions
colorOptions.Option1 = customValue1
colorOptions.Option2 = customValue2
...

SetCustomized(formatOptions, colorOptions)
```

Format options allow flexible output configuration:

- TimeLayout - `default = 15:04:05.000` (no date, time with ms).
Also could be set with date: "2006-01-02 15:04:05.000". This is the time
format we'll see in the log.

- LogLevelLength - `default = 3` (DEB / INF / WAR / ERR / ...).
Number of letters will be used to show log level.

- LogIDProvider - `default = empty string`.
Actually this is a function that returns string. You can set it to return
current "RequestID" or "SessionID" - depends infra code is running in.

- FirstEverPrintedFields - `default = none`. Sometimes we want print specific
log fields in defined order. By setting this option you'll force defined fields
be printed in defined order before other fields.

- LastEverPrintedFields - `default = none`. Same as previous, but "after other
fields". Fields that are not defined in First/Last... option, would be printed
between them sorted alphabetically.


Color options allow flexible output colorization ([available colors](https://github.com/artiomgiza/go-color-256)):

- ColorsEnabled - `default = true`.
Actually this is a function that return boolean. This function allows
dynamic color enabling/disabling. Useful to differentiate between prod
and dev/stage environments. Most likely the colorization will be enabled on
dev/stage (output is the terminal window - we want nice colored log) and
disabled on prod (log management systems)

- LogMessageColor - `default = green`. Log message color

- LogFieldKeyColor - `default = dark purple`.
When log fields are printed, keys and values are printed in different colors
to make the log readable. See attached colorized log example.

- LogFieldValColor - `default = dark gray`. See previous.

- ColorizeErrField - `default = true`.
Previous two options define the color scheme of the log fields. To emphasize
the error field - it will be printed in same color error header will be print.

- Log level color settings:
    - LevelTraceColor   - `default = white`
    - LevelDebugColor   - `default = blue`
    - LevelInfoColor    - `default = green`
    - LevelWarningColor - `default = yellow`
    - LevelErrorColor   - `default = red`
    - LevelFatalColor   - `default = red`
    - LevelPanicColor   - `default = red`

- OverrideLogColor - `default = none`.
Actually this is a function that return boolean (should override) and
integer (override color). This function allows you hide (by printing whole
log in color similar to the background. Why we would use this? For
example infra logs that could not be removed, but are not important to see.
The log line will exist, but it will be "hidden" for human eye.

\* 256 colors are defined (0..255) - see the colors mapping
[here](https://github.com/artiomgiza/go-color-256)

\* To choose the best color for your terminal's background color, you
can run short bash script ([same link](https://github.com/artiomgiza/go-color-256))
in the checked terminal to see how each color looks on it.

### Note

For most of the usages in real life, it can be useful to differentiate between prod
and dev/stage environments. Most likely this formatter will be enabled on
dev/stage (output is the terminal window - we want nice formatted log) and
disabled on prod (log management systems). Of course on prod env we'll use
another type of formatter, such as json formatter ([like this one](https://github.com/gtforge/logart/tree/master/log-formatters/logrus-json-formatter))

### Comparison:

Here is couple of possible options to format the log. **Same** log is shown
in three different ways:

Colorized log (made by this formatter):
![alt text](https://github.com/gtforge/logart/blob/master/log-formatters/logrus-human-formatter/readme_files/formatted.png "Example")

Default log (logrus default):
![alt text](https://github.com/gtforge/logart/blob/master/log-formatters/logrus-human-formatter/readme_files/default-formatter.png "Example")

JSON formatted log (log is printed as JSON):
![alt text](https://github.com/gtforge/logart/blob/master/log-formatters/logrus-human-formatter/readme_files/json-formatter.png "Example")
