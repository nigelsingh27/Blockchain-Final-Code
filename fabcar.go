/*
 * chaincode-fabcar-go
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright Institutionship.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

 package main

 /* Imports
  * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
  * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
  */
 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
	 "strconv"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
 
 // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
 type Certif struct {
	 Person   string `json:"person"`
	 Year  string `json:"year"`
	 Degree string `json:"degree"`
	 Institution  string `json:"institution"`
 }
 
 /*
  * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
  * Best practice is to have any Ledger initialization in separate function -- see initLedger()
  */
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 /*
  * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
  * The calling application program has also specified the particular smart contract function to be called, with arguments
  */
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 if function == "queryCar" {
		 return s.queryCar(APIstub, args)
	 } else if function == "initLedger" {
		 return s.initLedger(APIstub)
	 } else if function == "createCar" {
		 return s.createCar(APIstub, args)
	 } else if function == "queryAllCars" {
		 return s.queryAllCars(APIstub)
	 } else if function == "changeCertifInstitution" {
		 return s.changeCertifInstitution(APIstub, args)
	 } else if function == "changeCertifYear" {
		 return s.changeCertifYear(APIstub, args)
	 } else if function == "changeCertifDegree" {
		 return s.changeCertifDegree(APIstub, args)
	 } else if function == "changeCertifPerson" {
		 return s.changeCertifPerson(APIstub, args)
	 }
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 return shim.Success(carAsBytes)
 }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	 certifs := []Certif{
		 Certif{Person: "500684626", Year: "Year0", Degree: "Degree0", Institution: "Institution0"},
		 Certif{Person: "500684627", Year: "Year1", Degree: "Degree1", Institution: "Institution1"},
		 Certif{Person: "500684628", Year: "Year2", Degree: "Degree2", Institution: "Institution2"},
		 Certif{Person: "500684629", Year: "Year3", Degree: "Degree3", Institution: "Institution3"},
		 Certif{Person: "500684610", Year: "Year4", Degree: "Degree4", Institution: "Institution4"},
		 Certif{Person: "500684611", Year: "Year5", Degree: "Degree5", Institution: "Institution5"},
		 Certif{Person: "500684612", Year: "Year6", Degree: "Degree6", Institution: "Institution6"},
		 Certif{Person: "500684613", Year: "Year7", Degree: "Degree7", Institution: "Institution7"},
		 Certif{Person: "500684614", Year: "Year8", Degree: "Degree8", Institution: "Institution8"},
		 Certif{Person: "500684615", Year: "Year9", Degree: "Degree9", Institution: "Institution9"},
	 }
 
	 i := 0
	 for i < len(certifs) {
		 fmt.Println("i is ", i)
		 carAsBytes, _ := json.Marshal(certifs[i])
		 APIstub.PutState("PERSON"+strconv.Itoa(i), carAsBytes)
		 fmt.Println("Added", certifs[i])
		 i = i + 1
	 }
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 5 {
		 return shim.Error("Incorrect number of arguments. Expecting 5")
	 }
 
	 var certif = Certif{Person: args[1], Year: args[2], Degree: args[3], Institution: args[4]}
 
	 carAsBytes, _ := json.Marshal(certif)
	 APIstub.PutState(args[0], carAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 startKey := "PERSON0"
	 endKey := "PERSON999"
 
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
		 // Add a comma before array members, suppress it for the first array member
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
		 buffer.WriteString("{\"Key\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(queryResponse.Key)
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"Record\":")
		 // Record is a JSON object, so we write as-is
		 buffer.WriteString(string(queryResponse.Value))
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")
 
	 fmt.Printf("- queryAllCars:\n%s\n", buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }
 
 func (s *SmartContract) changeCertifInstitution(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 certif := Certif{}
 
	 json.Unmarshal(carAsBytes, &certif)
	 certif.Institution = args[1]
 
	 carAsBytes, _ = json.Marshal(certif)
	 APIstub.PutState(args[0], carAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) changeCertifYear(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 certif := Certif{}
 
	 json.Unmarshal(carAsBytes, &certif)
	 certif.Year = args[1]
 
	 carAsBytes, _ = json.Marshal(certif)
	 APIstub.PutState(args[0], carAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) changeCertifDegree(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 certif := Certif{}
 
	 json.Unmarshal(carAsBytes, &certif)
	 certif.Degree = args[1]
 
	 carAsBytes, _ = json.Marshal(certif)
	 APIstub.PutState(args[0], carAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) changeCertifPerson(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 certif := Certif{}
 
	 json.Unmarshal(carAsBytes, &certif)
	 certif.Person = args[1]
 
	 carAsBytes, _ = json.Marshal(certif)
	 APIstub.PutState(args[0], carAsBytes)
 
	 return shim.Success(nil)
 }
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
 