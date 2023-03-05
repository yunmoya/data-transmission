const ResponseContract = artifacts.require("ResponseContract");

module.exports = function (deployer) {
  deployer.deploy(ResponseContract);
};
