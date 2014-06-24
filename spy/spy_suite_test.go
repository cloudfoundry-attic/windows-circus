package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWindowsCircusSpy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Windows-Circus-Spy Suite")
}
