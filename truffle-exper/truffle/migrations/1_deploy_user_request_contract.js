const UserRequestContract = artifacts.require("UserRequestContract");

module.exports = function (deployer) {
  deployer.deploy(UserRequestContract);
};
