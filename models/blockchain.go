package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type response struct {
	id     int32  `json:"id"`
	result string `json:"result"`
	error  string `json:"error"`
}

var (
	Jrpc_Url = "http://124.223.171.19:8801" //对应于chain33.toml中的配置jrpcBindAddr
	privKey  = "3990969DF92A5914F7B71EEB9A4E58D6E255F32BF042FEA5318FC8B3D50EE6E8"
)

//func main() {
//	if len(os.Args) < 3 {
//		return
//	}
//	key := "key"
//	value := "value1,value2,value3"
//	privKey := "3990969DF92A5914F7B71EEB9A4E58D6E255F32BF042FEA5318FC8B3D50EE6E8"
//	_ = transfer(key, value, privKey)
//	queryTx("0x51ed9840f8d7ae5959b3f74403bf83a55e5852d4c21d34796c6bc99c31b3905f")
//	fmt.Printf("ok")
//}

// 获得json rpc服务的url信息
func getJrpc() string {
	return Jrpc_Url
}

// 格式化打印json结构
func printJson(data []byte) {
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(data), "", "    ")
	fmt.Println(str.String())
}

// 解析JSON响应，获得result字段或错误信息
func parseJsonResult(data []byte) (string, error) {
	var txdata = make(map[string]interface{})
	err := json.Unmarshal(data, &txdata)
	if err != nil {
		fmt.Println("err:", err.Error())
		return "", err
	}
	if hextx, ok := txdata["result"]; ok {
		return hextx.(string), nil
	}
	return "", fmt.Errorf("not have result!")
}

// 创建原始交易
func createTx(Key string, Value string) (string, error) {
	//按照“RPC接口>系统接口>交易接口”中的“构造交易 CreateRawTransaction”定义，来构造JSON RPC消息内容
	poststr := fmt.Sprintf(`{"jsonrpc":"2.0","id":0,"method":"Chain33.CreateTransaction","params":[{"execer":"storage","actionName":"ContentStorage", "payload":{"key":"%v","value":"%v"}}]}`, Key, Value)
	fmt.Println("---create tx request:")
	printJson([]byte(poststr))
	return constructTx(poststr)
}

// 向chain33节点调用JSON RPC接口，构造交易。
func constructTx(poststr string) (string, error) {
	resp, err := http.Post(getJrpc(), "application/json", bytes.NewBufferString(poststr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("---create tx response:")
	printJson(data)
	return parseJsonResult(data)
}

// 对原始交易，使用地址及私钥进行签名，得到签名后的交易。
func SignRawTx(rawTx, privKey, expire string, index int32) (string, error) {
	client := http.DefaultClient
	//按照“RPC接口>系统接口>交易接口”中的“交易签名 SignRawTx”定义，来构造JSON RPC消息内容
	postdata := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"Chain33.SignRawTx","params":[{"privKey":"%v","txHex":"%v","expire":"%v"}]}`, privKey, rawTx, expire)
	//向chain33节点调用JSON RPC接口，签名交易。
	req, err := http.NewRequest("post", getJrpc(), strings.NewReader(postdata))
	if err != nil {
		fmt.Println("err:", err.Error())
		return "", err
	}
	fmt.Println("---sign tx request:")
	printJson([]byte(postdata))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err:", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err.Error())
		return "", err
	}
	fmt.Println("---sign tx response:")
	printJson(data)
	return parseJsonResult(data)
}

// 向区块链提交交易
func SendTx(hexTx string) string {
	//按照“RPC接口>系统接口>交易接口”中的“发送交易 SendTransaction”定义，来构造JSON RPC消息内容
	poststr := fmt.Sprintf(`{"jsonrpc":"2.0","id":2,"method":"Chain33.SendTransaction","params":[{"data":"%v"}]}`, hexTx)
	fmt.Println("---send tx request:")
	printJson([]byte(poststr))
	//向chain33节点调用JSON RPC接口，发送交易。交易将被提交到区块链网络中。
	resp, err := http.Post(getJrpc(), "application/json", bytes.NewBufferString(poststr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("aaa%s\n", string(b[15]))
	fmt.Printf("%q\n", strings.Split(string(b), "\""))
	splits := strings.Split(string(b), "\"")
	fmt.Printf(splits[5] + "\n")
	var res response
	e := json.Unmarshal(b, &res)
	//res.result = "0x5c0f40ff01ebb11a54210322d4d4eafa7383a6ae683352cd5144fd15be4a13db"
	if e != nil {
		fmt.Println(e)
	}
	if err != nil {
		panic(err)
	}
	//结果中包括本交易的hash值，后续可以根据该值查询交易详情。
	fmt.Print("---send tx response:")
	printJson(b)
	return splits[5]
}

// 从一个地址转账到另一个地址
func transfer(Key, Value string) string {
	hexCreateTx, err := createTx(Key, Value) //构造原始交易。
	fmt.Println("---created tx:")
	fmt.Println(hexCreateTx)
	if err != nil {
		fmt.Println("testCreateGuess have a err:", err.Error())
		return "error"
	}
	hexTx, err := SignRawTx(hexCreateTx, privKey, "3600s", 0) //签名交易。
	if err != nil {
		fmt.Println("SignRawTx have a err:", err.Error())
		return "error"
	}
	return SendTx(hexTx) //发送交易
}

// 根据交易hash查询交易详情
// todo：返回查询值
func queryTx(txHash string) []byte {
	//按照“RPC接口>系统接口>交易接口”中的“根据哈希查询交易信息 QueryTransaction”定义，来构造JSON RPC消息内容
	poststr := fmt.Sprintf(`{"jsonrpc":"2.0","id":0,"method":"Chain33.QueryTransaction","params":[{"hash":"%v"}]}`, txHash)
	fmt.Println("---query tx request:")
	printJson([]byte(poststr))
	//向chain33节点调用JSON RPC接口，查询交易。
	resp, err := http.Post(getJrpc(), "application/json", bytes.NewBufferString(poststr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("---query tx response:")
	printJson(data)
	return data
}
