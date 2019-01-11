package main

import (
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/protos/peer"
  "fmt"
  "errors"
  "encoding/json"
)

type Owner struct {
	Id         string `json:"id"`        // Owner Id
	Username   string `json:"username"`  // User Name
	Company    string `json:"company"`   // User Company name
}

// ============================================================================================================================
// Create Owner - create a new owner and store it in the chaincode state
// ============================================================================================================================
func create_owner(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting at least 3")
	}

	var owner Owner
	owner.Id = args[0]
	owner.Username = args[1]
	owner.Company = args[2]

	// check if Owner already exists
	_, err = get_owner(stub, owner.Id)
	if err == nil {
		return shim.Error("This Owner already exists - " + owner.Id)
	}

	// save the Owner
	ownerAsBytes, _ := json.Marshal(owner)         // convert to array of bytes
	err = stub.PutState(owner.Id, ownerAsBytes)    // store Owner by its Id
	if err != nil {
		fmt.Println("Could not save an Owner")
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// Read Owner - get an Owner record as bytes from the ledger
// ============================================================================================================================
func read_owner(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error
  key := args[0]

	valAsbytes, err := stub.GetState(key)
	if err != nil {
    return shim.Error("Owner does not exist with Id '" + key + "'")
	}

	return shim.Success(valAsbytes)
}

// ============================================================================================================================
// Get Owner - get an Owner instance from ledger
// ============================================================================================================================
func get_owner(stub shim.ChaincodeStubInterface, id string) (Owner, error) {
	var owner Owner

  ownerAsBytes, err := stub.GetState(id)
  if err != nil {
    return owner, errors.New("Failed to find Owner with Id '" + id + "'")
  }

	json.Unmarshal(ownerAsBytes, &owner)

	if len(owner.Username) == 0 {  // test if Owner is actually here or just nil
		return owner, errors.New("Owner does not exist with Id '" + id + "'")
	}

	return owner, nil
}
