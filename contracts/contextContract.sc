CALLDATA

# ABI
PUSH 1
EQ
CALLIF inc 1

HALT

# Contract
inc:
LOAD 0
ISSUER
CALLER
EQ
CALLIF adjustBalance 2
RET

adjustBalance:
LOAD 0
SLOAD 0
ADD
SSTORE 0
RET
