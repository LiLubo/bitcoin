package main

import (
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
