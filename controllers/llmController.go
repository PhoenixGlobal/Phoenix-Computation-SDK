package controllers

import (
	"encoding/json"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/util"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
)

const LlmApiURL string = "https://www.phoenix.global/sdk/llm/v1/completions"
const CreateLLMJobURL string = "https://www.phoenix.global/sdk/computation/LLM/createLLMJob"
const QueryLLMPriceURL string = "https://www.phoenix.global/sdk/computation/LLM/queryLLMPrice"
const QueryLLMCountURL string = "https://www.phoenix.global/sdk/computation/LLM/queryLLMCount"

// CallLLM  call LLM api
func CallLLM(reqBody common.ReqLLM) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPostForLLM(LlmApiURL, reqJson)
	if err != nil {
		return nil, err
	}
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(result, &resultMap)
	if err != nil {
		return nil, err
	}
	//textMapArr := resultMap["choices"].([]interface{})
	//textMap := textMapArr[0].(map[string]interface{})
	text := resultMap["text"].(string)
	useTotalTokens := resultMap["useTotalTokens"].(float64)
	buyTokens,err:=GetLLMBuyCount(reqBody.UserToken)
	if err != nil {
		return nil, err
	}
	retResult := common.ResLLM{
		Code: 200,
		Msg:  "success",
		Text: text,
		TokensBalance: buyTokens-useTotalTokens,
	}
	res, err := json.Marshal(retResult)
	if err != nil {
		return nil, err
	}
	return res, err
}

func CreateLLMJob(reqBody common.ReqCreateLlmJob, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(CreateLLMJobURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func QueryLLMPrice(token string) (json.RawMessage, error) {
	result, err := util.SendHttpGet(QueryLLMPriceURL, nil, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func QueryLLMBuyCount(token string) (json.RawMessage, error) {
	result, err := util.SendHttpPost(QueryLLMCountURL, nil, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetLLMBuyCount(token string) (count float64,e error)  {
	res, err := QueryLLMBuyCount(token)
	if err!=nil{
		e=err
		return
	}
	resultMap := make(map[string]interface{})
	e = json.Unmarshal(res, &resultMap)
	if e!=nil{
		return
	}
	dataMap:=resultMap["data"].(map[string]interface{})
	count = dataMap["count"].(float64)
	return
}