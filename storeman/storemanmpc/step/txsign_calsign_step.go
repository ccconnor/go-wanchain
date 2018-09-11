package step

import (
	"bytes"
	"crypto/ecdsa"
	"github.com/wanchain/go-wanchain/common"
	"github.com/wanchain/go-wanchain/crypto"
	"github.com/wanchain/go-wanchain/log"
	mpcprotocol "github.com/wanchain/go-wanchain/storeman/storemanmpc/protocol"
	mpcsyslog "github.com/wanchain/go-wanchain/storeman/syslog"
	"math/big"
)

type TxSign_CalSignStep struct {
	TXSign_Lagrange_Step
}

func CreateTxSign_CalSignStep(peers *[]mpcprotocol.PeerInfo, resultKey string) *TxSign_CalSignStep {
	mpc := &TxSign_CalSignStep{*CreateTXSign_Lagrange_Step(peers, mpcprotocol.MpcTxSignSeed, resultKey)}
	return mpc
}

func (txStep *TxSign_CalSignStep) InitStep(result mpcprotocol.MpcResultInterface) error {
	ar, err := result.GetValue(mpcprotocol.MpcSignARResult)
	if err != nil {
		mpcsyslog.Err("TxSign_CalSignStep.InitStep, GetValue fail. key:%s", mpcprotocol.MpcSignARResult)
		return err
	}

	aPoint, err := result.GetValue(mpcprotocol.MpcSignAPoint)
	if err != nil {
		mpcsyslog.Err("TxSign_CalSignStep.InitStep, GetValue fail. key:%s", mpcprotocol.MpcSignAPoint)
		return err
	}

	r, err := result.GetValue(mpcprotocol.MpcSignR)
	if err != nil {
		mpcsyslog.Err("TxSign_CalSignStep.InitStep, GetValue fail. key:%s", mpcprotocol.MpcSignR)
		return err
	}

	c, err := result.GetValue(mpcprotocol.MpcSignC)
	if err != nil {
		mpcsyslog.Err("TxSign_CalSignStep.InitStep, GetValue fail. key:%s", mpcprotocol.MpcSignC)
		return err
	}

	txHash, err := result.GetValue(mpcprotocol.MpcTxHash)
	if err != nil {
		mpcsyslog.Err("TxSign_CalSignStep.InitStep, GetValue fail. key:%s", mpcprotocol.MpcTxHash)
		return err
	}

	privateKey, err := result.GetValue(mpcprotocol.MpcPrivateShare)
	if err != nil {
		mpcsyslog.Err("TxSign_CalSignStep.InitStep, GetValue fail. key:%s", mpcprotocol.MpcPrivateShare)
		return err
	}

	arInv := ar[0]
	arInv.ModInverse(&arInv, crypto.Secp256k1_N)
	invRPoint := new(ecdsa.PublicKey)
	invRPoint.Curve = crypto.S256()
	invRPoint.X, invRPoint.Y = crypto.S256().ScalarMult(&aPoint[0], &aPoint[1], arInv.Bytes())
	if invRPoint.X == nil || invRPoint.Y == nil {
		mpcsyslog.Err("TxSign_CalSignStep.InitStep, invalid r point")
		return mpcprotocol.ErrPointZero
	}

	log.Debug("calsign", "x", invRPoint.X.String(), "y", invRPoint.Y.String())
	SignSeed := new(big.Int).Set(invRPoint.X)
	SignSeed.Mod(SignSeed, crypto.Secp256k1_N)
	var v int64
	if invRPoint.X.Cmp(SignSeed) == 0 {
		v = 0
	} else {
		v = 2
	}

	invRPoint.Y.Mod(invRPoint.Y, big.NewInt(2))
	if invRPoint.Y.Cmp(big.NewInt(0)) != 0 {
		v |= 1
	}

	log.Debug("calsign", "v", v)
	result.SetValue(mpcprotocol.MpcTxSignResultR, []big.Int{*SignSeed})
	result.SetValue(mpcprotocol.MpcTxSignResultV, []big.Int{*big.NewInt(v)})
	SignSeed.Mul(SignSeed, &privateKey[0])
	SignSeed.Mod(SignSeed, crypto.Secp256k1_N)
	hash := txHash[0]
	SignSeed.Add(SignSeed, &hash)
	SignSeed.Mod(SignSeed, crypto.Secp256k1_N)
	SignSeed.Mul(SignSeed, &r[0])
	SignSeed.Mod(SignSeed, crypto.Secp256k1_N)
	SignSeed.Add(SignSeed, &c[0])
	SignSeed.Mod(SignSeed, crypto.Secp256k1_N)

	log.Debug("calsign", "seed", SignSeed.String())
	result.SetValue(mpcprotocol.MpcTxSignSeed, []big.Int{*SignSeed})
	return txStep.TXSign_Lagrange_Step.InitStep(result)
}

func (txStep *TxSign_CalSignStep) FinishStep(result mpcprotocol.MpcResultInterface, mpc mpcprotocol.StoremanManager) error {
	err := txStep.TXSign_Lagrange_Step.FinishStep(result, mpc)
	if err != nil {
		return err
	}

	err = mpc.SignTransaction(result)
	if err != nil {
		return err
	}

	from, err := result.GetValue(mpcprotocol.MpcAddress)
	if err != nil {
		mpcsyslog.Debug("TxSign_CalSignStep.FinishStep, GetValue fail. key:%s", mpcprotocol.MpcAddress)
		return nil
	}

	address := common.BigToAddress(&from[0])
	signedFrom, err := result.GetByteValue(mpcprotocol.MPCSignedFrom)
	if err != nil {
		mpcsyslog.Debug("TxSign_CalSignStep.FinishStep, GetValue fail. key:%s", mpcprotocol.MPCSignedFrom)
		return nil
	}

	log.Info("calsign finish. check signed from", "require", common.ToHex(address[:]), "actual", common.ToHex(signedFrom))
	mpcsyslog.Info("calsign finish. check signed from. require:%s, actual:%s", common.ToHex(address[:]), common.ToHex(signedFrom))
	if !bytes.Equal(address[:], signedFrom) {
		mpcsyslog.Err("TxSign_CalSignStep.FinishStep, unexpect signed data from address. require:%s, actual:%s", common.ToHex(address[:]), common.ToHex(signedFrom))
		return mpcprotocol.ErrFailSignRetVerify
	}

	return nil
}