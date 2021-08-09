package main

import "fmt"

const blockInterval = 21 // 区块衰减区间， 万

func totalBtc() {
	total := 0.0   // 总量
	reward := 50.0 // 最初奖励

	for reward > 0 {
		amount := reward * blockInterval // 在21万个块中产生的比特币数量
		total += amount

		reward *= 0.5 // 比特币奖励衰减一半，除法效率低
	}

	fmt.Printf("total Bitcoin: %f 万个\n", total) // 2100.000000 万个
}
