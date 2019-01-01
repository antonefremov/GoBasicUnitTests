package main

import (
    "testing"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

func initChaincode(test *testing.T) *shim.MockStub {
    stub := shim.NewMockStub("testingStub", new(SimpleChaincode))
    result := stub.MockInit("000", nil)

    if result.Status != shim.OK {
        test.FailNow()
    }
    return stub
}

func TestInstancesCreation(test *testing.T) {
  stub := initChaincode(test)

  assetExternalId := "ID01"
  ownerId := "o1"
  Invoke(test, stub, "create_owner", ownerId, "Username_1", "Company_1")
  Invoke(test, stub, "create_asset", assetExternalId, "Sernr1234", "Matnr1234", "ObjDesc", ownerId)
}
