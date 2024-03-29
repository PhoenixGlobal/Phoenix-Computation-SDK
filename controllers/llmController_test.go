package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
	"testing"
)

func TestCallLLM(t *testing.T) {
	reqBody := common.ReqLLM{
		Prompt:       "Funniest joke ever:",
		Temperature:  0.95,
		MaxNewTokens: 200,
		UserToken:    "",
	}
	result, err := CallLLM(reqBody)
	fmt.Println(111111, string(result), err)
}

func TestCreateLLMJob(t *testing.T) {
	reqBody := common.ReqCreateLlmJob{
		JobName:   "zckLLM115",
		TokensNum: 7}
	reqToken := common.ReqGenToken{
		Email:  "xxx@gmail.com",
		Passwd: "xxxxxx"}
	token, _ := GenToken(reqToken)
	tokenMap := make(map[string]interface{})
	err := json.Unmarshal(token, &tokenMap)
	if err != nil {
		fmt.Println("err=", err)
	}
	tokenStr := tokenMap["token"].(string)

	res, _ := CreateLLMJob(reqBody, tokenStr)
	fmt.Println("res", string(res))
}

func TestQueryLLMPrice(t *testing.T) {
	reqToken := common.ReqGenToken{
		Email:  "xxx@gmail.com",
		Passwd: "xxxxxx"}
	token, _ := GenToken(reqToken)
	tokenMap := make(map[string]interface{})
	err := json.Unmarshal(token, &tokenMap)
	if err != nil {
		fmt.Println("err=", err)
	}
	tokenStr := tokenMap["token"].(string)

	resPrice, _ := QueryLLMPrice(tokenStr)
	priceMap := make(map[string]interface{})
	err = json.Unmarshal(resPrice, &priceMap)
	priceStr := priceMap["price"].(string)
	fmt.Println("price=", priceStr)
}

func TestQueryLLMBuyCount(t *testing.T) {
	tokenStr := "XXXXXXXXXXXXX"

	res, err := QueryLLMBuyCount(tokenStr)
	resultMap := make(map[string]interface{})
	fmt.Println(11111,resultMap,err)
	err = json.Unmarshal(res, &resultMap)
	dataMap:=resultMap["data"].(map[string]interface{})
	buyCount := dataMap["count"].(float64)
	fmt.Println(22222,buyCount,err)
}
