package contract

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type ECAsset struct {
	ID               string    `json:"ID"`
	SequenceNumber   SeqNumber `json:"SequenceNumber"`
	EndorsementCount int       `json:"EndorsementCount"`
	Status           int       `json:"Status"`
}

type SeqNumber struct {
	BlockNumber uint64 `json:"BlockNumber"`
	Offset      int    `json:"Offset"`
}

// This type of transaction would typically only be run once by an application the first time it was started after its
// initial deployment. A new version of the chaincode deployed later would likely not need to run an "init" function.
func InitECLedger(ecContract *client.Contract) {
	log.Print("============ Init EC Ledger ============")
	_, err := ecContract.SubmitTransaction("InitLedger")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	log.Print("============ Transaction committed successfully ============")
}

// This type of transaction would typically only be run once by an application the first time it was started after its
// initial deployment. A new version of the chaincode deployed later would likely not need to run an "init" function.
func InitFairLedger(fairContract *client.Contract) {
	log.Print("============ Init Fair Ledger ============")
	_, err := fairContract.SubmitTransaction("InitLedger")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	log.Print("============ Transaction committed successfully ============")
}

func StatusAssetExists(contract *client.Contract, assetId string) bool {
	log.Print("============ EC Asset Exist ============")

	evaluateResult, err := contract.EvaluateTransaction("StatusAssetExists", assetId)
	if err != nil {
		log.Panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := false
	err = json.Unmarshal(evaluateResult, &result)
	if err != nil {
		log.Panic(fmt.Errorf("failed to parse json: %w", err))
	}

	log.Printf("The result is: %v", result)
	return result
}

func CreateECLedger(contract *client.Contract, assetId string) {
	log.Print("============ Create EC Ledger ============")

	_, err := contract.SubmitTransaction("CreateLedger", assetId)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	log.Printf("============ Transaction committed successfully ============")
}

func EndorsementIncAsync(contract *client.Contract, assetId string) {
	log.Print("============ Async Increse Endorsement ============")

	_, commit, err := contract.SubmitAsync("EndorsementInc", client.WithArguments(assetId))
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction asynchronously: %w", err))
	}

	log.Printf("Successfully submitted transaction to increase %v's ec.", assetId)
	log.Println("Waiting for transaction commit.")

	if commitStatus, err := commit.Status(); err != nil {
		log.Panic(fmt.Errorf("failed to get commit status: %w", err))
	} else if !commitStatus.Successful {
		log.Panic(fmt.Errorf("transaction %s failed to commit with status: %d", commitStatus.TransactionID, int32(commitStatus.Code)))
	}

	log.Printf("============ Transaction committed successfully ============ ")
}

func ReadECAssetByID(contract *client.Contract, assetId string) *ECAsset {
	log.Print("============ Read EC Asset By Id ============")

	evaluateResult, err := contract.EvaluateTransaction("ReadEC", assetId)
	if err != nil {
		log.Panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := new(ECAsset)
	err = json.Unmarshal(evaluateResult, &result)
	if err != nil {
		log.Panic(fmt.Errorf("failed to parse json: %w", err))
	}

	log.Printf("The result is: %v", result)
	return result
}

func ChangeECAssetStatus(contract *client.Contract, assetId string) error {
	log.Print("============ Async Change EC Asset Status ============")

	_, err := contract.SubmitTransaction("SetStatus", assetId)
	if err != nil {
		return fmt.Errorf("failed to submit transaction: %w", err)
	}

	log.Printf("============ Transaction committed successfully ============ ")
	return nil
}

func SelectCSP(contract *client.Contract, config int) int {
	log.Print("============ Select CSP============")

	evaluateResult, err := contract.EvaluateTransaction("SelectCSP", strconv.Itoa(config))
	if err != nil {
		log.Panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := -1
	err = json.Unmarshal(evaluateResult, &result)
	if err != nil {
		log.Panic(fmt.Errorf("failed to parse json: %w", err))
	}

	log.Printf("The result is: %v", result)
	return result
}

func UpdateProportionsAsync(contract *client.Contract) {
	log.Print("============ Async Update Proportions ============")

	_, commit, err := contract.SubmitAsync("UpdateProportions")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction asynchronously: %w", err))
	}

	log.Printf("Successfully submitted transaction to update proportions.")
	log.Println("Waiting for transaction commit.")

	if commitStatus, err := commit.Status(); err != nil {
		log.Panic(fmt.Errorf("failed to get commit status: %w", err))
	} else if !commitStatus.Successful {
		log.Panic(fmt.Errorf("transaction %s failed to commit with status: %d", commitStatus.TransactionID, int32(commitStatus.Code)))
	}

	log.Printf("============ Transaction committed successfully ============ ")
}
