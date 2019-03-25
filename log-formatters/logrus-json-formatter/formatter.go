package logrusjsonformatter

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type GettLogFormatter struct{}

func (f *GettLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(map[string]interface{}, len(entry.Data)+4)
	data["level"] = entry.Level.String()
	data["time"] = entry.Time.UTC().Format(options.TimestampFormat)
	data["request_id"] = options.LogIDProvider()

	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	data["message"] = entry.Message

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON. Data: %v, error: %v", data, err)
	}
	return append(serialized, '\n'), nil
}
