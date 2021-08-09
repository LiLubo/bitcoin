package main

// 4.引入区块链,使用Block数组模拟
type BlockChain struct {
	Blocks []*Block
}

// 创建区块链方法
func NewBlockChain() *BlockChain {
	// 在创建的时候添加一个区块， 创世块
	genesisBlock := NewBlock(
		"hello world",
		[]byte{0x0000000000000000},
	)

	bc := BlockChain{[]*Block{genesisBlock}}

	return &bc
}

// 5.添加区块
func (bc *BlockChain) AddBlock(data string) {
	// a.创建区块
	// 获取 bc.Blocks的最后一个区块的Hash值就是当前区块的 PrevBlockHash
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(data, lastBlock.Hash)

	// b.添加区块至bc.Blocks数组中
	bc.Blocks = append(bc.Blocks, block)
}
