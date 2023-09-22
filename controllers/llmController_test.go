package controllers

import (
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
	reqBody := common.ReqCreateLlmJob{JobName: "zckLLM", TokensNum: 7}
	reqToken := common.ReqGenToken{Email: "xxx@xx.com", Passwd: "xxx"}
	token, _ := GenToken(reqToken)
	fmt.Println("tk=", string(token))
	res, _ := CreateLLMJob(reqBody, "xxx")
	fmt.Println("res", string(res))
}
