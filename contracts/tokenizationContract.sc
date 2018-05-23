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
POP
CALL mint 3

DUP
PUSH 2
EQ
POP
CALL send 0
HALT


# Contract

mint:
LOAD 0
LOAD 1
ADDRESS
CALLER
EQ
POP
CALL adjustBalance 2
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