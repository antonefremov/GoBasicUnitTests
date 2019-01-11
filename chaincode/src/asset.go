package main

import (
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/protos/peer"
  "encoding/json"
  "fmt"
)

type Asset struct {
	ExternalId      string `json:"externalId"` // Asset Id
	Sernr						string `json:"sernr"`      // Serial Number
	Matnr						string `json:"matnr"`      // Material Number
	ObjDesc					string `json:"objDesc"`    // Description
  OwnerId         string `json:"ownerId"`    // Owner Id
}

// ============================================================================================================================
// Create Asset - create a new Asset and store it in the chaincode state
// ============================================================================================================================
func create_asset(stub shim.ChaincodeStubInterface, args []string) (peer.Response) {
	var err error

	if len(args) < 5 {
		return shim.Error("Incorrect number of arguments. Expecting at least 5")
	}

  var asset Asset
	asset.ExternalId = args[0]
	asset.Sernr = args[1]
	asset.Matnr = args[2]
  asset.ObjDesc = args[3]
  asset.OwnerId = args[4]

  owner, err := get_owner(stub, asset.OwnerId)
	if owner.Id == "" && err != nil {
		return shim.Error("Owner does not exist with Id '" + asset.OwnerId + "'")
	}

  assetAsBytes, _ := json.Marshal(asset)
	err = stub.PutState(asset.ExternalId, assetAsBytes)
	if err != nil {
    fmt.Println("Could not save an Asset")
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// Read Asset - get an Asset from the ledger
// ============================================================================================================================
func read_asset(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var jsonResp string

	key := args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ============================================================================================================================
// Change Owner for an Asset
// ============================================================================================================================
func set_owner(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	var asset_id = args[0]
	var new_owner_id = args[1]

	// check if user already exists
	owner, err := get_owner(stub, new_owner_id)
	if err != nil || len(owner.Username) == 0 {
		return shim.Error("This Owner does not exist - " + new_owner_id)
	}

	// get asset's current state
	assetAsBytes, err := stub.GetState(asset_id)
	if err != nil {
		return shim.Error("Failed to read Asset")
	}

	var asset Asset
	json.Unmarshal(assetAsBytes, &asset)

	asset.OwnerId = new_owner_id

	jsonAsBytes, _ := json.Marshal(asset)
	err = stub.PutState(asset_id, jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
