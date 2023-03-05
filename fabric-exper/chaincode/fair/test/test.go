package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

const CSPNUM = 3
const REQ_PORT = 8080
const REQ_PROPOTION_PATH = "/cloud/propotion"

var IPArr = [CSPNUM + 1]string{"", "10.1.0.133", "10.1.0.154", "10.1.0.155"}

type PropotionResp struct {
	Propotion float64 `json:"Propotion"` // CSP Propotion
}

type Asset struct {
	ID         string    `json:"ID"`         // asset id
	Propotions []float64 `json:"Propotions"` // propotions of csp contributions
	Historys   []int     `json:"Historys"`   // a queue containing results for past requests
}

func main() {
	var historys [10]int
	var propotions [CSPNUM + 1]float64
	var wg sync.WaitGroup
	respChannel := make(chan map[int]float64)

	for i := 0; i < len(historys); i++ {
		historys[i] = -1
	}
	go func() {
		for cspPropotionMap := range respChannel {
			for key, val := range cspPropotionMap {
				fmt.Printf("key:%d value:%f \n", key, val)
				propotions[key] = val
			}
		}
	}()

	for i := 1; i < 4; i++ {
		wg.Add(1)
		go getCSPPropotion(i, respChannel, &wg)
	}

	wg.Wait()
	close(respChannel)

	fmt.Print("The propotions are: ")
	fmt.Println(propotions)

	asset := Asset{ID: "1", Propotions: propotions[:], Historys: historys[:]}

	assetJSON, err := json.Marshal(asset)
	fmt.Println(string(assetJSON))
	if err != nil {
		fmt.Println(err)
	}
}

func getCSPPropotion(cspId int, respChannel chan map[int]float64, wg *sync.WaitGroup) {
	defer wg.Done()
	ip := IPArr[cspId]
	url := "http://" + ip + ":" + strconv.Itoa(REQ_PORT) + REQ_PROPOTION_PATH
	cspPropotionMap := make(map[int]float64)
	cspPropotionMap[cspId] = 0.0

	fmt.Println("Prepare to request url: " + url)
	rsps, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s", err)
		respChannel <- cspPropotionMap
		return
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		fmt.Printf("error: %s", err)
		respChannel <- cspPropotionMap
		return
	}

	var propotion PropotionResp
	err = json.Unmarshal(body, &propotion)
	if err != nil {
		fmt.Printf("error: %s", err)
		respChannel <- cspPropotionMap
		return
	}

	cspPropotionMap[cspId] = propotion.Propotion
	fmt.Printf("Get Csp%d propotion\n", cspId)
	respChannel <- cspPropotionMap
}
