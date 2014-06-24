package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Soldier", func() {
	var soldier string
	var appDir string

	BeforeEach(func() {
		var err error

		appDir, err = ioutil.TempDir("", "the-soldier-app-dir")
		Ω(err).ShouldNot(HaveOccurred())

		soldier, err = gexec.Build("github.com/cloudfoundry-incubator/windows-circus/soldier")
		Ω(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(appDir)
	})

	It("runs the given command", func() {
		session, err := gexec.Start(
			exec.Command(soldier, appDir, "ruby", "-v"),
			GinkgoWriter,
			GinkgoWriter,
		)
		Ω(err).ShouldNot(HaveOccurred())

		Eventually(session, 5).Should(gbytes.Say("ruby"))
	})

	It("handles arguments containing spaces (if the argument is quoted)", func() {
		session, err := gexec.Start(
			exec.Command(soldier, appDir, "ruby", "-e", "'puts true'"),
			GinkgoWriter,
			GinkgoWriter,
		)
		Ω(err).ShouldNot(HaveOccurred())

		Eventually(session, 5).Should(gbytes.Say("true"))
	})

	It("changes the working directory based on the first argument", func() {
		session, err := gexec.Start(
			exec.Command(soldier, appDir, "echo $pwd"),
			GinkgoWriter,
			GinkgoWriter,
		)
		Ω(err).ShouldNot(HaveOccurred())

		Eventually(session, 5).Should(gbytes.Say("the-soldier-app-dir"))
	})
})
