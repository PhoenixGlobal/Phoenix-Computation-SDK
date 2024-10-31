package controllers

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
)

func TestCallLLM3(t *testing.T) {
	reqBody := common.ReqLLM3{
		Prompt:       "Funniest joke ever:",
		MaxNewTokens: 200,
		UserToken:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MDg2NjE4NjR9.0IaBfwOrjEo5XCkrvPcJJ4dgdzDf6p5hxSFHvx1GygQ",
	}
	result, err := CallLLM3(reqBody)
	fmt.Println(111111, string(result), err)
}

func TestVerticalLLM(t *testing.T) {
	reqBody := common.ReqLLM3{
		Prompt:       "Write a quicksort algorithm in Golang.",
		MaxNewTokens: 200,
		UserToken:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MDg2NjE4NjR9.0IaBfwOrjEo5XCkrvPcJJ4dgdzDf6p5hxSFHvx1GygQ",
	}
	result, err := VerticalLLM(reqBody)
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

	res, _ := BuyLLMToken(reqBody, tokenStr)
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

type ReqGenVideo struct {
	Uidname  string `json:"uidname"`
	Prompt   string `json:"prompt"`
	Numsteps int    `json:"numsteps"`
}

func TestGenMotion(t *testing.T) {
	reqBody := common.ReqTxt2Motion{
		Prompt:        "iron man skiing on steep slope",
		NegativePromt: "blurry",
		UserToken:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MDY2MTAyNTN9.k7nc5Ev_YUFKd5ciGSQZZ0Y3yfNQ6_gy9vY8WuBBtOs",
	}
	result, err := Txt2Motion(reqBody)
	fmt.Println(result, err)
}

func TestGenSDXl(t *testing.T) {
	reqBody := common.ReqGenImgSDXL{
		Prompt:        "iron man skiing on steep slope",
		NegativePromt: "blurry",
		UserToken:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MDg5NDYzNzJ9._5nm-dNuxTRmvdAESbQW7HlXPvGmc3g2EoFz2d5NL_0",
	}
	result, err := GenImgSDXL(reqBody)
	fmt.Println(result, err)
}
