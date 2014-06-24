package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSoldier(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Soldier Suite")
}
