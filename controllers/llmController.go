package controllers

import (
	"encoding/json"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/util"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
)

const CallLLMURL string = "https://www.phoenix.global/sdk/computation/LLM/callLLM"
const GenImageURL string = "https://www.phoenix.global/sdk/computation/LLM/callGenImage"
const CreateLLMJobURL string = "https://www.phoenix.global/sdk/computation/LLM/createLLMJob"
const QueryLLMPriceURL string = "https://www.phoenix.global/sdk/computation/LLM/queryLLMPrice"
const QueryLLMCountURL string = "https://www.phoenix.global/sdk/computation/LLM/queryLLMActualCount"
const QueryLLMTokensBalanceURL string = "https://www.phoenix.global/sdk/computation/LLM/queryLLMTokensBalance"
const QueryLLMFreeTokensBalanceURL string = "https://www.phoenix.global/sdk/computation/LLM/queryLLMFreeTokensBalance"

// CallLLM  call LLM api
func CallLLM(reqBody common.ReqLLM) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(CallLLMURL, reqJson, reqBody.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
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

func GetLLMBuyCount(token string) (count float64, e error) {
	res, err := QueryLLMBuyCount(token)
	if err != nil {
		e = err
		return
	}
	resultMap := make(map[string]interface{})
	e = json.Unmarshal(res, &resultMap)
	if e != nil {
		return
	}
	dataMap := resultMap["data"].(map[string]interface{})
	count = dataMap["count"].(float64)
	return
}

func QueryLLMTokensBalance(token string) (llmTokens float64, e error) {
	result, err := util.SendHttpGet(QueryLLMTokensBalanceURL, nil, token)
	if err != nil {
		e = err
		return
	}
	resultMap := make(map[string]interface{})
	e = json.Unmarshal(result, &resultMap)
	if e != nil {
		return
	}
	dataMap := resultMap["data"].(map[string]interface{})
	llmTokens = dataMap["llmTokens"].(float64)
	return
}

func QueryLLMFreeTokensBalance(token string) (llmTokens float64, e error) {
	result, err := util.SendHttpGet(QueryLLMFreeTokensBalanceURL, nil, token)
	if err != nil {
		e = err
		return
	}
	resultMap := make(map[string]interface{})
	e = json.Unmarshal(result, &resultMap)
	if e != nil {
		return
	}
	dataMap := resultMap["data"].(map[string]interface{})
	llmTokens = dataMap["llmTokens"].(float64)
	return
}

// GenImage  call LLM api
func GenImage(reqBody common.ReqGenImage) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(GenImageURL, reqJson, reqBody.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}
