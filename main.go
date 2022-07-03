package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//pos的挖矿原理
type Block struct {
	Index int
	Data string
	PreHash string
	Hash string
	Timestamp string
	Validator *Node //记录挖矿节点
}

func genesisBlock() Block  {
	var genesBlock  = Block{0, "Genesis block","","",time.Now().String(),&Node{0, 0, "dd"}}
	genesBlock.Hash = hex.EncodeToString(BlockHash(&genesBlock))
	return genesBlock
}

func BlockHash(block *Block) []byte{
	record := strconv.Itoa(block.Index) + block.Data + block.PreHash + block.Timestamp + block.Validator.Address
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hashed
}

//创建全节点类型
type Node struct {
	Tokens int //持币数量
	Days int //持币时间
	Address string //地址
}

//创建5个节点
var nodes = make([]Node, 5)
//存放节点的地址
var addr = make([]*Node, 15)

func InitNodes()  {

	nodes[0] = Node{1, 1, "0x12341"}
	nodes[1] = Node{2, 1, "0x12342"}
	nodes[2] = Node{3, 1, "0x12343"}
	nodes[3] = Node{4, 1, "0x12344"}
	nodes[4] = Node{5, 1, "0x12345"}

	cnt :=0
	for i:=0;i<5;i++ {
		for j:=0;j<nodes[i].Tokens * nodes[i].Days;j++{
			addr[cnt] = &nodes[i]
			cnt++
		}
	}

}

//采用Pos共识算法进行挖矿
func CreateNewBlock(lastBlock *Block, data string) Block{

	var newBlock Block
	newBlock.Index = lastBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.PreHash = lastBlock.Hash
	newBlock.Data = data


	//通过pos计算由那个村民挖矿
	rand.Seed(time.Now().Unix())
	//[0,15)产生0-15的随机值
	var rd =rand.Intn(15)

	//选出挖矿的旷工
	node := addr[rd]
	//设置当前区块挖矿地址为旷工
	newBlock.Validator = node
	//简单模拟 挖矿所得奖励
	node.Tokens += 1
	newBlock.Hash = hex.EncodeToString(BlockHash(&newBlock))
	return newBlock
}

func main()  {

	InitNodes()

	//创建创世区块
	var genesisBlock = genesisBlock()

	//创建新区快
	var newBlock = CreateNewBlock(&genesisBlock, "new block")

	//打印新区快信息
	fmt.Println(newBlock)
	fmt.Println(newBlock.Validator.Address)
	fmt.Println(newBlock.Validator.Tokens)

}









