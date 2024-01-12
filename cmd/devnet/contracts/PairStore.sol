// SPDX-License-Identifier: GPL-3.0

pragma solidity >0.7.0 < 0.9.0;
/**
* @title Storage
* @dev store or retrieve variable value
*/

contract PairStore {
  uint256 private storedDataFirst;
  uint256 private storedDataSecond;

  function setFirst(uint256 value) public {
    storedDataFirst = value;
  }

  function setSecond(uint256 value) public {
    storedDataSecond = value;
  }

  function getFirst() public view returns (uint256) {
    return storedDataFirst;
  }

  function getSecond() public view returns (uint256) {
    return storedDataSecond;
  }
}