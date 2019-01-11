package main

import (
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "testing"
  "fmt"
  "bytes"
  "encoding/json"
)

// ============================================================================================================================
// Init Mock Chaincode helper
// ============================================================================================================================
func InitChaincode(test *testing.T) *shim.MockStub {
    stub := shim.NewMockStub("testingStub", new(SimpleChaincode))
    result := stub.MockInit("000", nil)

    if result.Status != shim.OK {
        test.FailNow()
    }
    return stub
}

// ============================================================================================================================
// Invoke wrapper
// ============================================================================================================================
func Invoke(test *testing.T, stub *shim.MockStub, function string, args [][]byte) []byte {
    const transactionId = "000"

    // prepend the function name as the first item
    args = append([][]byte{[]byte(function)}, args...)

    // prepare the parameters for printing
    byteDivider := []byte{','}
    byteArrayToPrint := bytes.Join(args[1:], byteDivider)

    // print information just before the call
    fmt.Println("Call:    ", function, "(", string(byteArrayToPrint), ")")

    // perform the MockInvoke call
    result := stub.MockInvoke(transactionId, args)

    // print the Invoke results
    fmt.Println("RetCode: ", result.Status)
    fmt.Println("RetMsg:  ", result.Message)
    fmt.Println("Payload: ", string(result.Payload))

    if result.Status != shim.OK {
      fmt.Println("Invoke", function, "failed", string(result.Message))
      return nil
    }

    return []byte(result.Payload)
}

// ============================================================================================================================
// Get a mock Owner
// ============================================================================================================================
func GetFirstOwnerForTesting() [][]byte {
  return [][]byte{
		[]byte("o1"),         // Id
		[]byte("Username_1"), // Username
		[]byte("Company_1")}  // Company
}

// ============================================================================================================================
// Get another mock Owner
// ============================================================================================================================
func GetSecondOwnerForTesting() [][]byte {
  return [][]byte{
		[]byte("o2"),         // Id
		[]byte("Username_2"), // Username
		[]byte("Company_2")}  // Company
}

// ============================================================================================================================
// Get a mock Asset
// ============================================================================================================================
func GetAssetForTesting() [][]byte {
  owner := GetFirstOwnerForTesting()

  return [][]byte{
		[]byte("a1"),         // ExternalId
		[]byte("1234"),       // Sernr
		[]byte("4321"),       // Matnr
    []byte("Desc_1"),     // ObjDesc
    owner[0]}             // OwnerId
}

// ============================================================================================================================
// Convert the Owner passed in as bytes to an Owner instance presented as bytes
// ============================================================================================================================
func ConvertBytesToOwnerAsBytes(ownerAsBytes [][]byte) []byte {
  var owner Owner
	owner.Id = string(ownerAsBytes[0])
	owner.Username = string(ownerAsBytes[1])
	owner.Company = string(ownerAsBytes[2])
	bagJSON, err := json.Marshal(owner)
	if err != nil {
		fmt.Println("Error converting an Owner record to JSON")
		return nil
	}
	return []byte(bagJSON)
}

// ============================================================================================================================
// Convert the Asset passed in as bytes to an Asset instance presented as bytes
// ============================================================================================================================
func ConvertBytesToAssetAsBytes(assetAsBytes [][]byte) []byte {
  var asset Asset
	asset.ExternalId = string(assetAsBytes[0])
	asset.Sernr = string(assetAsBytes[1])
	asset.Matnr = string(assetAsBytes[2])
  asset.ObjDesc = string(assetAsBytes[3])
  asset.OwnerId = string(assetAsBytes[4])

	bagJSON, err := json.Marshal(asset)
	if err != nil {
		fmt.Println("Error converting an Asset record to JSON")
		return nil
	}
	return []byte(bagJSON)
}
