


The storage layout of the peg token contract: 


```text
┌────────────────────┬────────────────┬──────────────┬────────┬─────────────────────────────────────────────────────┬─────┬───────────────────────────────────────────────────┬───────────────┐
│      contract      │ state_variable │ storage_slot │ offset │                        type                         │ idx │                     artifact                      │ numberOfBytes │
├────────────────────┼────────────────┼──────────────┼────────┼─────────────────────────────────────────────────────┼─────┼───────────────────────────────────────────────────┼───────────────┤
│   AccessControl    │     _roles     │      0       │   0    │  t_mapping(t_bytes32,t_struct(RoleData)18_storage)  │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│       ERC20        │   _balances    │      0       │   0    │           t_mapping(t_address,t_uint256)            │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│       ERC20        │  _allowances   │      1       │   0    │ t_mapping(t_address,t_mapping(t_address,t_uint256)) │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│       ERC20        │  _totalSupply  │      2       │   0    │                      t_uint256                      │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│       ERC20        │     _name      │      3       │   0    │                  t_string_storage                   │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│       ERC20        │    _symbol     │      4       │   0    │                  t_string_storage                   │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│       ERC20        │   _decimals    │      5       │   0    │                       t_uint8                       │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │       1       │
│ ERC20BlackListAble │     _roles     │      0       │   0    │  t_mapping(t_bytes32,t_struct(RoleData)18_storage)  │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│ ERC20BlackListAble │   BlackList    │      1       │   0    │             t_mapping(t_address,t_bool)             │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Burnable    │   _balances    │      0       │   0    │           t_mapping(t_address,t_uint256)            │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Burnable    │  _allowances   │      1       │   0    │ t_mapping(t_address,t_mapping(t_address,t_uint256)) │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Burnable    │  _totalSupply  │      2       │   0    │                      t_uint256                      │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Burnable    │     _name      │      3       │   0    │                  t_string_storage                   │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Burnable    │    _symbol     │      4       │   0    │                  t_string_storage                   │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Burnable    │   _decimals    │      5       │   0    │                       t_uint8                       │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │       1       │
│    ERC20Factory    │     tokens     │      0       │   0    │            t_array(t_address)dyn_storage            │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│    ERC20Factory    │      root      │      1       │   0    │                      t_address                      │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      20       │
│   ERC20Pausable    │   _balances    │      0       │   0    │           t_mapping(t_address,t_uint256)            │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Pausable    │  _allowances   │      1       │   0    │ t_mapping(t_address,t_mapping(t_address,t_uint256)) │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Pausable    │  _totalSupply  │      2       │   0    │                      t_uint256                      │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Pausable    │     _name      │      3       │   0    │                  t_string_storage                   │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Pausable    │    _symbol     │      4       │   0    │                  t_string_storage                   │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20Pausable    │   _decimals    │      5       │   0    │                       t_uint8                       │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │       1       │
│   ERC20Pausable    │    _paused     │      5       │   1    │                       t_bool                        │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │       1       │
│   ERC20PegToken    │   _balances    │      0       │   0    │           t_mapping(t_address,t_uint256)            │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20PegToken    │  _allowances   │      1       │   0    │ t_mapping(t_address,t_mapping(t_address,t_uint256)) │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20PegToken    │  _totalSupply  │      2       │   0    │                      t_uint256                      │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20PegToken    │     _name      │      3       │   0    │                  t_string_storage                   │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20PegToken    │    _symbol     │      4       │   0    │                  t_string_storage                   │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20PegToken    │   _decimals    │      5       │   0    │                       t_uint8                       │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │       1       │
│   ERC20PegToken    │     _roles     │      6       │   0    │  t_mapping(t_bytes32,t_struct(RoleData)18_storage)  │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│   ERC20PegToken    │    _paused     │      7       │   0    │                       t_bool                        │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │       1       │
│   ERC20PegToken    │   BlackList    │      8       │   0    │             t_mapping(t_address,t_bool)             │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │      32       │
│      Pausable      │    _paused     │      0       │   0    │                       t_bool                        │  0  │ /build-info/2c4471da4a7c15b8e645e04f18dcbf07.json │       1       │
└────────────────────┴────────────────┴──────────────┴────────┴─────────────────────────────────────────────────────┴─────┴───────────────────────────────────────────────────┴───────────────┘
```



In the `genesis.json`, we have manuipluated the storage, adding some USDT to the USDT contract's address:  

```json
    "0039f574ee5cc39bdd162e9a88e3eb1f111baf48": {
      "balance": "0x0",
      "code": "0x6080",
      "storage": {
        "6f20d50b7632994afdeb57b0a84f1f7aea84bfe45ef97c691dd856c2af5abf6c": "00000000000000000000000000000000000000000000022d2499dd996382ce88"
      }
    }
```