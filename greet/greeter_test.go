package greet_test

import (
	"bytes"
	"errors"

	"github.com/cholick/pipeline-test/greet"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type errorWriter struct{}

func (ew errorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("write failure!")
}

var _ = Describe("Greeter", func() {
	It("writes greeting", func() {
		b := &bytes.Buffer{}
		g := greet.Greeter{
			Message: "Hello",
		}

		err := g.Greet(b)

		Expect(err).To(BeNil())
		Expect(b.String()).To(Equal("Hello"))
	})

	It("returns error", func() {
		w := &errorWriter{}

		g := greet.Greeter{
			Message: "Hello",
		}
		err := g.Greet(w)

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("write failure!"))
	})
})
