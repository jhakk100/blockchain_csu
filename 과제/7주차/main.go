package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.blocks); i++ {
		if !bytes.Equal(
			bc.blocks[i].PrevBlockHash,
			bc.blocks[i-1].Hash,
		) {
			return false
		}
	}
	return true
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	bc.AddBlock("Send 3 BTC to Alice")

	//bc.blocks[1].Data = []byte("Send 9 BTC to Ivan")
	//bc.blocks[1].SetHash()

	if bc.IsValid() {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

// 미션 1: 그대로 실행하기, 지문이 바뀌었는가?
// 바뀝니다
// 미션 2: 블록 하나 더 붙이기, prev hash가 앞 블록의 hash와 같나?
// 같습니다
// 미션 3: 몰래 고쳐보기, 데이터 바꾼 후 실행하면 SetHash()를 호출하는가?
// 이전 hash를 가져오기 때문에, 변조를 한 들 숨길 수 없음
// 미션 4: 검사기 만들기
// 이전 값과 이후값만을 비교하여 출력을 하는 코드 작성, 아마 이정도면 되지 않을까 함
