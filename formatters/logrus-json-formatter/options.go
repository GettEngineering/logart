package logrusjsonformatter

import "time"

func Set() {
	options = DefaultFormatOptions
}

func SetCustomized(o FormatOptions) {
	options = o
}

var options = DefaultFormatOptions

type FormatOptions struct {
	TimestampFormat string
	LogIDProvider   func() string
}

var DefaultFormatOptions = FormatOptions{
	TimestampFormat: time.RFC3339,
	LogIDProvider:   func() string { return "" },
}
