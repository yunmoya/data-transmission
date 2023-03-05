package main

import (
	myabi "app/abi"
	"app/config"
	"app/contract"
	"app/cosiUtil"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// global configuration
const (
	ConfigName             = "154"
	TestNetName            = "sepolia" //goerli or sepolia
	CspNum                 = 3
	ReqPort                = 8080
	PeerConsensusThreshold = 2
)

// configuration about fabric
const (
	ChannelName      = "lccch"
	EcContractName   = "ec"
	FairContractName = "fair"
)

var PeerInfoName = fmt.Sprintf("config%s.yaml", ConfigName)

// configuration about ethereum
const (
	SepoliaUserReqContractAddr  = "0x3211f20Ec71C9f8E32CAF306c41766BA80aCbe87"
	SepoliaResponseContractAddr = "0x12A0B1A747024096B59CC55BB2F3A9a9A42CbcB4"
	GoerilUserReqContractAddr   = "0xcD6695d74CC71fcdf65a2E5A52126fEDeA44b256"
	GoerilResponseContractAddr  = "0xEf666c73091a4D66632760F28289f6A54e7b0EEb"
)

var UserReqContractAddr = GoerilUserReqContractAddr
var ResponseContractAddr = GoerilResponseContractAddr

var WalletName = fmt.Sprintf("wallet/wallet%s.yaml", ConfigName)
var WalletInfo = new(config.WalletInfo)
var InfuraURL = fmt.Sprintf("wss://%s.infura.io/ws/v3/547bcf201d264561b95c50860d2dddff", TestNetName)

// configuration about cosi
const (
	rosterFileName = "/home/lcc/config/roster.toml"
)

// configuration about cloudservice
const (
	assignPath = "/assignInfo/assign"
)

var PeerInfo = new(config.PeerInfo)
var IPArr = [CspNum + 1]string{"", "10.1.0.133", "10.1.0.154", "10.1.0.155"}

type VMReq struct {
	ID              string `json:"ID"`              // user request id
	UserId          string `json:"UserId"`          // user id
	PubKey          string `json:"PubKey"`          // user public key
	Config          uint32 `json:"Configuration"`   // cpu num
	Duration        uint32 `json:"Duration"`        // duration for which the VM is requested
	EncryptRequired bool   `json:"EncryptRequired"` //true：encryption required false: no encryption required
}

type ReqData struct {
	User            common.Address
	UserId          string // user id
	PubKey          string // user public key
	Config          uint32 // cpu num
	Duration        uint32 // duration for which the VM is requested
	EncryptRequired bool   //true：encryption required false: no encryption required
}

type ReqBody struct {
	Config   int    `json:"config"`
	Duration int    `json:"duration"`
	UserId   string `json:"userId"`
}

type RespBody struct {
	Msg      string   `json:"msg"`
	Code     int      `json:"code"`
	AssignVO AssignVO `json:"data"`
}

type AssignVO struct {
	VmId       uint64 `json:"vmId"`
	UserId     string `json:"userId"`
	VmName     string `json:"vmName"`
	AssignTime string `json:"assignTime"`
	EndTime    string `json:"endTime"`
}

func main() {
	log.Println("============ application starts ============")
	err := config.LoadPeerInfo(PeerInfoName, PeerInfo)
	if err != nil {
		log.Panic(err)
	}

	// fabric ethClient config
	// The gRPC ethClient connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()
	gw := createFabricClient(clientConnection)
	defer clientConnection.Close()
	defer gw.Close()
	network := gw.GetNetwork(ChannelName)

	// fabric contract
	ecContract := network.GetContract(EcContractName)
	fairContract := network.GetContract(FairContractName)
	contract.InitECLedger(ecContract)
	contract.InitFairLedger(fairContract)

	// ehereum ethClient config
	ethClient, err := ethclient.Dial(InfuraURL)
	//ethClient, err := ethclient.Dial("ws://localhost:7545") //ganache
	if err != nil {
		log.Fatal(err)
	}
	// eth account config
	err = config.LoadWalletInfo(WalletName, WalletInfo)
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
	ethFromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if TestNetName == "sepolia" {
		UserReqContractAddr = SepoliaUserReqContractAddr
		ResponseContractAddr = SepoliaResponseContractAddr
	}
	userReqContractAddr := common.HexToAddress(UserReqContractAddr)
	//userReqContractAddr := common.HexToAddress("0x59Af076b2A8f82A17062c78626B226A1cAd5CACC") //ganache
	query := ethereum.FilterQuery{
		Addresses: []common.Address{userReqContractAddr},
	}

	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(myabi.UserreqMetaData.ABI)))

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			log.Printf("Now is %v, get new event from tx%v: ", time.Now().UnixMilli(), vLog.TxHash)
			//log.Print("New Event:")
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
			reqData := new(ReqData)
			err = contractAbi.UnpackIntoInterface(reqData, "NewRequest", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			log.Print("Parse request data successfully: ", reqData)
			assetECId := getECAssetId(blockNumber, int(offset))
			req := VMReq{
				ID:              assetECId,
				UserId:          reqData.UserId,
				Config:          reqData.Config,
				Duration:        reqData.Duration,
				PubKey:          reqData.PubKey,
				EncryptRequired: reqData.EncryptRequired,
			}
			log.Print("The user request for private blockchain is: ", req)
			// 2. update the EC Asset Id
			log.Printf("Now is %v, update ec", time.Now().UnixMilli())
			contract.EndorsementIncAsync(ecContract, assetECId)

			// 3. Check if ecNum > 3/2 && request's status doesn't exist
			ecAsset := contract.ReadECAssetByID(ecContract, assetECId)

			for {
				if ecAsset.EndorsementCount < PeerConsensusThreshold || contract.StatusAssetExists(ecContract, assetECId) {
					break
				} else {
					// 4. update status to completed(1)
					err = contract.ChangeECAssetStatus(ecContract, assetECId)
					if err != nil {
						log.Print(err)
						continue
					} else {
						// 5. fair scheduled contract process request
						processReq(fairContract, ethClient, privateKey, ethFromAddress, reqData.User, req)
					}
				}
			}
			log.Printf("Now is %v, complete ec", time.Now().UnixMilli())
		}
	}
}

