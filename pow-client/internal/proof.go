package internal

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/big"
)

type ProofOfWork struct {
	Nonce      int64
	Complexity int64
	Data       []byte
	Target     *big.Int
}

func DoWork(complexity int64, data string) (int64, []byte) {
	pow := NewProof(complexity, []byte(data), 1)
	nonce, hash := pow.Run()

	return nonce, hash
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var intHash big.Int
	var hash [32]byte
	var nonce int64
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func (pow *ProofOfWork) InitData(nonce int64) []byte {
	data := bytes.Join([][]byte{
		pow.Data,
		ToHex(nonce),
		ToHex(pow.Complexity),
	}, []byte{})

	return data
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
