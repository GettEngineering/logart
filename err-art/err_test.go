package errart_test

import (
	"testing"
	"time"

	"github.com/gtforge/logart/log-art"

	"github.com/gtforge/logart/err-art"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ErrorWithFields suit")
}

var _ = Describe("ErrorWithFields", func() {

	It("Error comparison", func() {
		constErr := errart.NewConstError("err")

		err := func() error {
			return constErr
		}()

		same := err == constErr
		Expect(same).To(BeTrue())
		Expect(err).To(Equal(constErr))
	})

	It(".Error", func() {
		err1 := errart.New("err_msg_1")
		err2 := errart.Wrap(err1, "err_msg_2").WithField("f", "v")
		Expect(err2.Error()).To(Equal("err_msg_2: err_msg_1 {\"f\":\"v\"}"))
	})

	It("Wrap nil error", func() {
		err2 := errart.Wrap(nil, "err_msg_2")
		Expect(err2).To(BeNil())
	})

	It("Basic usage", func() {

		// call function that might (and will) return error
		err := func() error {

			fieldsHolder := fieldsHolder{
				fields: map[string]interface{}{
					"field1": "val1",
					"field2": 2,
					"field3": true,
				},
			}

			// call another function that might (and will) returns error
			err := func() error {
				return errart.New("most_internal_err_msg")
			}()

			// regular error check
			if err != nil {
				return errart.Wrap(err, "something went wrong").
					WithFieldsFrom(fieldsHolder).
					WithField("field4", "val4").
					WithFields(errart.Fields{
						"field5": "val5",
					})
			}

			return nil // will never happened
		}()

		logart.WithError(err).Error("log_error")

		time.Sleep(100 * time.Millisecond)
	})

})

type fieldsHolder struct {
	fields map[string]interface{}
}

func (fp fieldsHolder) GetFields() map[string]interface{} {
	return fp.fields
}
