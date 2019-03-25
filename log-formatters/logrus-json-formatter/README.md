# logrus json formatter


JSON formatted log (log is printed as JSON):
![alt text](https://github.com/gtforge/logart/blob/master/log-formatters/logrus-json-formatter/readme_files/json-formatter.png "Example")

### Usage:

Default setup:
```
Set()
```

Custom setup:
```
formatOptions := DefaultFormatOptions
formatOptions.Option1 = customValue1
formatOptions.Option2 = customValue2
...

SetCustomized(formatOptions)
```

Format options allow flexible output configurations:

- TimeLayout - `default = RFC3339 (2006-01-02T15:04:05Z07:00)`

- LogIDProvider - `default = empty string`.
Actually this is a function that returns string. You can set it to return
current "RequestID" or "SessionID" - depends infra code is running in.


### Note

JSON formatter is good for log managements systems.

To make the log more "human readable" (for example in dev/stage environments)
another formatter should be used, see [human readable formatter](../log-formatters/logrus-human-formatter/)
for more info.
