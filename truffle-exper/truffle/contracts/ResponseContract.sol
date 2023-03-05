// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract ResponseContract {
    struct Resp {
        string userId; // user id
        string dataHash; // hash of response data
        string dataSig; // signature of data
        bytes data; // data
        bool encryptRequired; //true：encryption required false: no encryption required
    }

    // 存储最新响应
    mapping(address => Resp) public newestResponse;

    event NewResponse(address from, address to, string userId, string dataHash, string dataSig, bytes data, bool encryptRequired);

    function responseToUser(address to, string memory userId, string memory dataHash, string memory dataSig, bytes memory data, bool encryptRequired) public {
        newestResponse[msg.sender] = Resp({
            userId: userId,
            dataHash: dataHash,
            dataSig: dataSig,
            data: data,
            encryptRequired: encryptRequired
        });

        emit NewResponse(msg.sender, to, userId, dataHash, dataSig, data, encryptRequired);
    }
}