# logrus formatter


### Usage:

Default setup:
```
logformatter.Set()
```

If you prefer manage the order of printed fields, you can initialize the formatter with this:
```
logformatter.SetWithFieldsOrder(
    logformatter.OrderedFields{"env", "order_id"},
    logformatter.OrderedFields{"error"},
)
```

First param defines the fields that would be printed on the beginning of the log's fields.
Second param defines the fields that would be printed on the ending of the log's fields.
Fields that are not defined, would be printed between these (alphabetically sorted).


Formatter allows you to change some of its configurations:

1. Log time format. By default it's only time, you can add also a date.
2. Length of request ID. By default it's 6 chars.
3. Log level length. By default it's 5 chars ("DEBUG"), you can put less, e.g.: 3 to get "DEB"
4. Color scheme of the log ([available colors](https://github.com/artiomgiza/go-color-256/blob/master/colors.png))

This could be done by setting `logformatter.DefaultOptions.` before colling ".Set/.SetWith..."

```
// format:
    TimeLayout              = "15:04:05.000"
    RequestIdLength         = 6
    levelLen                = 5

// colors:
    LogHeaderColor:          254 // almost white
    LogHeaderErrorColor:     88  // dark red
    LogMessageColor:         34  // sort of green
    FieldKeyColor:           62  // dark purple
    FieldValColor:           244 // dark gray
    FieldsJsonFrameValColor: 244 // dark gray
```

### Example:

With the formatter log could look like:

![alt text](https://github.com/gtforge/rex_common/blob/master/log_formatter/readme_files/example.png "Example")