func createFabricClient(clientConnection *grpc.ClientConn) *client.Gateway {
	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		log.Panic(err)
	}
	return gw
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	var tlsCertPath = PeerInfo.TlsCertPath
	var gatewayPeer = PeerInfo.GatewayPeer
	var peerEndpoint = PeerInfo.PeerEndpoint

	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		log.Panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		log.Panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity() *identity.X509Identity {
	var certPath = PeerInfo.CertPath
	var mspID = PeerInfo.Mspid

	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	var keyPath = PeerInfo.KeyPath

	files, err := os.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(keyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func getECAssetId(blockNumber uint64, offset int) string {
	return "asset" + strconv.FormatUint(blockNumber, 10) + "-" + strconv.Itoa(offset)
}

func processReq(
	fairContract *client.Contract,
	ethClient *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	fromAddr common.Address,
	toAddr common.Address,
	req VMReq,
) {
	log.Printf("Now is %v, select CP", time.Now().UnixMilli())
	cspId := contract.SelectCSP(fairContract, int(req.Config))
	if cspId == -1 {
		log.Printf("select csp fail, the request's detail is %v", req.UserId, req)
		return
	}

	log.Printf("Now is %v, req for assign", time.Now().UnixMilli())
	assignVO, err := getAssignInfo(cspId, req)
	if err != nil {
		log.Panic(err)
	}
	if assignVO == nil {
		log.Printf("process user:%s request fail, the request's detail is %v", req.UserId, req)
		return
	} else {
		log.Printf("process %s request success, the request's detail is %v, the assign information is %v", req.UserId, req, assignVO)
	}
	log.Printf("Now is %v, updateProportion", time.Now().UnixMilli())
	contract.UpdateProportionsAsync(fairContract)
	go postToPublicChain(ethClient, fromAddr, toAddr, privateKey, req, assignVO)
}

func getAssignInfo(cspId int, vmReq VMReq) (*AssignVO, error) {
	ip := IPArr[cspId]
	url := "http://" + ip + ":" + strconv.Itoa(ReqPort) + assignPath
	reqBodyContent := ReqBody{Config: int(vmReq.Config), Duration: int(vmReq.Duration), UserId: vmReq.UserId}
	reqBodyContentJson, err := json.Marshal(reqBodyContent)
	if err != nil {
		return nil, fmt.Errorf("Marshal RequestParam fail, err:%v", err)
	}
	reqBody := strings.NewReader(string(reqBodyContentJson))
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		return nil, fmt.Errorf("Http post fail, url: %s, reqBody: %s, err: %v", url, reqBody, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Response body read fail, err: %v", err)
	}
	respBody := new(RespBody)
	err = json.Unmarshal(body, &respBody)
	log.Printf("Get response msg:\ncode: %d\nmsg: %s", respBody.Code, respBody.Msg)
	if respBody.Code == 200 {
		return &respBody.AssignVO, nil
	} else {
		return nil, nil
	}
}

func postToPublicChain(
	client *ethclient.Client,
	fromAddr common.Address,
	toAddr common.Address,
	privateKey *ecdsa.PrivateKey,
	req VMReq,
	assignVO *AssignVO,
) {
	msg, err := json.Marshal(assignVO)
	if err != nil {
		log.Fatal("Failed to convert assignVO to json: ", err)
	}

	if req.EncryptRequired {
		block, _ := pem.Decode([]byte(req.PubKey))
		if block == nil || block.Type != "PUBLIC KEY" {
			log.Fatal("Failed to decode PEM block containing public key.")
		}
		pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			log.Fatal(err)
		}
		switch pubKey.(type) {
		case *rsa.PublicKey:
			log.Println("Pub is of type RSA.")
			log.Println("The message before encryption is: ", string(msg))
			msg, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), msg)
			if err != nil {
				log.Fatal("Failed to encrypt message: ", err)
			}
			log.Println("The encrypted message is: ", string(msg))
			// msg, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey.(*rsa.PublicKey), msg, nil)
		default:
			log.Fatal("Unknown type of public key")
		}

	}
	log.Printf("Now is %v, getCollectiveSig", time.Now().UnixMilli())
	hash, sig, err := cosiUtil.GetCollectiveSignature(msg, rosterFileName)
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Fatal("Failed to get nonce: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to get gasPrice: %w", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(900000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(ResponseContractAddr)
	instance, err := myabi.NewResponse(address, client)
	if err != nil {
		log.Fatal("Failed to create instance: %w", err)
	}
	tx, err := instance.ResponseToUser(auth, toAddr, req.UserId, hash, sig, msg, req.EncryptRequired)
	if err != nil {
		log.Fatal("Failed to write to response contract: %w", err)
	}
	log.Printf("now is %v, tx sent: %v", time.Now().UnixMilli(), tx.Hash().Hex())
}
