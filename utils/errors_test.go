package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"dt-logger/logger"
)

const (
	validErrMsg = "Example Error"
)

var _ = Describe("Errors", func() {

	Describe("When sending an error string", func() {
		Context("If error string is not empty", func() {
			It("error should be logged", func() {

				pwd, err := os.Getwd()
				if err != nil {
					fmt.Println(err)
				}

				logfile := pwd + "/mytest.log"
				logger.SetupFileOutput("Error Test", logfile, true)

				err = HandleError(validErrMsg)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(validErrMsg))

				time.Sleep(2)
				logContent, err := ioutil.ReadFile(logfile)
				if err != nil {
					fmt.Println(err)
				}

				Expect(string(logContent)).Should(ContainSubstring(validErrMsg))
				os.Remove(logfile)
			})
		})
	})

	Describe("When sending an error string", func() {
		Context("If error string is not empty", func() {
			It("error should be returned", func() {

				err := HandleError(validErrMsg)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(validErrMsg))

			})
		})
	})
})
