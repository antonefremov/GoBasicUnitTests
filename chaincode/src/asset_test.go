package main

import (
  "github.com/hyperledger/fabric/core/chaincode/shim"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "encoding/json"
)

var stub = shim.NewMockStub("testingStub", new(SimpleChaincode))
var status200 = int32(200)
var mockAsset = GetAssetForTesting()
var mockAssetAsBytes = ConvertBytesToAssetAsBytes(mockAsset)
var mockOwner = GetFirstOwnerForTesting()
var anotherMockOwner = GetSecondOwnerForTesting()
var argsToCreateOwner = append([][]byte{[]byte("create_owner")}, mockOwner...)
var argsToCreateAnotherOwner = append([][]byte{[]byte("create_owner")}, anotherMockOwner...)
var argsToCreateAsset = append([][]byte{[]byte("create_asset")}, mockAsset...)
var argsToReadAsset = [][]byte{[]byte("read_asset"), mockAsset[0]}
var argsToPassAsset = [][]byte{[]byte("set_owner"), mockAsset[0], anotherMockOwner[0]}
var payload = []byte{}

var _ = Describe("Tests for Assets", func() {

  BeforeSuite(func() {
    stub.MockInit("000", nil)
  })

  Describe("Checking the CRUD operations", func() {
    Context("Checking that create/read work fine", func() {
      It("An Owner should be created successfully first", func() {
        receivedStatus := stub.MockInvoke("000", argsToCreateOwner).Status
        Expect(receivedStatus).Should(Equal(status200))
      })
      It("First Asset instance should be saved successfully", func() {
        receivedStatus := stub.MockInvoke("000", argsToCreateAsset).Status
        Expect(receivedStatus).Should(Equal(status200))
      })
      It("Another Asset instance is retrieved by the same Id", func() {
        result := stub.MockInvoke("000", argsToReadAsset)
        payload = []byte(result.Payload)
        Expect(result.Status).Should(Equal(status200))
      })
      Specify("Asset instances are identical", func() {
        Expect(payload).To(Equal(mockAssetAsBytes))
      })
    })
  })

  Describe("Running tests to check Assets tracking", func() {
    Context("Checking that Asset is transferred correctly between Owners", func() {
      It("First Owner should be created successfully", func() {
        stub.DelState(string(mockOwner[0]))
        receivedStatus := stub.MockInvoke("000", argsToCreateOwner).Status
        Expect(receivedStatus).Should(Equal(status200))
      })
      It("Second Owner should be created successfully", func() {
        receivedStatus := stub.MockInvoke("000", argsToCreateAnotherOwner).Status
        Expect(receivedStatus).Should(Equal(status200))
      })
      It("Asset should be created successfully belonging to the first Owner", func() {
        receivedStatus := stub.MockInvoke("000", argsToCreateAsset).Status
        Expect(receivedStatus).Should(Equal(status200))
      })
      It("Asset should be passed to the second Owner", func() {
        result := stub.MockInvoke("000", argsToPassAsset)
        Expect(result.Status).Should(Equal(status200))
      })
      Specify("Check that Asset belongs to the second Owner", func() {
        result := stub.MockInvoke("000", argsToReadAsset)
        payload = []byte(result.Payload)
        var asset Asset
        json.Unmarshal(payload, &asset)
        Expect(asset.OwnerId).Should(Equal(string(anotherMockOwner[0])))
      })
    })
  })
})
