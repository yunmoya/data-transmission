const InfoRequestContract = artifacts.require("InfoRequestContract");

module.exports = function (deployer) {
  deployer.deploy(InfoRequestContract);
};
