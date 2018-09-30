package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	peer "github.com/hyperledger/fabric/protos/peer"
)

type chaincode struct {
}
type aid struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

func (c *chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (c *chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "read" {
		return c.read(stub, args)
	}
	return shim.Success(nil)
}

func (c *chaincode) read(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Wrong number of arguments....expected 1")
	}
	aidasbytes, _ := stub.GetState(args[0])
	return shim.Success(aidasbytes)
}

func (c *chaincode) write(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	name := args[0]
	quantity := args[1]
	provider := args[2]

	Aid := aid{name, quantity}
	aidasbytes, _ := json.Marshal(Aid)
	err := stub.PutState(provider, aidasbytes)
	if err != nil {
		return shim.Error("Write failed")
	}
	return shim.Success(nil)
}

func main() {

	err := shim.Start(new(chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode......")
	}
}
