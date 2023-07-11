package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const (
	THRESHOLD              = 10
	CSPNUM                 = 3
	REQ_PROPORTION_PATH    = "/config/proportion/"
	REQ_FINDCSPCONFIG_PATH = "/config/exist/"
	REQ_PORT               = 8080
)

var IPArr = [CSPNUM + 1]string{"", "10.1.0.133", "10.1.0.154", "10.1.0.155"}

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type Asset struct {
	ID          string    `json:"ID"`          // asset id
	Proportions []float64 `json:"Proportions"` // proportions of csp contributions
	Historys    []int     `json:"Historys"`    // a queue containing results for past requests
}

type VMReq struct {
	ID       uint64 `json:"ID"`            // user request id
	pubKey   []byte `json:"PubKey"`        // user public key
	config   int    `json:"Configuration"` // cpu num
	duration uint64 `json:"Duration"`      // duration for which the VM is requested
}

type ProportionResp struct {
	Proportion float64 `json:"Proportion"` // CSP Proportion
}

type FindCSPResp struct {
	Exist bool `json:"Exist"` // Is config exist in the csp
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) (string, error) {
	historys := []int{}
	var proportions [CSPNUM + 1]float64
	// var wg sync.WaitGroup

	// TODO muti-thread
	// respChannel := make(chan map[int]float64)
	// go func() {
	// 	for cspProportionMap := range respChannel {
	// 		for key, val := range cspProportionMap {
	// 			fmt.Printf("key:%d value:%f \n", key, val)
	// 			proportions[key] = val
	// 		}
	// 	}
	// }()
	sum := 0.0
	for i := 1; i < CSPNUM+1; i++ {
		// wg.Add(1)
		// go getCSPProportionThread(i, respChannel, &wg)
		proportion, err := getCSPProportion(i)
		if err != nil {
			return "", err
		}
		proportions[i] = proportion
		sum += proportion
	}
	// wg.Wait()
	for i := 1; i < CSPNUM+1; i++ {
		proportions[i] /= sum
	}
	fmt.Print("The proportions are: ")
	fmt.Println(proportions)

	asset := Asset{ID: "1", Proportions: proportions[:], Historys: historys[:]}

	assetJSON, err := json.Marshal(asset)
	// fmt.Println(assetJSON)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(asset.ID, assetJSON)
	if err != nil {
		return "", fmt.Errorf("failed to put to world state. %v", err)
	}

	return string(assetJSON), nil
}

func (s *SmartContract) SelectCSP(ctx contractapi.TransactionContextInterface, config int) (int, error) {
	assetId := "1"
	exists, err := s.AssetExists(ctx, assetId)
	if err != nil {
		return -1, err
	}
	if !exists {
		return -1, fmt.Errorf("the asset %s does not exists", assetId)
	}

	asset, err := s.ReadAsset(ctx, assetId)
	if err != nil {
		return -1, err
	}

	// Find all csps which can provide VM. The VMs needs match user's requirement.
	cspList := FindCSPList(config)

	maxCSPId := -1
	maxScore := -100.0
	for _, cspId := range cspList {
		sum := 0.0
		for _, curId := range asset.Historys {
			if curId == cspId {
				sum += 1
			}
		}

		if len(asset.Historys) > 0 {
			sum /= float64(len(asset.Historys))
		}

		score := sum - asset.Proportions[cspId]
		if score >= float64(maxScore) {
			maxScore = score
			maxCSPId = cspId
		}
	}

	if maxCSPId == -1 {
		return -1, fmt.Errorf("There must be some problems about proportion computation")
	}
	asset.Historys = append(asset.Historys, maxCSPId)
	if len(asset.Historys) > THRESHOLD {
		asset.Historys = asset.Historys[1:] // delete the first element
	}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return -1, err
	}

	err = ctx.GetStub().PutState(assetId, assetJSON)
	if err != nil {
		return -1, err
	}

	return maxCSPId, nil
}

func (s *SmartContract) UpdateProportions(ctx contractapi.TransactionContextInterface) error {
	// respChannel := make(chan map[int]float64)
	// go getCSPProportionThread(cspId, respChannel)

	assetId := "1"
	exists, err := s.AssetExists(ctx, assetId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exists", assetId)
	}

	asset, err := s.ReadAsset(ctx, assetId)
	if err != nil {
		return err
	}
	// cspProportionMap := <-respChannel
	// asset.Proportions[cspId] = cspProportionMap[cspId]
	sum := 0.0
	var proportions [CSPNUM + 1]float64

	for i := 1; i < CSPNUM+1; i++ {
		proportion, err := getCSPProportion(i)
		if err != nil {
			return err
		}
		proportions[i] = proportion
		sum += proportion
	}
	// wg.Wait()
	for i := 1; i < CSPNUM+1; i++ {
		proportions[i] /= sum
	}
	asset.Proportions = proportions[:]

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(assetId, assetJSON)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(assetId, assetJSON)
}

func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func FindCSPList(config int) []int {
	result := []int{}
	for i := 1; i < CSPNUM+1; i++ {
		exist, err := FindCSPConfig(config, i)
		if err != nil {
			fmt.Printf("Can't judge whether the configuration%d exists in CSP%d, caused by: %v", config, i, err)
		}
		if exist {
			result = append(result, i)
		}
	}

	fmt.Print("The findCSPList is: ")
	fmt.Println(result)
	return result
}

func FindCSPConfig(config int, cspId int) (bool, error) {
	ip := IPArr[cspId]
	url := "http://" + ip + ":" + strconv.Itoa(REQ_PORT) + REQ_FINDCSPCONFIG_PATH + strconv.Itoa(config)

	// fmt.Println("Prepare to request url: " + url)
	rsps, err := http.Get(url)

	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}
	var findCSPConfig FindCSPResp
	err = json.Unmarshal(body, &findCSPConfig)
	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}

	return findCSPConfig.Exist, nil
}

func getCSPProportionThread(cspId int, respChannel chan map[int]float64, wg *sync.WaitGroup) {
	defer wg.Done()
	cspProportionMap := make(map[int]float64)
	cspProportionMap[cspId] = -1.0

	proportion, err := getCSPProportion(cspId)
	if err != nil {
		fmt.Println(err)
		respChannel <- cspProportionMap
		return
	}
	cspProportionMap[cspId] = proportion
	respChannel <- cspProportionMap
}

func getCSPProportion(cspId int) (float64, error) {
	ip := IPArr[cspId]
	url := "http://" + ip + ":" + strconv.Itoa(REQ_PORT) + REQ_PROPORTION_PATH

	// fmt.Println("Prepare to request url: " + url)
	rsps, err := http.Get(url)
	if err != nil {
		return -1, fmt.Errorf("error: %s", err)
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		return -1, fmt.Errorf("error: %s", err)
	}

	var proportion ProportionResp
	err = json.Unmarshal(body, &proportion)
	if err != nil {
		return -1, fmt.Errorf("error: %s", err)
	}

	return proportion.Proportion, nil
}

func main() {
	assetChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
