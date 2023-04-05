// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

contract CircleContract {

    struct CircleMember {
        uint256 totalSeedsAmount;
        address member;
    }

    CircleMember[] public seedsOwners;

    mapping(address => uint256) public nameToSeedsAmount;

    bytes16 public immutable CIRCLE_ADDRESS;

    uint256 public constant TOTAL_SEEDS = 1000;

    constructor(address _circleDirector, bytes16 _circleID) {
        CIRCLE_ADDRESS = _circleID;
        nameToSeedsAmount[_circleDirector] = TOTAL_SEEDS;
       // seedsOwners.push(_circleDirector);
    }

    function circle() public view returns (bytes16)  {
        return CIRCLE_ADDRESS;
    }
}