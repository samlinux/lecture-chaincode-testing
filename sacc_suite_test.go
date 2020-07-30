package sacc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSacc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sacc Suite")
}
