package main

import (
  "github.com/hyperledger/fabric/core/chaincode/shim"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Owner Tests", func() {

  stub := shim.NewMockStub("testingStub", new(SimpleChaincode))
  status200 := int32(200)
  mockOwner := GetFirstOwnerForTesting()
  mockOwnerAsBytes := ConvertBytesToOwnerAsBytes(mockOwner)
  argsToCreate := append([][]byte{[]byte("create_owner")}, mockOwner...)
  argsToRead := [][]byte{[]byte("read_owner"), mockOwner[0]}
  payload := []byte{}

  BeforeEach(func() {
    stub.MockInit("000", nil)
  })

  Describe("Running tests for the Owner", func() {
      Context("Checking that create/read for Owner work fine", func() {
          It("Should be created successfully", func() {
              receivedStatus := stub.MockInvoke("000", argsToCreate).Status
              Expect(receivedStatus).Should(Equal(status200))
          })
          It("Another Owner instance is retrieved successfully by the same Id", func() {
              result := stub.MockInvoke("000", argsToRead)
              payload = []byte(result.Payload)
              Expect(result.Status).Should(Equal(status200))
          })
          Specify("Owner instances are identical", func() {
              Expect(payload).To(Equal(mockOwnerAsBytes))
          })
      })
  })
})
