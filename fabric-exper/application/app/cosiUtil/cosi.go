package cosiUtil

import (
	pb "app/gocosi"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go.dedis.ch/cothority/v3/blscosi"
	"go.dedis.ch/cothority/v3/blscosi/blscosi/check"
	"go.dedis.ch/onet/v3/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

type sigHex struct {
	Hash      string
	Signature string
}

// signWithBLSCosi takes a stream and a toml file defining the servers
func signWithBLSCosi(msg []byte, tomlFileName string) (*blscosi.SignatureResponse, error) {
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

func signWithGoCosi(msg string, addr string) (*pb.NewMsgResp, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGocosiRPCClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	r, err := c.NewMsg(ctx, &pb.NewMsgReq{Msg: msg})
	if err != nil {
		return nil, fmt.Errorf("could not send msg: %v", err)
	}
	log.Printf("Msg: %v\n Signature: %v\n Signers: %v\n", msg, r.Signature, r.Signers)
	return r, nil
}

func GetCollectiveSignatureByBLSCosi(msg []byte, rosterFileName string) (hash string, sig string, err error) {
	resp, err := signWithBLSCosi([]byte(msg), rosterFileName)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign msg: %w", err)
	}

	return hex.EncodeToString(resp.Hash), hex.EncodeToString(resp.Signature), nil
}

func GetCollectiveSignatureByGoCosi(msgStr string, addr string) (string, []int32, error) {
	resp, err := signWithGoCosi(msgStr, addr)
	if err != nil {
		return "", nil, fmt.Errorf("failed to sign msg: %w", err)
	}

	return resp.Signature, resp.Signers, nil
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
