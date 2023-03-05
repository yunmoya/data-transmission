package cosiUtil

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go.dedis.ch/cothority/v3/blscosi"
	"go.dedis.ch/cothority/v3/blscosi/blscosi/check"
	"go.dedis.ch/onet/v3/app"
	"log"
	"os"
)

type sigHex struct {
	Hash      string
	Signature string
}

// sign takes a stream and a toml file defining the servers
func sign(msg []byte, tomlFileName string) (*blscosi.SignatureResponse, error) {
	log.Println("Starting signature")
	f, err := os.Open(tomlFileName)
	if err != nil {
		return nil, err
	}
	g, err := app.ReadGroupDescToml(f)
	if err != nil {
		return nil, err
	}
	if len(g.Roster.List) <= 0 {
		return nil, fmt.Errorf("Empty or invalid blscosi group file: %s", tomlFileName)
	}

	log.Println("Sending signature to", g.Roster)
	return check.SignStatement(msg, g.Roster)
}

func GetCollectiveSignature(msg []byte, rosterFileName string) (hash string, sig string, err error) {
	resp, err := sign([]byte(msg), rosterFileName)
	if err != nil {
		return "", "", fmt.Errorf("Failed to sign msg: %w", err)
	}

	return hex.EncodeToString(resp.Hash), hex.EncodeToString(resp.Signature), nil
}

// writeSigAsJSON - writes the JSON out to a file
func WriteSigAsJSON(res *blscosi.SignatureResponse) (string, error) {
	b, err := json.Marshal(sigHex{
		Hash:      hex.EncodeToString(res.Hash),
		Signature: hex.EncodeToString(res.Signature)},
	)

	if err != nil {
		return "", fmt.Errorf("Couldn't encode signature: %s", err.Error())
	}

	return string(b), nil
}
