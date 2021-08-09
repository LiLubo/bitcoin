package main

import (
	"crypto/sha256"
	"fmt"
)

// 1.定义结构
type Block struct {
	PrevBlockHash []byte // a.前区块哈希
	Hash          []byte // b.当前区块哈希
	Data          []byte // c.数据
}

// 2.创建区块，对Block的每一个字段进行填充
func NewBlock(data string, prevBlockHash []byte) *Block {

	block := Block{
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
	}

	block.setHash()

	return &block
}

// 3.生成哈希(无随机值)
func (block *Block) setHash() {
	var data []byte
	data = append(data, block.PrevBlockHash...)
	data = append(data, block.Data...)

	hash /* [32]byte */ := sha256.Sum256(data)
	block.Hash = hash[:]
}

// 4.引入区块链
// 5.添加区块
// 6.重构代码

func main() {
	block := NewBlock(
		"hello world",
		[]byte{0x0000000000000000},
	)

	fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("Data: %s\n", block.Data)
}
