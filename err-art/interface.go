package errart

type ErrArt interface {
	WithField(key string, val interface{}) ErrArt
	WithFields(fields Fields) ErrArt
	WithFieldsFrom(holder FieldsHolder) ErrArt
	Error() string
}

type ErrorWithFields interface {
	GetError() error
	GetFields() map[string]interface{}
}

type FieldsHolder interface {
	GetFields() map[string]interface{}
}

type Fields map[string]interface{}
