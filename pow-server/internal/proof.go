package internal

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/big"
)

type ProofOfWork struct {
	Nonce      int64
	Complexity int64
	Data       []byte
	Target     *big.Int
}

func NewProof(complexity int64, data []byte, nonce int64) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-complexity))

	pow := &ProofOfWork{nonce, complexity, data, target}

	return pow
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}

func (pow *ProofOfWork) InitData(nonce int64) []byte {
	data := bytes.Join([][]byte{
		pow.Data,
		ToHex(nonce),
		ToHex(pow.Complexity),
	}, []byte{})

	return data
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
