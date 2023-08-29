package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func submitReivew(id string, target string, review string) {
	//数据hash
	h := sha256.New()
	h.Write([]byte(review))
	hr := h.Sum(nil)
	fmt.Printf("%s is giving review to %s:%x\n", id, target, hr)
	//数据上链

	//数据发送给服务器

}

// 创建一个结构体来表示JSON数据的格式
type dataJSON struct {
	Sender     string `json:"sender"`
	Seller     string `json:"seller"`
	Comment    string `json:"comment"`
	Investment string `json:"investment"`
	Polarity   string `json:"polarity"`
	Ratings    int    `json:"ratings"`
}

func recvReview() {
	//从轻节点接受评论

	// 创建TCP监听
	listener, err := net.Listen("tcp", ":9887")
	if err != nil {
		fmt.Println("无法创建TCP监听:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器已启动，监听地址:", 9887)

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("客户端连接错误:", err)
			continue
		}

		// 处理客户端连接
		go handleClient(conn)

	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// 读取客户端发送的数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("读取数据错误:", err)
		//return
	}

	// 解析，转发JSON数据
	//var data dataJSON

	data := buffer[:n]

	//err = json.Unmarshal(buffer[:n], &data)
	//if err != nil {
	//	fmt.Println("解析JSON数据错误:", err)
	//	return
	//}

	// 在服务器端处理数据
	fmt.Printf("收到JSON数据： %+v\n", data)

	// 可以在这里添加进一步处理逻辑
	//记录评论哈希
	var jsonData dataJSON
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		// 处理解码错误
		fmt.Println("json解码错误")
		return
	}
	seller := jsonData.Seller
	comment := jsonData.Comment
	//sender :=jsonData.Sender
	//investment :=jsonData.Investment
	hashReview(seller, comment)
	//发送http请求到网页
	url := "http://repustation.000webhostapp.com/index.php"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("HTTP POST error:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful!")
	} else {
		fmt.Println("Request failed. Status code:", resp.StatusCode)
	}
	//return sender, investment
}

func hashReview(seller string, comment string) {
	dbPath := "./reviewHash/" + seller + ".txt"
	if _, err := os.Stat(dbPath); !os.IsNotExist(err) {
		// 读取文件内容
		existingContent, err := ioutil.ReadFile(dbPath)
		if err != nil {
			fmt.Println("无法读取文件:", err)
			return
		}

		// 拼接内容和comment
		updatedContent := string(existingContent) + comment

		// 计算SHA-256哈希值
		hash := sha256.Sum256([]byte(updatedContent))

		// 将哈希值写入文件
		err = ioutil.WriteFile(dbPath, []byte(hex.EncodeToString(hash[:])), os.ModePerm)
		if err != nil {
			fmt.Println("无法写入文件:", err)
			return
		}

		fmt.Println("文件已创建并插入comment的SHA-256值。")

	} else {
		// 文件不存在，创建并插入comment的SHA-256值

		// 计算SHA-256哈希值
		hash := sha256.Sum256([]byte(comment))

		// 将哈希值写入文件
		err := ioutil.WriteFile(dbPath, []byte(hex.EncodeToString(hash[:])), os.ModePerm)
		if err != nil {
			fmt.Println("无法创建文件并写入哈希值:", err)
			return
		}
		fmt.Println("文件已创建并插入comment的SHA-256值:%s。")
	}

}

//func sendReviewJson(seller string, comment string, ratings int) {
//	url := "http://repustation.000webhostapp.com/index.php" // 替换为实际的 PHP 脚本地址
//
//	// 构建要发送的 JSON 数据
//	data := map[string]interface{}{
//		"seller":  seller,
//		"comment": comment,
//		"ratings": ratings,
//	}
//
//	jsonData, err := json.Marshal(data)
//	if err != nil {
//		fmt.Println("JSON encoding error:", err)
//		return
//	}
//
//	// 发送 HTTP POST 请求
//	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
//	if err != nil {
//		fmt.Println("HTTP POST error:", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	// 处理响应
//	if resp.StatusCode == http.StatusOK {
//		fmt.Println("Request successful!")
//	} else {
//		fmt.Println("Request failed. Status code:", resp.StatusCode)
//	}
//
//}
