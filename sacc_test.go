package sacc_test

import (
	// import the chaincode
	. "sacc"

	// import shim
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// import testing packages
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sacc", func() {
	// create a new MockStub
	stub := shim.NewMockStub("uuid", new(SimpleAsset))

	// define some variables
	textValue1 := string("100")
	textValue2 := string("110")
	textValue3 := string("50")
	args := [][]byte{[]byte("A"), []byte("100")}

	// init some values
	BeforeSuite(func() {
		stub.MockInit("000", args)
	})

	// create the context
	Describe("Checking Saac chaincode operations", func() {
		Context("Checking the value of an Asset", func() {
			// create the first test case
			It("It should check Asset A with the intial value of 100", func() {
				// query the new value of asset A
				result, _ := stub.GetState("A")

				// compare the result with our expectation
				r := string(result)
				Expect(r).Should(Equal(textValue1))
			})

			// create the second test case
			It("It should check Asset A with an updated value of 50", func() {
				// set the new value for A
				args := [][]byte{[]byte("set"), []byte("A"), []byte("110")}

				// call the set function of the chaincode
				stub.MockInvoke("1", args)

				// query the new value of asset A
				result, _ := stub.GetState("A")

				// compare the result with our expectation
				r := string(result)
				Expect(r).Should(Equal(textValue2))
			})

			// create the third test case
			It("It should check the value of the new asset B with a value of 50", func() {
				// set the new value for B
				args := [][]byte{[]byte("set"), []byte("B"), []byte("50")}

				// call the set function of the chaincode
				stub.MockInvoke("set", args)

				// query the new value of asset B
				result, _ := stub.GetState("B")

				// compare the result with our expectation
				r := string(result)
				Expect(r).Should(Equal(textValue3))
			})
		})
	})
})
