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
ISSUER
CALLER
EQ
CALLIF adjustBalance 2
RET

adjustBalance:
LOAD 1
SLOAD 0
MAPGETVAL
LOAD 0
ADD
MAPSETVAL
SSTORE 0
RET
