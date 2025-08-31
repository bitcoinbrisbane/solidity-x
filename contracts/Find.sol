// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract FindResultContract {

    int[] private myCollection;

    constructor() {
        // Constructor logic
    }

    function example(int target) public view returns (int) {
        int result = _find(target);
        return result;
    }

    function _find(int target) private view returns (int) {
        for (uint i = 0; i < myCollection.length; i++) {
            if (myCollection[i] == target) {
                return int(i);
            }
        }
        return -1;
    }
}