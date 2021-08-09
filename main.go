package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("I'm Li Lubo")
	bc.AddBlock("I'm Cai Jinxiao")

	// 遍历打印所有的区块
	for i, block := range bc.Blocks {
		fmt.Printf("========================== %d ==========================\n", i) // 打印区块高度
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
	}

}
