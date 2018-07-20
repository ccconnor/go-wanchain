package protocol

import (
	"bytes"
	"github.com/wanchain/go-wanchain/p2p/discover"
	"math/big"
	"time"
)

const (
	MpcCreateLockAccountLeader = iota + 0
	MpcCreateLockAccountPeer
	MpcTXSignLeader
	MpcTXSignPeer
)
const (
	StatusCode = iota + 0 // used by storeman protocol
	KeepaliveCode
	KeepaliveOkCode
	MPCError
	RequestMPC // ask for a new mpc Context
	MPCMessage // get a message for a Context
	RequestMPCNonce
	KeepaliveCycle
	NumberOfMessageCodes

	MPCDegree          = 8
	MPCTimeOut         = time.Second * 100
	ProtocolName       = "storeman"
	ProtocolVersion    = uint64(1)
	ProtocolVersionStr = "1.0"
)
const (
	MpcPrivateShare  = "MpcPrivateShare"
	MpcPrivateKey    = "MpcPrivateKey"
	MpcPublicShare   = "MpcPublicShare"
	MpcSignA         = "MpcSignA"
	MpcSignA0        = "MpcSignA0"
	MpcSignR         = "MpcSignR"
	MpcSignR0        = "MpcSignR0"
	MpcSignB         = "MpcSignB"
	MpcSignC         = "MpcSignC"
	MpcSignARSeed    = "MpcSignARSeed"
	MpcSignARResult  = "MpcSignARResult"
	MpcTxSignSeed    = "MpcTxSignSeed"
	MpcTxSignResultR = "MpcTxSignResultR"
	MpcTxSignResultV = "MpcTxSignResultV"
	MpcTxSignResult  = "MpcTxSignResult"
	MpcContextResult = "MpcContextResult"

	PublicKeyResult = "PublicKeyResult"
	MpcSignAPoint   = "MpcSignAPoint"
	MpcTxHash       = "MpcTxHash"
	MpcTransaction  = "MpcTransaction"
	MpcChainType    = "MpcChainType"
	MpcSignType     = "MpcSignType"
	MpcChainID      = "MpcChainID"
	MpcAddress      = "MpcAddress"
	MPCActoin       = "MPCActoin"
	MPCSignedFrom   = "MPCSignedFrom"
)

type PeerInfo struct {
	PeerID discover.NodeID
	Seed   uint64
}
type SliceStoremanGroup []discover.NodeID

func (s SliceStoremanGroup) Len() int {
	return len(s)
}
func (s SliceStoremanGroup) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SliceStoremanGroup) Less(i, j int) bool {
	return bytes.Compare(s[i][:], s[j][:]) < 0
}

type GetMessageInterface interface {
	HandleMessage(*StepMessage) bool
}

type StepMessage struct {
	Msgcode   uint64 //message code
	PeerID    *discover.NodeID
	Peers     *[]PeerInfo
	Data      []big.Int //message data
	BytesData [][]byte
}
type MpcMessage struct {
	ContextID uint64
	StepID    uint64
	Peers     []byte
	Data      []big.Int //message data
	BytesData [][]byte
}