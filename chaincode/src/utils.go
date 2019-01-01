package main

import (
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "testing"
  "fmt"
  "strings"
)

// ============================================================================================================================
// Invoke wrapper
// ============================================================================================================================
func Invoke(test *testing.T, stub *shim.MockStub, function string, args ...string) {

    cc_args := make([][]byte, 1+len(args))
    cc_args[0] = []byte(function)
    for i, arg := range args {
        cc_args[i + 1] = []byte(arg)
    }
    result := stub.MockInvoke("000", cc_args)
    fmt.Println("Call:    ", function, "(", strings.Join(args,","), ")")
    fmt.Println("RetCode: ", result.Status)
    fmt.Println("RetMsg:  ", result.Message)
    fmt.Println("Payload: ", string(result.Payload))

    if result.Status != shim.OK {
        test.FailNow()
    }
}
