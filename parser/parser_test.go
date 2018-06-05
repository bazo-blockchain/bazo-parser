package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParser_ProgramAddNums(t *testing.T) {
	contract, err := ioutil.ReadFile("../contracts/addNums.sc")
	if err != nil {
		fmt.Print(err)
	}

	contractAsString := string(contract) // convert content to a 'string'

	instructionCode := Parse(contractAsString)
	for _, v := range instructionCode {
		fmt.Printf("%v, ", v)
	}

	if !bytes.Equal(instructionCode, []byte{0, 0, 8, 0, 0, 5, 4, 50}) {
		t.Errorf("Expected tos to be '0, 0, 8, 0, 0, 5, 4, 48' error message but was %v", instructionCode)
	}
}

func TestParser_ProgrammFunctionCall(t *testing.T) {
	contract, err := ioutil.ReadFile("../contracts/functionCall.sc")
	if err != nil {
		fmt.Print(err)
	}

	contractAsString := string(contract) // convert content to a 'string'

	instructionCode := Parse(contractAsString)
	for _, v := range instructionCode {
		fmt.Printf("%v, ", v)
	}

	if !bytes.Equal(instructionCode, []byte{0, 1, 217, 228, 0, 0, 5, 21, 0, 13, 2, 50, 28, 0, 28, 1, 4, 24}) {
		t.Errorf("Expected tos to be '0, 1, 217, 228, 0, 0, 5, 21, 0, 13, 2, 48, 27, 0, 27, 1, 4, 24' error message but was %v", instructionCode)
	}
}

/*
func TestParser_ProgrammTokenizationContract(t *testing.T) {
	contract, err := ioutil.ReadFile("../contracts/tokenizationContract.sc")
	if err != nil {
		fmt.Print(err)
	}

	contractAsString := string(contract) // convert content to a 'string'

	instructionCode := Parse(contractAsString)
	for _, v  := range instructionCode {
		fmt.Printf( "%v, ", v)
	}
}
*/
