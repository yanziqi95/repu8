package main

import "math"

func ratings(sender string, receiver string, p int) (int, int) {
	dis, _ := checkDistance(sender, receiver)

	bc := NewBlockchain() //打开数据库，读取区块链并构建区块链实例
	defer bc.Db.Close()
	bal := getElectorBalance(bc, sender)

	var rating int
	var ratingReward int
	if dis <= disThreshold {
		rating = (disThreshold / dis) * bal * p
		ratingReward = disThreshold / dis
	} else {
		rating = (1 / 2) * bal * p
		ratingReward = 1 / 2
	}
	return rating, ratingReward
}

func reviewReward(rating int, invest int) int {
	reward := rating * invest * roi
	return reward
}

//punishment is an index [0,1], should multiply purchased tokens
func punishment(address string) float64 {
	//找到大笔交易的数量和距离
	var dis int
	//计算惩罚数量
	var index float64
	if dis <= 10 {
		index = math.Exp(-0.2 * float64(dis))
	} else {
		index = 0
	}
	return index
}
