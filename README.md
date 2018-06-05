# bazo-parser
Minimalistic parser, basis for building a compiler 

## Usage
### Prerequisites
The programming language Go (developed and tested with version >= 1.9) must be installed, the properties $GOROOT and $GOPATH must be set. For more information, please check out the official documentation.

### Getting started
1. Download the bazo-parser application.
```
go get github.com/bazo-blockchain/bazo-parser
```
2. Build the application.
```
$GOPATH/src/github.com/bazo-blockchain/bazo-parser
go build
```
3. Run the application.
```
./bazo-parser
```
4. Define the path to your smart contract.
```
Define the path to your contract
./contracts/addNums.sc
```
After hitting enter the parser prints the compiled byte code instructions, ready to copy to the VM or Miner
