// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract InfoRequestContract {

    // 存储每个用户的最新请求
    mapping(address => string) public newestRequests;

    event NewRequest(address user, string userId, uint32 reqType);

    function getCategory(string memory userId) public {
        newestRequests[msg.sender] = userId;
        emit NewRequest(msg.sender, userId, 1);
    }

    function getPubKey(string memory userId) public {
        newestRequests[msg.sender] = userId;
        emit NewRequest(msg.sender, userId, 0);
    }
}