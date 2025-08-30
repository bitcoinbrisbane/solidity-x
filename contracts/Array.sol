// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ArrayContract {

    uint[] private myCollection;

    constructor() {
        // Constructor logic
    }

    function sort() public {
        for (uint i = 0; i < myCollection.length - 1; i++) {
            for (uint j = i + 1; j < myCollection.length; j++) {
                if (myCollection[i] > myCollection[j]) {
                    (myCollection[i], myCollection[j]) = (myCollection[j], myCollection[i]);
                }
            }
        }
    }

    function find(uint target) public view returns (int) {
        for (uint i = 0; i < myCollection.length; i++) {
            if (myCollection[i] == target) {
                return int(i);
            }
        }
        return -1;
    }

    function filter(uint[] memory arr, uint target) public pure returns (int) {
        for (uint i = 0; i < arr.length; i++) {
            if (arr[i] == target) {
                return int(i);
            }
        }
        return -1;
    }
}