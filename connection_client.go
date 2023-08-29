package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

//发送交易信息到Ip

//func conn_sendTx(ip string, data []byte) {
//	// 连接到服务端建立的tcp连接
//	conn, err := net.Dial("tcp", ip+":9888")
//	// 输出当前建Dial函数的返回值类型, 属于*net.TCPConn类型
//	fmt.Printf("客户端: %T\n", conn)
//	if err != nil {
//		// 连接的时候出现错误
//		fmt.Println("err :", err)
//		return
//	}
//	// 当函数返回的时候关闭连接
//	defer conn.Close()
//
//	_, err = conn.Write(data)
//	if err != nil {
//		return
//	}
//	buf := [512]byte{}
//	//	// 读取服务端发送的数据
//	n, err := conn.Read(buf[:])
//	if err != nil {
//		fmt.Println("recv failed, err:", err)
//		return
//	}
//	fmt.Println("客户端接收服务端发送的数据: ", string(buf[:n]))
//}

// 定义一个结构体

func send_tx(ip string, tx *Transaction) {
	// 建立TCP连接
	conn, err := net.Dial("tcp", ip+":9888")
	if err != nil {
		fmt.Println("无法建立TCP连接:", err)
		return
	}
	defer conn.Close()

	// 创建一个编码器和解码器
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	// 创建一个结构体实例

	// 发送结构体到服务器
	err = encoder.Encode(tx)
	if err != nil {
		fmt.Println("编码错误:", err)
		return
	}

	// 接收服务器的响应
	var response string
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("解码错误:", err)
		return
	}

	fmt.Println("服务器响应:", response)
}
