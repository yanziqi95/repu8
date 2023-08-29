package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//从其他节点同步区块链
func handler_sendFile(conn net.Conn) error {
	file, err := os.Open(dbFile)
	if err != nil {
		return fmt.Errorf("Error opening Bolt file: %s", err)
	}
	defer file.Close()

	_, err = conn.Write([]byte("upd"))
	if err != nil {
		return fmt.Errorf("Error sending status code: %s", err)
	}

	_, err = io.Copy(conn, file)
	if err != nil {
		return fmt.Errorf("Error sending Bolt file: %s", err)
	}

	return nil
}

//挖矿节点间确认区块链并同步

func send_file(ip string) {

	// 建立TCP连接
	conn, err := net.Dial("tcp", ip+":9888")
	if err != nil {
		fmt.Println("无法建立TCP连接:", err)
		return
	}
	defer conn.Close()

	err = handler_sendFile(conn)
	if err != nil {
		fmt.Println("Error sending Bolt file:", err)
		return
	}
	// 创建缓冲区
	//buffer := make([]byte, 1024)

	//逐块读取文件并发送
	//for {
	//	// 读取文件数据到缓冲区
	//	n, err := file.Read(buffer)
	//	if err != nil {
	//		if err != io.EOF {
	//			fmt.Println("读取文件错误:", err)
	//		}
	//		break
	//	}
	//
	//	// 发送数据块到服务器
	//	_, err = conn.Write(buffer[:n])
	//	if err != nil {
	//		fmt.Println("发送数据错误:", err)
	//		return
	//	}
	//}

	fmt.Println("文件发送完成")
}
