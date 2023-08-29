package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkReview() {
	//从轻节点接受评论检查请求

	// 创建TCP监听
	listener, err := net.Listen("tcp", ":9886")
	if err != nil {
		fmt.Println("无法创建TCP监听:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器已启动，监听地址:", 9886)

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("客户端连接错误:", err)
			continue
		}

		// 处理客户端连接
		go handleCheck(conn)
	}
}

func handleCheck(conn net.Conn) {
	defer conn.Close()

	// 读取客户端发送的数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("读取数据错误:", err)
		return
	}
	target := string(buffer[:n])

	fmt.Printf("Server is checking the hash of seller : %s", target)
	dbPath := "./reviewHash/" + target + ".txt"
	if _, err := os.Stat(dbPath); !os.IsNotExist(err) {
		// 读取文件内容
		existingContent, err := ioutil.ReadFile(dbPath)
		_, err = conn.Write(existingContent)
		if err != nil {
			fmt.Println("Error sending response:", err.Error())
			return
		}
	} else {
		fmt.Println("文件不存在")

	}

}