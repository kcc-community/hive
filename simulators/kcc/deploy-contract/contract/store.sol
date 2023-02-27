pragma solidity >=0.8.0;

contract Store {
    event ItemSet(bytes32 key, bytes32 value);

    mapping(bytes32 => bytes32) public items;

    function setItem(bytes32 key, bytes32 value) external {
        items[key] = value;
        emit ItemSet(key, value);
    }
}