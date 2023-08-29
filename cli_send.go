package main

import (
	"fmt"
	"log"
)

// send 转账
func (cli *CLI) send(from, to string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: 发送地址非法")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: 接收地址非法")
	}

	bc := NewBlockchain() //打开数据库，读取区块链并构建区块链实例
	defer bc.Db.Close()   //转账完毕，关闭数据库

	tx := NewUTXOTransaction(from, to, amount, bc) //创建交易

	//最后一个区块
	bci := bc.Iterator()
	lastBlock := bci.Next()
	for _, ip := range lastBlock.Ip {
		currentIP := string(getIPV4())
		fmt.Printf("current ip : %s, target ip :%s", currentIP, ip)
		if ip == currentIP { //如果当前为挖矿节点,打包交易
			fmt.Println("当前节点为矿工节点")
			reward := NewCoinbaseTX(from, "")
			bc.MineBlock([]*Transaction{reward, tx}, []string{currentIP})

			//发送给dns
			//send_status("3.8.194.158", "upd")
			send_file(dns)

		} else { //发送给所有挖矿节点
			//conn_sendTx(ip, tx.Serialize())
			//send_status(currentIP, "snd")
			send_tx(ip, tx)
			fmt.Printf("正在向发送数据" + ip)
		}

	}

	//挖出包含交易的区块，上链（写入区块链数据库）
	//reward := NewCoinbaseTX(from, "")
	//txs := []*Transaction{reward, tx}
	//bc.MineBlock(txs)
	fmt.Println("成功发送金钱")
}
