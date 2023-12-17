package fsociety

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0x71797875
	TestnetMagic wire.BitcoinNet = 0x4e455552
)

// chain parameters
var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{35}
	MainNetParams.ScriptHashAddrID = []byte{95}

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{127}
	TestNetParams.ScriptHashAddrID = []byte{196}
}

// FsocietyParser handle
type FsocietyParser struct {
	*btc.BitcoinLikeParser
	baseparser *bchain.BaseParser
}

// NewFsocietyParser returns new FsocietyParser instance
func NewFsocietyParser(params *chaincfg.Params, c *btc.Configuration) *FsocietyParser {
	return &FsocietyParser{
		BitcoinLikeParser: btc.NewBitcoinLikeParser(params, c),
		baseparser:        &bchain.BaseParser{},
	}
}

// GetChainParams contains network parameters
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}

// PackTx packs transaction to byte array using protobuf
func (p *FsocietyParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *FsocietyParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
