package parser

const (
	OPCODE = iota
	BYTES
	BYTE
	LABEL
	ADDR
)

const (
	PUSH = iota
	DUP
	ROLL
	POP
	ADD
	SUB
	MULT
	DIV
	MOD
	NEG
	EQ
	NEQ
	LT
	GT
	LTE
	GTE
	SHIFTL
	SHIFTR
	NOP
	JMP
	JMPIF
	CALL
	CALLIF
	CALLEXT
	RET
	SIZE
	STORE
	SSTORE
	LOAD
	SLOAD
	ADDRESS // Address of account
	ISSUER  // Issuer of account
	BALANCE // Balance of account
	CALLER
	CALLVAL  // Amount of bazo coins transacted in transaction
	CALLDATA // Parameters and function signature hash
	NEWMAP
	MAPHASKEY
	MAPPUSH
	MAPGETVAL
	MAPSETVAL
	MAPREMOVE
	NEWARR
	ARRAPPEND
	ARRINSERT
	ARRREMOVE
	ARRAT
	SHA3
	CHECKSIG
	ERRHALT
	HALT
)

type Token struct {
	tokenType int
	value     string
}

type OpCode struct {
	Name     string
	Nargs    int
	ArgTypes []int
}

// Must be the same as in the vm
var OpCodes = map[int]OpCode{
	PUSH:      {"push", 1, []int{BYTES}},
	DUP:       {"dup", 0, nil},
	ROLL:      {"roll", 1, []int{BYTE}},
	POP:       {"pop", 0, nil},
	ADD:       {"add", 0, nil},
	SUB:       {"sub", 0, nil},
	MULT:      {"mult", 0, nil},
	DIV:       {"div", 0, nil},
	MOD:       {"mod", 0, nil},
	NEG:       {"neg", 0, nil},
	EQ:        {"eq", 0, nil},
	NEQ:       {"neq", 0, nil},
	LT:        {"lt", 0, nil},
	GT:        {"gt", 0, nil},
	LTE:       {"lte", 0, nil},
	GTE:       {"gte", 0, nil},
	SHIFTL:    {"shiftl", 1, []int{BYTE}},
	SHIFTR:    {"shiftl", 1, []int{BYTE}},
	NOP:       {"nop", 0, nil},
	JMP:       {"jmp", 1, []int{LABEL}},
	JMPIF:     {"jmpif", 1, []int{LABEL}},
	CALL:      {"call", 2, []int{LABEL, BYTE}},
	CALLIF:    {"callif", 2, []int{LABEL, BYTE}},
	CALLEXT:   {"callext", 3, []int{ADDR, BYTE, BYTE, BYTE, BYTE, BYTE}},
	RET:       {"ret", 0, nil},
	SIZE:      {"size", 0, nil},
	STORE:     {"store", 0, nil},
	SSTORE:    {"sstore", 1, []int{BYTE}},
	LOAD:      {"load", 1, []int{BYTE}},
	SLOAD:     {"sload", 1, []int{BYTE}},
	ADDRESS:   {"address", 0, nil},
	ISSUER:    {"issuer", 0, nil},
	BALANCE:   {"balance", 0, nil},
	CALLER:    {"caller", 0, nil},
	CALLVAL:   {"callval", 0, nil},
	CALLDATA:  {"calldata", 0, nil},
	NEWMAP:    {"newmap", 0, nil},
	MAPHASKEY: {"maphaskey", 0, nil},
	MAPPUSH:   {"mappush", 0, nil},
	MAPGETVAL: {"mapgetval", 0, nil},
	MAPSETVAL: {"mapsetval", 0, nil},
	MAPREMOVE: {"mapremove", 0, nil},
	NEWARR:    {"newarr", 0, nil},
	ARRAPPEND: {"arrappend", 0, nil},
	ARRINSERT: {"arrinsert", 0, nil},
	ARRREMOVE: {"arrremove", 0, nil},
	ARRAT:     {"arrat", 0, nil},
	SHA3:      {"sha3", 0, nil},
	CHECKSIG:  {"checksig", 0, nil},
	HALT:      {"halt", 0, nil},
	ERRHALT:   {"errhalt", 0, nil},
}
