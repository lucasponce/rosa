package interactive

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOcm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interactive Suite")
}
