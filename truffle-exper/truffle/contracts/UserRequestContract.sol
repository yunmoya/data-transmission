// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract UserRequestContract {
    struct Req {
        string userId; // user id
        string pubKey; // user public key
        uint32 config; // cpu num
        uint32 duration; // duration for which the VM is requested
        bool encryptRequired; //true：encryption required false: no encryption required
    }

    // 存储每个用户的最新请求
    mapping(address => Req) public newestRequests;

    event NewRequest(address user, string userId, string pubKey, uint32 config, uint32 duration, bool encryptRequired);

    function requestVMWithoutEncryption(string memory userId, uint32 config, uint32 duration) public {
        newestRequests[msg.sender] = Req({
            userId: userId,
            pubKey: "",
            config: config,
            duration: duration,
            encryptRequired: false
        });

        emit NewRequest(msg.sender, userId, "", config, duration, false);
    }

    function requestVMWithEncryption(string memory userId, uint32 config, uint32 duration, string memory pubKey) public {
        newestRequests[msg.sender] = Req({
            userId: userId,
            pubKey: pubKey,
            config: config,
            duration: duration,
            encryptRequired: true
        });

        emit NewRequest(msg.sender, userId, pubKey, config, duration, true);
    }
}