package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// 1.定义结构
type Block struct {
	Version       uint64 // 区块版本号
	PrevBlockHash []byte // 前区块哈希
	MerKleRoot    []byte // 默克尔根
	TimeStamp     uint64 // 时间戳，从1970年1月1日至今的秒数
	Difficulity   uint64 // 挖矿难度
	Nonce         uint64 // 随机数
	Data          []byte // 数据
	Hash          []byte // 当前区块哈希，区块中本不存在的字段，为了方便而添加
}

// 2.创建区块，对Block的每一个字段进行填充
func NewBlock(data string, prevBlockHash []byte) *Block {

	block := Block{
		Version:       00,
		PrevBlockHash: prevBlockHash,
		MerKleRoot:    []byte{},
		TimeStamp:     uint64(time.Now().Unix()),
		Difficulity:   10, // 后续会调整
		Nonce:         10, // 后续会调整
		Hash:          []byte{},
		Data:          []byte(data),
	}

	block.setHash()

	return &block
}

// 3.生成哈希(无随机值)
func (block *Block) setHash() {
	var data []byte
	// uintToByte() 将数字转换成 []byte{},在utils.go中实现
	// data = append(data, uintToByte(block.Version)...)
	// data = append(data, block.PrevBlockHash...)
	// data = append(data, block.MerKleRoot...)
	// data = append(data, uintToByte(block.TimeStamp)...)
	// data = append(data, uintToByte(block.Difficulity)...)
	// data = append(data, uintToByte(block.Nonce)...)
	// data = append(data, block.Data...)

	tmp := [][]byte{
		uintToByte(block.Version),
		block.PrevBlockHash,
		block.MerKleRoot,
		uintToByte(block.TimeStamp),
		uintToByte(block.Difficulity),
		uintToByte(block.Nonce),
		block.Data,
	}

	data = bytes.Join(tmp, []byte{})

	hash /* [32]byte */ := sha256.Sum256(data)
	block.Hash = hash[:]
}
