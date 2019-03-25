package logrusjsonformatter

import (
	"time"

	"github.com/sirupsen/logrus"
)

func Set() {
	logrus.SetFormatter(&JsonLogFormatter{
		formatOptions: DefaultFormatOptions,
	})
}

func SetCustomized(options FormatOptions) {
	logrus.SetFormatter(&JsonLogFormatter{
		formatOptions: options,
	})
}

type FormatOptions struct {
	TimestampFormat string
	LogIDProvider   func() string
}

var DefaultFormatOptions = FormatOptions{
	TimestampFormat: time.RFC3339,
	LogIDProvider:   func() string { return "" },
}
