package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

//var txs []*Transaction
//
//func conn_recv_tx(conn net.Conn) {
//	// 函数执行完之后关闭连接
//	defer conn.Close()
//	// 输出主函数传递的conn可以发现属于*TCPConn类型, *TCPConn类型那么就可以调用*TCPConn相关类型的方法, 其中可以调用read()方法读取tcp连接中的数据
//	fmt.Printf("服务端: %T\n", conn)
//	var buf [128]byte
//	// 将tcp连接读取到的数据读取到byte数组中, 返回读取到的byte的数目
//	n, err := conn.Read(buf[:])
//	if err != nil {
//		// 从客户端读取数据的过程中发生错误
//		fmt.Println("read from client failed, err:", err)
//	}
//	recvStr := string(buf[:n])
//	fmt.Println("服务端收到客户端发来的数据：", recvStr)
//	data := buf[:n]
//	txs = append(txs, DeserializeTx(data))
//	conn.Write([]byte("gotcha it"))
//}
//
//func conn_recv(addr string) []*Transaction {
//	reward := NewCoinbaseTX(addr, "")
//	txs = []*Transaction{reward}
//	listen, err := net.Listen("tcp", ":9888")
//	fmt.Printf("服务端: %T=====\n", listen)
//	if err != nil {
//		fmt.Println("listen failed, err:", err)
//	}
//	//i<打包交易个数
//	for i := 0; i < 3; i++ {
//		conn, err := listen.Accept() // 建立连接
//		fmt.Printf("当前建立了tcp连接,第%v个", i)
//		if err != nil {
//			fmt.Println("accept failed, err:", err)
//			continue
//		}
//		// 对于每一个建立的tcp连接使用go关键字开启一个goroutine处理
//		go conn_recv_tx(conn)
//	}
//	fmt.Println("结束打包")
//	return txs
//}

// 定义一个结构体
var txs []*Transaction

func recv_tx(addr string) []*Transaction {
	// 监听TCP连接
	//reward := NewCoinbaseTX(addr, "")
	//txs = []*Transaction{reward}
	txs = []*Transaction{}

	listener, err := net.Listen("tcp", ":9888")
	if err != nil {
		fmt.Println("无法监听端口:", err)
	}
	defer listener.Close()

	fmt.Println("服务器已启动，等待连接...")

	// 接受连接并处理
	//i := 0
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		fmt.Println("接受连接错误:", err)
	//		continue
	//	}
	//	//接受状态码，如果是同步区块链请求，则发送给该ip
	//	status := handleStatus(conn)
	//	if status[0:3] == "bal" {
	//		//返回余额
	//		fmt.Printf("get balance, this should not happened")
	//	}
	//	if status[0:3] == "syn" {
	//		//同步区块链
	//		SyncTx(status[3:])
	//	} else {
	//		go handleConnection(conn)
	//		i++
	//	}
	//	if i == 3 {
	//		return txs
	//	}
	//}

	//i = 每个区块交易的个数
	for i := 0; i < blockSize; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接错误:", err)
			continue
		}
		go handleConnection(conn)
	}

	//等待下一区块打包地址报名报名

	//更新给dns服务器
	send_status(dns, "upd")
	send_file(dns)
	fmt.Println("向dns更新完成")
	return txs
}

func handleConnection(conn net.Conn) {

	// 创建一个解码器和编码器
	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	//接收客户端发送的结构体
	var tx *Transaction
	err := decoder.Decode(&tx)
	if err != nil {
		fmt.Println("解码错误:", err)
		return
	}

	fmt.Println("接收到结构体:", tx)
	// 处理结构体
	txs = append(txs, tx)
	// ...

	// 发送响应到客户端
	response := "结构体已收到"
	err = encoder.Encode(response)
	if err != nil {
		fmt.Println("编码错误:", err)
		return
	}
}

//func handleStatus(conn net.Conn) string {
//	defer conn.Close()
//
//	// 创建缓冲区
//	buffer := make([]byte, 1024)
//
//	// 读取接收到的文字数据
//	n, err := conn.Read(buffer)
//	if err != nil {
//		fmt.Println("读取数据错误:", err)
//	}
//	// 提取文字内容
//	status := string(buffer[:n])
//	fmt.Println("接收到的文字:", status)
//	return status
//}
