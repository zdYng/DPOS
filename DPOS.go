package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Block struct {
	Index int
	TimeStamp string
	BPM int
	Hash string
	PrevHash string
	Delegate string
}

// 产生新的区块
func generateBlock(oldBlock Block, _BMP int , address string)(Block, error){
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.TimeStamp = t.String()
	newBlock.BPM = _BMP
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = createBlockHash(newBlock)
	newBlock.Delegate = address
	fmt.Println("NewBlock: ", newBlock)
	return newBlock, nil
}

// 生成区块的hash = sha256('当前区块的index序号' +  '时间戳' + '区块的BPM' + '上一个区块的hash').string()
func createBlockHash(block Block) string{
	record := string(block.Index) + block.TimeStamp + string(block.BPM) + block.PrevHash
	sha3 := sha256.New()
	sha3.Write([] byte(record))
	hash := sha3.Sum(nil)
	fmt.Println("NewHash: ",hex.EncodeToString(hash))
	return hex.EncodeToString(hash)
}

//检视区块是否合法
func isBlockValid(newBlock, oldBlock Block) bool{
	if oldBlock.Index + 1 != newBlock.Index{
		fmt.Println("失败！！index非法")
		return false
	}
	if newBlock.PrevHash != oldBlock.Hash{
		fmt.Println("失败！！PrevHash非法")
		return false
	}
	fmt.Println("合法")
	return true
}

var blockChain []Block
type Trustee struct{
	name string
	votes int
}

type trusteeList [] Trustee

func (_trusteeList trusteeList) Len() int{
	return len(_trusteeList)
}

func (_trusteeList trusteeList) Swap(i,j int){
	_trusteeList[i],_trusteeList[j] = _trusteeList[j],_trusteeList[i]
}

func (_trusteeList trusteeList) Less(i,j int) bool{
	return _trusteeList[j].votes < _trusteeList[i].votes
}

func selecTrustee()([]Trustee){
	_trusteeList := []Trustee{
		{"node1", rand.Intn(100)},
		{"node2", rand.Intn(100)},
		{"node3", rand.Intn(100)},
		{"node4", rand.Intn(100)},
		{"node5", rand.Intn(100)},
		{"node6", rand.Intn(100)},
		{"node7", rand.Intn(100)},
		{"node8", rand.Intn(100)},
		{"node9", rand.Intn(100)},
		{"node10", rand.Intn(100)},
		{"node11", rand.Intn(100)},
		{"node12", rand.Intn(100)},
	}
	sort.Sort(trusteeList(_trusteeList))
	result := _trusteeList[:5]
	_trusteeList = result[1:]
	_trusteeList = append(_trusteeList, result[0])
	fmt.Println("当前超级节点代表列表是：",_trusteeList)
	return _trusteeList
}

func main(){
	t := time.Now()
	genesisBlock := Block{0, t.String(),0,createBlockHash(Block{}),"",""}
	fmt.Println("创世块block: ", genesisBlock)
	blockChain = append(blockChain, genesisBlock)
	var trustee Trustee
	for _, trustee = range selecTrustee(){
		_BPM := rand.Intn(100)
		blockHeight := len(blockChain)
		oldBlock := blockChain[blockHeight-1]
		newBlock,err := generateBlock(oldBlock, _BPM, trustee.name)
		if err!=nil{
			fmt.Println("新生成区块失败：",err)
			continue
		}
		if isBlockValid(newBlock, oldBlock){
			blockChain = append(blockChain, newBlock)
			fmt.Println("当前操作区块节点为：", trustee.name)
			fmt.Println("当前区块数：", len(blockChain))
			fmt.Println("当前区块信息：", blockChain[len(blockChain)-1])
		}
	}
}



