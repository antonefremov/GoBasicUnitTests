package main

import (
    "testing"
)

func TestInstancesCreation(test *testing.T) {
  stub := InitChaincode(test)

  assetExternalId := "ID01"
  ownerId := "o1"
  Invoke(test, stub, "create_owner", ownerId, "Username_1", "Company_1")
  Invoke(test, stub, "create_asset", assetExternalId, "Sernr1234", "Matnr1234", "ObjDesc", ownerId)
}
