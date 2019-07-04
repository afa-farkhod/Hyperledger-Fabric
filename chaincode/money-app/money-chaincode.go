// SPDX-License-Identifier: Apache-2.0
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Money struct {
	Bank  string `json:"bank"`
	Date  string `json:"date"`
	Serial string `json:"serial"`
	Holder  string `json:"holder"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryMoney" {
		return s.queryMoney(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordMoney" {
		return s.recordMoney(APIstub, args)
	} else if function == "queryAllMoney" {
		return s.queryAllMoney(APIstub)
	} else if function == "changeMoneyHolder" {
		return s.changeMoneyHolder(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}


func (s *SmartContract) queryMoney(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	moneyAsBytes, _ := APIstub.GetState(args[0])
	if moneyAsBytes == nil {
		return shim.Error("Could not locate money")
	}
	return shim.Success(moneyAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	money := []Money{
		Money{Bank: "AsakaBank", Serial: "1235AF545", Date: "17052018", Holder: "Mr.Farkhod1"},
		Money{Bank: "NationalBank", Serial: "254845AS", Date: "15062019", Holder: "Mr.Farkhod2"},
		Money{Bank: "WorldBank", Serial: "545A54515", Date: "02042018", Holder: "Mr.Farkhod3"},
		Money{Bank: "KhalkBank", Serial: "TF5541564", Date: "03062019", Holder: "Mr.Farkhod4"},
		Money{Bank: "Infinbank", Serial: "54545F515", Date: "12062019", Holder: "Mr.Farkhod5"},
		Money{Bank: "TrustBank", Serial: "5487813RJ45", Date: "21062018", Holder: "Mr.Farkhod6"},
		Money{Bank: "TuronBank", Serial: "F1564875468", Date: "22062019", Holder: "Mr.Farkhod7"},
		Money{Bank: "HamkorBank", Serial: "784A54545A", Date: "25072019", Holder: "Mr.Farkhod8"},
		Money{Bank: "CapitalBank", Serial: "1213448T45", Date: "30042019", Holder: "Mr.Farkhod9"},
		Money{Bank: "AlokaBank", Serial: "54512854AS54", Date: "02052019", Holder: "Mr.Farkhod10"},
	}

	i := 0
	for i < len(money) {
		fmt.Println("i is ", i)
		moneyAsBytes, _ := json.Marshal(money[i])
		APIstub.PutState(strconv.Itoa(i+1), moneyAsBytes)
		fmt.Println("Added", money[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) recordMoney(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var money = Money{ Bank: args[1], Serial: args[2], Date: args[3], Holder: args[4] }

	moneyAsBytes, _ := json.Marshal(money)
	err := APIstub.PutState(args[0], moneyAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record money packet: %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryAllMoney(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllMoney:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeMoneyHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	moneyAsBytes, _ := APIstub.GetState(args[0])
	if moneyAsBytes == nil {
		return shim.Error("Could not locate money")
	}
	money := Money{}

	json.Unmarshal(moneyAsBytes, &money)
	
	money.Holder = args[1]

	moneyAsBytes, _ = json.Marshal(money)
	err := APIstub.PutState(args[0], moneyAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change money holder: %s", args[0]))
	}

	return shim.Success(nil)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
