package wifi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWifi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wifi Suite")
}
