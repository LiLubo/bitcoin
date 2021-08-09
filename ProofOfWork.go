package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 1.定义一个工作量证明的结构 ProofOfWork
type ProofOfWork struct {
	block *Block // block

	// 存储哈希值，内置一了些方法
	// Cmp: 比较方法
	// SetBytes: 可以将 bytes 转成 big.int 类型
	// SetString: 可以把 string 转成 big.int 类型
	target *big.Int // 目标值,系统提供的，是固定的
}

// 2.创建 POW 的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	// 难度值，应该是推导出来的，此处简化先固定
	// 0000100000000000000000000000000000000000000000000000000000000000

	// 16 进制格式的字符串
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"

	var bigIntTmp big.Int
	bigIntTmp.SetString(targetStr, 16)

	pow.target = &bigIntTmp

	return &pow
}

// 3.不断计算 hash 的函数, 获取挖矿的 Nonce,同时返回区块的哈希
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	// 声明随机变量
	var nonce uint64
	// 声明哈希数组
	var hash [32]byte

	for {
		// 拼接 Block 和 Nonce
		data := pow.prepareData(nonce)

		// c.sha256
		hash /* 数组类型 */ = sha256.Sum256(data)

		// d.与难度值比较
		// 将 hash(数组类型) 转成 big.int,然后与 pow.target 进行比较，需要引入局部变量
		var bigIntTmp big.Int
		bigIntTmp.SetBytes(hash[:])

		// 比较
		/*
			func (x *Int) Cmp(y *Int) (r int)
			Cmp compares x and y and returns:
			-1 if x <  y
			0  if x == y
			+1 if x >  y

		*/
		result := bigIntTmp.Cmp(pow.target)

		if result == -1 {
			// Ⅰ：若哈希值小于难度值，挖矿成功，退出
			fmt.Printf("挖矿成功！Nonce: %v, hash: %v\n", nonce, hash)
			break
		} else {
			// Ⅱ: 若哈希值大于难度值，nonce ++
			nonce++
		}
	}

	return hash[:], nonce
}

// 4. 提供一个校验函数
//     - `IsValid()`

func (pow *ProofOfWork) prepareData(nonce uint64) []byte {
	// a.获取 block 数据
	block := pow.block

	// b.拼接 nonce
	tmp := [][]byte{
		uintToByte(block.Version),
		block.PrevBlockHash,
		block.MerKleRoot,
		uintToByte(block.TimeStamp),
		uintToByte(block.Difficulity),
		block.Data,
		uintToByte(nonce),
	}

	return bytes.Join(tmp, []byte{})
}
