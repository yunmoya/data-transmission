package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	userreq "test/abi"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VMReq struct {
	ID              string `json:"ID"`              // user request id
	UserId          string `json:"UserId"`          // user id
	PubKey          []byte `json:"PubKey"`          // user public key
	Config          int    `json:"Configuration"`   // cpu num
	Duration        int    `json:"Duration"`        // duration for which the VM is requested
	EncryptRequired bool   `json:"EncryptRequired"` //true：encryption required false: no encryption required
}

type ReqData struct {
	UserId          string `json:"UserId"`          // user id
	PubKey          []byte `json:"PubKey"`          // user public key
	Config          int    `json:"Configuration"`   // cpu num
	Duration        int    `json:"Duration"`        // duration for which the VM is requested
	EncryptRequired bool   `json:"EncryptRequired"` //true：encryption required false: no encryption required
}

func main() {
	// client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/547bcf201d264561b95c50860d2dddff")
	client, err := ethclient.Dial("ws://localhost:7545") //ganache
	if err != nil {
		log.Fatal(err)
	}

	// contractAddress := common.HexToAddress("0x74770669068090D6dCeA5163aBD5af61829A647a")
	contractAddress := common.HexToAddress("0x59Af076b2A8f82A17062c78626B226A1cAd5CACC") //ganache
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			log.Print("New Event:")
			vLogJson, err := json.Marshal(vLog)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(string(vLogJson)) // pointer to event log
			// if get new block in public blockchain
			// request propagation contract(ec)
			// ec >= sum/2 => schedule contract(fair)
			// 1. Get request, blockNumber and offset in the newest block
			blockNumber := vLog.BlockNumber
			offset := vLog.TxIndex
			var reqData ReqData

			contractAbi, err := abi.JSON(strings.NewReader(string(userreq.UserreqMetaData.ABI)))
			if err != nil {
				log.Fatal(err)
			}
			err = contractAbi.UnpackIntoInterface(&reqData, "NewRequest", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			req := VMReq{
				ID:              GetECAssetId(blockNumber, int(offset)),
				UserId:          reqData.UserId,
				Config:          reqData.Config,
				Duration:        reqData.Duration,
				PubKey:          reqData.PubKey,
				EncryptRequired: reqData.EncryptRequired,
			}
			log.Printf("The user request for private blockchain is: %v", req)
		}
	}
}

func GetECAssetId(blockNumber uint64, offset int) string {
	return "asset" + strconv.FormatUint(blockNumber, 10) + "-" + strconv.Itoa(offset)
}
