package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type GetBlockReply struct {
	Number     string `json:"number"`
	Timestamp  string `json:"timestamp"`
	ParentHash string `json:"parentHash"`
	Hash       string `json:"hash"`

	Nonce string `json:"nonce"`
	Size  string `json:"size"`
	Miner string `json:"miner"`

	Difficulty      string `json:"difficulty"`
	TotalDifficulty string `json:"totalDifficulty"`

	GasLimit string `json:"gasLimit"`
	GasUsed  string `json:"gasUsed"`

	Uncles []string `json:"uncles"`

	Transactions []Tx `json:"transactions"`
}

type Tx struct {
	BlockNumber      string `json:"blockNumber"`
	BlockHash        string `json:"blockHash"`
	TransactionIndex string `json:"transactionIndex"`
	TransactionHash  string `json:"hash"`

	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`

	Input string `json:"input"`

	Gas      string `json:"gas"`
	GasPrice string `json:"gasprice"`
	Nonce    string `json:"nonce"`
}

type JSONRpcResp struct {
	Id     *json.RawMessage       `json:"id"`
	Result *json.RawMessage       `json:"result"`
	Error  map[string]interface{} `json:"error"`
}

func main() {

	HttpClient = CreateHTTPClient(30, "")

	var params []interface{}
	jsonReq := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "method",
		"params":  params,
		"id":      0,
	}
	data, _ := json.Marshal(jsonReq)

	setterFunc := func(req *http.Request) {
		req.Header.Set("Content-Length", strconv.Itoa(len(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
	}

	Req("POST", "baseUrl", "url", data, setterFunc)

	// var rpcResp *JSONRpcResp
	// err = json.NewDecoder(resp.Body).Decode(&rpcResp)
	// if err != nil {
	// 	return nil, err
	// }
	// if rpcResp.Error != nil {
	// 	return nil, errors.New(rpcResp.Error["message"].(string))
	// }

	// if rpcResp.Result != nil {
	// 	var reply *GetBlockReply
	// 	err = json.Unmarshal(*rpcResp.Result, &reply)
	// 	return reply, err
	// }

}
