package heartbeat_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGRPC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Heartbeat Service Suite")
}
