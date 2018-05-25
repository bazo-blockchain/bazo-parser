# index 0 ist leere MAP
# Minter

# amount
# receiver
# func hash

CALLDATA

# ABI
DUP
PUSH 1
EQ
CALLIF mint 3

HALT

# Contract
mint:
LOAD 0
LOAD 1
SLOAD 1
CALLER
EQ
CALLIF adjustBalance 2
RET

adjustBalance:
LOAD 0
SLOAD 2
MAPGETVAL
LOAD 1
ADD
SLOAD 2
MAPSETVAL
SSTORE 2
RET
