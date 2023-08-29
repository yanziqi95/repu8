package main

import (
	"fmt"
	"log"
	"time"
)

func (cli *CLI) pack(addr string) {
	if !ValidateAddress(addr) {
		log.Panic("ERROR: 发送地址非法")
	}

	bc := NewBlockchain() //打开数据库，读取区块链并构建区块链实例
	defer bc.Db.Close()   //转账完毕，关闭数据库
	//txs := conn_recv(addr)

	//打开上一个区块得nonce值作为随机种子
	bci := bc.Iterator()
	lastBlock := bci.Next()
	nonce := lastBlock.Nonce
	fmt.Println("nonce的值为", nonce)
	//开启一个主进程处理交易数据
	//ctx, cancel := context.WithCancel(context.Background())
	go func() {
		txs := recv_tx(addr)
		fmt.Println(txs)
		electors := recvElection(bc, nonce)
		fmt.Println(electors)
		bc.MineBlock(txs, electors)

		//	cancel()
	}()
	//
	//go func() {
	//	select {
	//	case <-ctx.Done():
	//		// 当第一个goroutine通知停止时，这里可以执行清理操作
	//		fmt.Println("Second goroutine stopped")
	//		return
	//	default:
	//		// 执行一些操作
	//		recvReview()
	//		fmt.Println("Second goroutine running")
	//	}
	//}()
	//
	//go func() {
	//	select {
	//	case <-ctx.Done():
	//		// 当第一个goroutine通知停止时，这里可以执行清理操作
	//		fmt.Println("Third goroutine stopped")
	//		return
	//	default:
	//		// 执行一些操作
	//		checkReview()
	//		fmt.Println("Third goroutine running")
	//	}
	//}()

	//开启一个进程处理评论数据
	go func() {

		recvReview()
	}()

	go func() {

		checkReview()
	}()

	select {}

	fmt.Println("成功打包交易")
	time.Sleep(time.Second * 2)
}
