package main

import (
	"flag"
	"fmt"
	"github.com/bazo-blockchain/bazo-parser/parser"
	"io/ioutil"
)

func main() {
	var filePath string
	flag.String(filePath, "path", "Define path to your contract")
	flag.Parse()

	fmt.Println("Define the path to your contract")
	if _, err := fmt.Scanf("%v", &filePath); err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	contract, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}

	contractAsString := string(contract) // convert content to a 'string'

	instructionCode := parser.Parse(contractAsString)
	for _, v := range instructionCode {
		fmt.Printf("%v, ", v)
	}
}
