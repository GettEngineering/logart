package errart

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Wrap(err error, msg string) ErrArt {

	if err == nil {
		return nil // pay attention, can't use WithField/s if err == nil
	}

	errArt := errArt{
		fields: Fields{},
	}

	// copy fields and error from input error
	if e, ok := err.(ErrorWithFields); ok {
		err = e.GetError()
		errArt = errArt.addFields(e.GetFields())
	}

	errArt.err = errors.Wrap(err, msg)
	return errArt
}

func NewConstError(msg string) error {
	return errors.New(msg)
}

func New(msg string) ErrArt {
	return errArt{
		fields: Fields{},
		err:    errors.New(msg),
	}
}

// container to pass error and fields as error object
type errArt struct {
	fields Fields
	err    error
}

func (r errArt) WithFieldsFrom(holder FieldsHolder) ErrArt {
	for k, v := range holder.GetFields() {
		r.fields[k] = v
	}
	return r
}

func (r errArt) WithFields(fields Fields) ErrArt {
	for k, v := range fields {
		r.fields[k] = v
	}
	return r
}

func (r errArt) addFields(fields Fields) errArt {
	for k, v := range fields {
		r.fields[k] = v
	}
	return r
}

func (r errArt) WithField(key string, val interface{}) ErrArt {
	r.fields[key] = fmt.Sprintf("%+v", val) // pretty print complex objects
	return r
}

func (r errArt) GetError() error {
	return r.err
}

func (r errArt) GetFields() map[string]interface{} {
	return Fields(r.fields)
}

// Implements error interface, so can be returned as error
func (r errArt) Error() string {
	return fmt.Sprint(r.err, " ", map2String(r.fields))
}

func map2String(m map[string]interface{}) string {
	// json.Marshal impl without using json.Marshal
	var pairs []string
	for k, v := range m {
		pairs = append(pairs, fmt.Sprintf("\"%v\":\"%+v\"", k, v))
	}
	return fmt.Sprintf("{%v}", strings.Join(pairs, ","))
}
