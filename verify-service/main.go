package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.dedis.ch/cothority/v3/blscosi"
	"go.dedis.ch/cothority/v3/blscosi/blscosi/check"
	"go.dedis.ch/onet/v3/app"
	"log"
	"net/http"
	"os"
	"strings"
)

const GroupTomlName = "roster.toml"

type verifyReq struct {
	Msg       string `json:"Msg"`
	Hash      string `json:"Hash"`
	Signature string `json:"Signature"`
}

type verifyResult struct {
	IsValid bool   `json:"IsValid"`
	Desc    string `json:"Desc"`
}

type sigHex struct {
	Hash      string
	Signature string
}

func main() {
	router := gin.Default()
	router.Use(Cors())
	router.POST("/verify", verify)
	router.GET("/test", test)

	router.Run("localhost:8080")
}

func verify(c *gin.Context) {
	var req verifyReq
	var result verifyResult

	if err := c.BindJSON(&req); err != nil {
		return
	}

	err := verifySig(req.Msg, req.Hash, req.Signature)
	if err != nil {
		result.IsValid = false
		result.Desc = fmt.Sprint(err)
	} else {
		result.IsValid = true
		result.Desc = "Signature verified and is correct!"
	}

	c.IndentedJSON(http.StatusOK, result)
}

func test(c *gin.Context) {
	key := c.Query("key")
	block, _ := pem.Decode([]byte(key))
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("Failed to decode PEM block containing public key.")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	msg := []byte("Hello World")
	log.Println(msg)
	switch pub := pubKey.(type) {
	case *rsa.PublicKey:
		log.Println("Pub is of type RSA:", pub)
		msg, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), msg)
		log.Println(hex.EncodeToString(msg))
		// msg, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey.(*rsa.PublicKey), msg, nil)
	default:
		log.Fatal("Unknown type of public key")
	}
	c.IndentedJSON(http.StatusOK, hex.EncodeToString(msg))
}

// verifySig takes a group-definition, calls the signature
// verification and prints the result. If sigFileName is empty it
// assumes to find the standard signature in fileName.sig
func verifySig(msg string, hash string, signature string) error {
	sig := &blscosi.SignatureResponse{}
	decodeHash, err := hex.DecodeString(hash)
	if err != nil {
		return err
	}
	sig.Hash = decodeHash
	sig.Signature, err = hex.DecodeString(signature)
	if err != nil {
		return err
	}
	msgBytes, err := hex.DecodeString(strings.TrimPrefix(msg, "0x"))
	if err != nil {
		return err
	}

	fGroup, err := os.Open(GroupTomlName)
	if err != nil {
		return err
	}

	log.Println("Reading group definition")
	g, err := app.ReadGroupDescToml(fGroup)
	if err != nil {
		return err
	}

	log.Println("Verifying signature: ", msg)
	if err := check.VerifySignatureHash(msgBytes, sig, g.Roster); err != nil {
		return err
	}

	return nil
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
