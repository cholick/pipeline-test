package greet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGreet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Greet Suite")
}
