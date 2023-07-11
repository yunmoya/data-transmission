package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/big"
	"simul/inforeq"
	"simul/response"
	"time"
)

type WalletInfo struct {
	WalletPrivateKey    string `yaml:"wallet-private-key"`
	WalletPublicAddress string `yaml:"wallet-public-address"`
}

const (
	InfuraURL                   = "wss://goerli.infura.io/ws/v3/547bcf201d264561b95c50860d2dddff" //goerli or sepolia
	UserReqContractAddr         = "0xA9785c5302eF4Ac6A43B00B41fE2ADDf15d79Ed1"
	ResponseContractAddr        = "0xf64646DF059840890AaDdaaAE6D2A14Bd8Ecb5f9"
	GoerliInfoContractAddr      = "0xB2286b07f2282aE74EA82254D0174258a9EADBB0"
	SepoliaInfoContractAddr     = "0xD693a92435a4f2D5A44ea868a8eC09ada5032FDF"
	GoerliResponseContractAddr  = "0xEf666c73091a4D66632760F28289f6A54e7b0EEb"
	SepoliaResponseContractAddr = "0x12A0B1A747024096B59CC55BB2F3A9a9A42CbcB4"
)

func main() {
	// ehereum ethClient config
	ethClient, err := ethclient.Dial(InfuraURL)
	if err != nil {
		log.Fatal(err)
	}

	//tx := common.HexToHash("0x35d67e24d6ba543b6b2d6969819db6432591bdbca6f73a40d134391403a5144a")
	//txRec, err := ethClient.TransactionReceipt(context.Background(), tx)
	//if err != nil {
	//	log.Print(err)
	//} else {
	//	blockNumber := txRec.BlockNumber
	//	gasUsed := txRec.GasUsed
	//
	//	block, err := ethClient.BlockByNumber(context.Background(), blockNumber)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	commitTime := block.Time()
	//	log.Printf("gas %v commitTime %v", gasUsed, commitTime)
	//}

	// eth account config
	var WalletName = fmt.Sprintf("wallet.yaml")
	var WalletInfo = new(WalletInfo)
	err = LoadWalletInfo(WalletName, WalletInfo)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(WalletInfo.WalletPrivateKey)
	if err != nil {
		log.Printf("Failed to get private key: %w", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Fatal("Failed to get nonce: %w", err)
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to get gasPrice: %w", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice
	print(gasPrice)

	sendCategory(ethClient, auth)
	//sendPubkey(ethClient, auth)
	//reqPubkey(ethClient, auth)
	//reqCategory(ethClient, auth)

}

func sendCategory(client *ethclient.Client, auth *bind.TransactOpts) {
	address := common.HexToAddress(GoerliResponseContractAddr)
	instance, err := response.NewResponse(address, client)
	if err != nil {
		log.Fatal("Failed to create instance: %w", err)
	}
	toAddr := common.HexToAddress("0x457ea1A9E67403DAdeF225BDBDDA2cE92fc148bf")
	userId := "0x457ea1A9E67403DAdeF225BDBDDA2cE92fc148bf"
	hash := "baba3d3071f983d3fa4fdaa4832fbc86598f61d9983a538f16bc8eb1eb722337"
	sig := "50bb52a2d34fd840796cedca8d2689c11f6b5fa52991145ec916a2abe2608c92083ff7b9c18a1cdbadec0c2dc9b4c05ba19fdc06562931c396706e55fb3327e307"
	msg := map[int]int{8: 100, 64: 20, 32: 10, 128: 5, 256: 30}
	msgBytes, err := json.Marshal(msg)
	print(len(msgBytes))
	tx, err := instance.ResponseToUser(auth, toAddr, userId, hash, sig, msgBytes, false)
	if err != nil {
		log.Fatal("Failed to write to response contract: ", err)
	}
	log.Printf("now is %v, tx sent: %v", time.Now().UnixMilli(), tx.Hash().Hex())
}

func sendPubkey(client *ethclient.Client, auth *bind.TransactOpts) {
	address := common.HexToAddress(SepoliaResponseContractAddr)
	instance, err := response.NewResponse(address, client)
	if err != nil {
		log.Fatal("Failed to create instance: %w", err)
	}
	toAddr := common.HexToAddress("0x457ea1A9E67403DAdeF225BDBDDA2cE92fc148bf")
	userId := "0x457ea1A9E67403DAdeF225BDBDDA2cE92fc148bf"
	hash := "baba3d3071f983d3fa4fdaa4832fbc86598f61d9983a538f16bc8eb1eb722337"
	sig := "50bb52a2d34fd840796cedca8d2689c11f6b5fa52991145ec916a2abe2608c92083ff7b9c18a1cdbadec0c2dc9b4c05ba19fdc06562931c396706e55fb3327e307"
	msg := [...]string{"2e2295087d7011d183204066924abe42ea6270e3968d6f3e3389180d4be3683d", "5587544811422d359a880f7cb9289a5f8ca8627228320b7e9f92d6bfeb7e26f9", "e171a5f49a355eaf054273e879c1b6242843101bc3ef49585ccfdc807947db8c"}
	msgBytes, err := json.Marshal(msg)
	tx, err := instance.ResponseToUser(auth, toAddr, userId, hash, sig, msgBytes, false)
	if err != nil {
		log.Fatal("Failed to write to response contract: %w", err)
	}
	log.Printf("now is %v, tx sent: %v", time.Now().UnixMilli(), tx.Hash().Hex())
}

func reqPubkey(client *ethclient.Client, auth *bind.TransactOpts) {
	address := common.HexToAddress(GoerliInfoContractAddr)

	instance, err := inforeq.NewInforeq(address, client)
	if err != nil {
		log.Fatal("Failed to create instance: %w", err)
	}

	tx, err := instance.GetPubKey(auth, "user1")
	if err != nil {
		log.Fatal("Failed to write to response contract: %w", err)
	}

	log.Printf("now is %v, tx sent: %v", time.Now().UnixMilli(), tx.Hash().Hex())
}

func reqCategory(client *ethclient.Client, auth *bind.TransactOpts) {
	address := common.HexToAddress(GoerliInfoContractAddr)

	instance, err := inforeq.NewInforeq(address, client)
	if err != nil {
		log.Fatal("Failed to create instance: %w", err)
	}

	tx, err := instance.GetCategory(auth, "user1")
	if err != nil {
		log.Fatal("Failed to write to response contract: %w", err)
	}

	log.Printf("now is %v, tx sent: %v", time.Now().UnixMilli(), tx.Hash().Hex())
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

func infoEvent(ethClient *ethclient.Client, contractAddress string) {
	contractAddr := common.HexToAddress(contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			tx := vLog.TxHash
			log.Printf("Now is %v, get new event from tx%v: ", time.Now().UnixMilli(), tx)
			txRec, err := ethClient.TransactionReceipt(context.Background(), tx)
			if err != nil {
				log.Print(err)
			} else {
				blockNumber := txRec.BlockNumber
				gasUsed := txRec.GasUsed

				block, err := ethClient.BlockByNumber(context.Background(), blockNumber)
				if err != nil {
					log.Fatal(err)
				}
				commitTime := block.Time()
				log.Printf("gas %v commitTime %v", gasUsed, commitTime)
			}
		}
	}
}
