CALLDATA
# ------- ABI -------------------------
DUP
PUSH 1
EQ
CALLIF mint 3
HALT

# ------ Contract ---------------------
mint:
LOAD 1 # load value
LOAD 0 # load key
SLOAD 1 # load address of minter
CALLER
EQ
CALLIF adjustBalance 2
RET

adjustBalance:
LOAD 1 # load value
LOAD 0 # load key
DUP
SLOAD 2 # load map
MAPHASKEY
CALLIF addKeyIfNotExists 2
LOAD 1 # load value
LOAD 0 # load key
SLOAD 2 # load map
MAPPUSH
SSTORE 2 # store the map
HALT

addKeyIfNotExists:
LOAD 1 # load key
SLOAD 2 # load map
MAPGETVAL
LOAD 0 # load value
ADD
LOAD 1 # load key
SLOAD 2  # load map
MAPSETVAL
SSTORE 2 # store the map
HALT