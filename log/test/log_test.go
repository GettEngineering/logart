package logtest

import (
	"fmt"
	"github.com/gtforge/logart/formatters/logrus-formatter"
	"strings"

	"github.com/golang/mock/gomock"
	"github.com/gtforge/logart/log/mock"

	"testing"
	"time"

	"github.com/gtforge/logart/log"
	"github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Log suit")
}

func Set4Gett() {
	o := logrusformatter.DefaultFormatOptions

	o.FormatterEnabled = func() bool {
		return true
	}

	o.FirstEverPrintedFields = logrusformatter.OrderedFields{"field2"}
	o.LastEverPrintedFields = logrusformatter.OrderedFields{"error"}
	o.LogIDProvider = func() string {
		return "xxxxxx"
	}

	co := logrusformatter.DefaultColorOptions
	co.OverrideLogColor = func(m string) (bool, int) {
		return strings.HasPrefix(m, "curl "), 234 // should be almost black
	}
	//co.LevelInfoColor = 222
	//co.ColorsEnabled = func() bool { return false }

	logrusformatter.SetCustomized(o, co)
}

var _ = Describe("OverrideLogColor usages", func() {

	BeforeEach(func() {
		logrus.SetLevel(logrus.DebugLevel)

		Set4Gett()
		//logrus.SetFormatter(&gettOps.GettLogFormatter{})

		o := logart.DefaultLogOptions
		o.PrintAsFields = []string{"error", "field1"}
		o.PrintFieldsInsideMessage = true
		o.PrintErrorWithStackTrace = false
		logart.SetLogOptions(o)

	})

	Describe("Basic usage", func() {
		It("Basic", func() {

			log := logart.
				WithField("field1", "value1").
				WithField("field2", "value2")

			log.Info("1st info log")

			log.
				WithField("field3", "value3").
				Debug("2nd debug log")

			log.Debug("1st info log")

			log.
				WithField("field5", "value5").
				WithError(fmt.Errorf("jopa")).
				Error("3th error log")

			err := fmt.Errorf("test_error")
			log.WithError(err).Error("error_with_stack?")

			// OUTPUT Scrum --> all fields - as fields
			//13:59:30.187 381 INFO  1st info  log {field1: value1, field2: value2}
			//13:59:30.187 381 DEBUG 2nd debug log {field1: value1, field2: value2, field3: value3}
			//13:59:30.187 381 ERROR 4th error log {field1: value1, field2: value2, field3: value3, field5: value5, error: jopa}

			// OUTPUT Prod --> first json - part of the message, second - as fields
			//14:08:13.307 bc5 INFO  1st info  log{"field2": "value2"} {field1: value1}
			//14:08:13.307 bc5 DEBUG 2nd debug log{"field2": "value2", "field3": "value3"} {field1: value1}
			//14:08:13.308 bc5 ERROR 4th error log{"field2": "value2", "field3": "value3", "field5": "value5"} {field1: value1, error: jopa}

			time.Sleep(100 * time.Millisecond)
		})
	})

	Describe("Mock", func() {

		Context("when mock is on", func() {

			var (
				mockCtrl *gomock.Controller
				mockLog  *mock_logart.MockLogger
			)

			BeforeEach(func() {
				mockCtrl = gomock.NewController(GinkgoT())
				mockLog = mock_logart.NewMockLogger(mockCtrl)
				logart.MockLog(mockLog) // <---------------------- replace log with mock
			})

			AfterEach(func() {
				logart.UnMockLog() // <---------------------- replace log back to normal
				mockCtrl.Finish()
			})

			Context("when logging", func() {
				BeforeEach(func() {
					// each action - separate Expect
					// pay attention - mockLog returned from WithError/Field/s
					mockLog.EXPECT().WithField("k1", "v1").Return(mockLog)
					mockLog.EXPECT().WithField("k2", "v2").Return(mockLog)
					mockLog.EXPECT().Debug("deb")
				})

				JustBeforeEach(func() {
					logart.
						WithField("k1", "v1").
						WithField("k2", "v2").Debug("deb")
				})

				It("should satisfy mocks", func() {
					// nothing here
				})
			})
		})

		Context("when mock is off", func() {
			// in case tests run not randomly - this part will ensure
			// logart.UnMockLog() worked
			// (in case of random tests run - it's not necessary)

			It("should use default log", func() {
				// check we are back to regular (un-mocked) log
				logart.Warn("yey! log is back!")
			})
		})

		Context("when value was fat", func() {

			It("should use default log", func() {
				fatValue := `
						{
							"k1": "v1",
						    "k2": "v2"
						}
						`

				logart.WithField("fat", fatValue).Debug("yey! no fat!")
			})
		})

	})
})

