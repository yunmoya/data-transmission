package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// mspID = "Org1MSP"
// cryptoPath    = "../../test-network/organizations/peerOrganizations/org1.example.com"
// certPath      = "../../test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/cert.pem"
// keyPath       = "../../test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/"
// tlsCertPath   = "../../test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt"
// peerEndpoint  = "localhost:7051"
// gatewayPeer   = "peer0.org1.example.com"

type PeerInfo struct {
	Mspid        string `yaml:"msp-id"`
	CryptoPath   string `yaml:"crypto-path"`
	CertPath     string `yaml:"cert-path"`
	KeyPath      string `yaml:"key-path"`
	TlsCertPath  string `yaml:"tls-cert-path"`
	PeerEndpoint string `yaml:"peer-endpoint"`
	GatewayPeer  string `yaml:"gateway-peer"`
}

type WalletInfo struct {
	WalletPrivateKey    string `yaml:"wallet-private-key"`
	WalletPublicKey     string `yaml:"wallet-public-key"`
	WalletPublicAddress string `yaml:"wallet-public-address"`
}

func LoadPeerInfo(path string, peerInfo *PeerInfo) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &peerInfo)
	if err != nil {
		return err
	}

	log.Printf("Load peer information successfully! The information is %v", peerInfo)
	return nil
}

func LoadWalletInfo(path string, walletInfo *WalletInfo) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &walletInfo)
	if err != nil {
		return err
	}

	log.Printf("Load wallet information successfully! The information is %v", walletInfo)
	return nil
}

func WritePeerInfo() {
	peerInfo := PeerInfo{Mspid: "test"}

	fmt.Println(peerInfo)
	yamlContent, err := yaml.Marshal(peerInfo)
	fmt.Println(err)

	fmt.Println(string(yamlContent))
}
