package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWifiRedirector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WifiRedirector Suite")
}
