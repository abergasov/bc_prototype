// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

contract CircleStorage {

    struct CircleMember {
        uint256 totalSeedsAmount;
        address member;
    }

    CircleMember[] public seedsOwners;

    mapping(address => uint256) public nameToSeedsAmount;

    bytes16 public immutable CIRCLE_ADDRESS;

    uint256 public constant TOTAL_SEEDS = 1000;

    constructor(bytes16 _circleID) {
        CIRCLE_ADDRESS = _circleID;
        nameToSeedsAmount[msg.sender] = TOTAL_SEEDS;
//        CircleMember memory member = CircleMember({
//            totalSeedsAmount: TOTAL_SEEDS,
//            member: msg.sender
//        });
//        seedsOwners.push(member);
    }
}