const HDWalletProvider = require('@truffle/hdwallet-provider');
const fs = require('fs');
module.exports = {
  // contracts_build_directory: "../client/src/ganache/contracts",
  contracts_build_directory: "../client/src/contracts",
  networks: {
    loc_development_development: {
      network_id: "*",
      port: 7545,
      host: "127.0.0.1"
    },
    inf_CloudService_goerli: {
      network_id: 5,
      gasPrice: 10000000000,
      provider: new HDWalletProvider(fs.readFileSync('d:\\onedrive\\workspace\\fabric-exper\\application\\gas.env', 'utf-8'), "https://goerli.infura.io/v3/547bcf201d264561b95c50860d2dddff")
    },
    inf_CloudService_sepolia: {
      network_id: 11155111,
      gasPrice: 10000000000,
      provider: new HDWalletProvider(fs.readFileSync('d:\\onedrive\\workspace\\fabric-exper\\application\\gas.env', 'utf-8'), "https://sepolia.infura.io/v3/547bcf201d264561b95c50860d2dddff")
    }
  },
  mocha: {},
  compilers: {
    solc: {
      version: "0.8.14"
    }
  }
};
