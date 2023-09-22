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
	reqBody := common.ReqCreateLlmJob{JobName: "LLM test11", Count: 1}
	reqToken := common.ReqGenToken{Email: "515588290@qq.com", Passwd: "111"}
	token, _ := GenToken(reqToken)
	fmt.Println("tk=", string(token))
	res, _ := CreateLLMJob(reqBody, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjUxNTU4ODI5MEBxcS5jb20iLCJleHAiOjE2OTUzNjMzODl9.f3uO27mpkAsG-oMgzc4pDuE_sGGrn4UMO7sdH_-xtUU")
	fmt.Println("res", string(res))
}
