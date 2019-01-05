package main

import (
    "testing"
    "fmt"
    "bytes"
    "encoding/json"
    "strings"
)

func TestOwnerCreateRead(test *testing.T) {

  // prepare all the necessary objects and keys
  stub := InitChaincode(test)
  ownerForTesting := GetFirstOwnerForTesting()
  ownerForTestingAsBytes := ConvertBytesToOwnerAsBytes(ownerForTesting)
  ownerForTestingKey := [][]byte{ownerForTesting[0]}

  // invoke the functions
                  Invoke(test, stub, "create_owner", ownerForTesting)
  ownerAsBytes := Invoke(test, stub, "read_owner", ownerForTestingKey)

  // check the results
  if bytes.Compare(ownerForTestingAsBytes, ownerAsBytes) != 0 {
    fmt.Println("\n>>> FAILED TEST: read_owner.\n", "\nExpected:\n", string(ownerForTestingAsBytes), "\nActual:\n", string(ownerAsBytes), "\n ")
    test.FailNow()
  }
}

func TestAssetCreateRead(test *testing.T) {

  // prepare all the necessary objects and keys
  stub := InitChaincode(test)
  ownerForTesting := GetFirstOwnerForTesting()
  assetForTesting := GetAssetForTesting()
  assetForTestingAsBytes := ConvertBytesToAssetAsBytes(assetForTesting)
  assetForTestingKey := [][]byte{assetForTesting[0]}

  // invoke the functions
                  Invoke(test, stub, "create_owner", ownerForTesting)
                  Invoke(test, stub, "create_asset", assetForTesting)
  assetAsBytes := Invoke(test, stub, "read_asset", assetForTestingKey)

  // check the results
  if bytes.Compare(assetForTestingAsBytes, assetAsBytes) != 0 {
    fmt.Println("\n>>> FAILED TEST: read_asset.\n", "\nExpected:\n", string(assetForTestingAsBytes), "\nActual:\n", string(assetAsBytes), "\n ")
    test.FailNow()
  }
}

func TestAssetTransfer(test *testing.T) {
  // prepare all the necessary objects and keys
  stub := InitChaincode(test)

  firstOwnerForTesting := GetFirstOwnerForTesting()
  secondOwnerForTesting := GetSecondOwnerForTesting()
  assetForTesting := GetAssetForTesting()

  args := [][]byte{assetForTesting[0], secondOwnerForTesting[0]}

  // invoke the functions
                  Invoke(test, stub, "create_owner", firstOwnerForTesting)
                  Invoke(test, stub, "create_asset", assetForTesting)
                  Invoke(test, stub, "create_owner", secondOwnerForTesting)
                  Invoke(test, stub, "set_owner", args)
  assetAsBytes := Invoke(test, stub, "read_asset", [][]byte{assetForTesting[0]})

  // check the results
  var asset Asset
  json.Unmarshal(assetAsBytes, &asset)
  if strings.Compare(asset.OwnerId, string(secondOwnerForTesting[0])) != 0 {
    fmt.Println("\n>>> FAILED TEST: set_owner.\n", "\nExpected:\n", asset.OwnerId, "\nActual:\n", string(secondOwnerForTesting[0]), "\n ")
    test.FailNow()
  }
}
