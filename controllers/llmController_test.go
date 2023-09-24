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
		Email:  "515588290@qq.com",
		Passwd: "111"}
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
		Email:  "515588290@qq.com",
		Passwd: "111"}
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
