package main

import (
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/protos/peer"
  "fmt"
)

// SimpleChaincode example Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode - %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
  return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	// Handle invoke functions
	if function == "init" {                   // initialise the chaincode state
		return t.Init(stub)
	} else if function == "create_asset" {    // create a new asset
		return create_asset(stub, args)
	} else if function == "read_asset" {   	  // read an asset
		return read_asset(stub, args)
	} else if function == "create_owner" {    // create an owner
    return create_owner(stub, args)
  } else if function == "read_owner" {   	  // read an owner
		return read_owner(stub, args)
	} else if function == "set_owner" {       // change owner of an asset
		return set_owner(stub, args)
	}

  // error out
	fmt.Println("Received unknown invoke function name - '" + function + "'")
	return shim.Error("Received unknown invoke function name - '" + function + "'")
}
