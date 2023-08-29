package main

import (
	"fmt"
	"net"
)

func send_status(ip string, status string) {
	// 建立TCP连接
	conn, err := net.Dial("tcp", ip+":9888")
	if err != nil {
		fmt.Println("无法建立TCP连接:", err)
		return
	}
	defer conn.Close()

	// 待发送的文字
	message := status

	// 发送文字数据
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("发送数据错误:", err)
		return
	}

	fmt.Println("文字发送完成")
}
