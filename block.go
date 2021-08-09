package main

import (
	"crypto/sha256"
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
