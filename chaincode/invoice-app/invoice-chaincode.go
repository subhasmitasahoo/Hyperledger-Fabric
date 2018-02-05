package main
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

type Invoice struct {
	SGSTN string `json:"sgstn"`
	SState string `json:"sstate"`
	InvoiceNo string `json:"invoiceno"`
	Date string `json:"date"`
	CGSTN string `json:"cgstn"`
	CName string `json:"cname"`
	BillAdd string `json:"billadd"`
	ShipAdd string `json:"shipadd"`
	TaxableAmount string `json:"taxableamount"`
	TotalTax string `json:"totalTax"`
	InvoiceTotal string `json:"invoicetotal"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "getInvoice" {
		return s.getInvoice(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createInvoice" {
		return s.createInvoice(APIstub, args)
	} else if function == "getAllInvoices" {
		return s.getAllInvoices(APIstub)
	}
	return shim.Error("Invalid Smart Contract function name...... "+function)
}

func (s *SmartContract) getInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	invoiceAsBytes, _ := APIstub.GetState(args[0])
	if invoiceAsBytes == nil {
		return shim.Error("Could not find the invoice")
	}
	return shim.Success(invoiceAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	invoices := []Invoice{
		Invoice{SGSTN: "sgstn1", SState: "Odisha", InvoiceNo: "invoice1", Date: "04/02/17",CGSTN:"cgstn1",CName:"custom1",BillAdd:"BillAdd1",ShipAdd:"ShipAdd1",TaxableAmount:"100.00",TotalTax:"18.00",InvoiceTotal:"118.00"},
    Invoice{SGSTN: "sgstn2", SState: "Karnataka", InvoiceNo: "invoice2", Date: "04/02/17",CGSTN:"cgstn2",CName:"custom2",BillAdd:"BillAdd2",ShipAdd:"ShipAdd2",TaxableAmount:"200.00",TotalTax:"36.00",InvoiceTotal:"236.00"},
    Invoice{SGSTN: "sgstn3", SState: "Bihar", InvoiceNo: "invoice3", Date: "04/02/17",CGSTN:"cgstn3",CName:"custom3",BillAdd:"BillAdd3",ShipAdd:"ShipAdd3",TaxableAmount:"100.00",TotalTax:"18.00",InvoiceTotal:"118.00"},
    Invoice{SGSTN: "sgstn4", SState: "UP", InvoiceNo: "invoice4", Date: "04/02/17",CGSTN:"cgstn4",CName:"custom4",BillAdd:"BillAdd4",ShipAdd:"ShipAdd4",TaxableAmount:"200.00",TotalTax:"36.00",InvoiceTotal:"236.00"},
    Invoice{SGSTN: "sgstn5", SState: "Kashmir", InvoiceNo: "invoice5", Date: "04/02/17",CGSTN:"cgstn5",CName:"custom5",BillAdd:"BillAdd5",ShipAdd:"ShipAdd5",TaxableAmount:"100.00",TotalTax:"18.00",InvoiceTotal:"118.00"},
	}

	i := 0
	for i < len(invoices) {
		fmt.Println("i is ", i)
		invoiceAsBytes, _ := json.Marshal(invoices[i])
		APIstub.PutState(strconv.Itoa(i+1), invoiceAsBytes)
		fmt.Println("Added", invoices[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 12 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var invoice = Invoice{ SGSTN: args[1], SState: args[2], InvoiceNo: args[3], Date: args[4], CGSTN: args[5], CName: args[6], BillAdd: args[7],ShipAdd: args[8],TaxableAmount: args[9],TotalTax: args[10],InvoiceTotal: args[11] }

	invoiceAsBytes, _ := json.Marshal(invoice)
	err := APIstub.PutState(args[0], invoiceAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create invoice: %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) getAllInvoices(APIstub shim.ChaincodeStubInterface) sc.Response {

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
		// Add comma before array members,suppress it for the first array member
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

	fmt.Printf("- getAllInvoices:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * main function *
calls the Start function
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
