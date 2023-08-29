package main

import (
	"errors"
)

func checkDistance(senderAddress string, receiverAddress string) (int, error) {
	bc := NewBlockchain()
	defer bc.Db.Close()
	bci := bc.Iterator()
	var lastBlockHeight int
	foundTransaction := false

	for {
		block := bci.Next()
		lastBlockHeight++

		for _, tx := range block.Transactions {
			// 检查交易是否包含发送者地址和接收者地址
			if ContainsAddress(tx, senderAddress) && ContainsAddress(tx, receiverAddress) {
				foundTransaction = true
				break
			}
		}

		if foundTransaction {
			return lastBlockHeight, nil
		}

		// 如果已经达到创世区块，但未找到交易，则返回错误
		if len(block.PrevBlockHash) == 0 {
			return 0, errors.New("未找到包含指定交易的区块")
		}
	}
}

// ContainsAddress 检查交易是否包含指定地址
func ContainsAddress(tx *Transaction, address string) bool {
	// 检查交易的输入
	for _, vin := range tx.Vin {
		if vin.UsesKey([]byte(address)) {
			return true
		}
	}

	// 检查交易的输出
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	for _, vout := range tx.Vout {
		if vout.IsLockedWithKey(pubKeyHash) {
			return true
		}
	}
	return false
}
