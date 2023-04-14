package conversions_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConversions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Conversions Suite")
}
