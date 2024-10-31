package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/util"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
)

const (
	llmEndPoint      = "LLM/"
	buyLLMTokenMehod = "buyLLMToken"
	llmPriceMethod   = "queryLLMPrice"
	llm3Method       = "callLLM3"
	verticalMethod   = "verticalLLM"
)

// BuyLLMToken bug LLM token
func BuyLLMToken(reqJSON common.ReqBuyLlmToken) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, llmEndPoint, buyLLMTokenMehod)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryLLMPrice query LLM price
func QueryLLMPrice(token string) (json.RawMessage, error) {
	rawURL := fmt.Sprint(apiEndpoint, llmEndPoint, llmPriceMethod)
	result, err := util.SendHttpGet(rawURL, nil, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CallLLM3  call LLM api
func CallLLM3(reqJSON common.ReqLLM3) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, llmEndPoint, llm3Method)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VerticalLLM  VerticalLLM api
func VerticalLLM(reqJSON common.ReqLLM3) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, llmEndPoint, verticalMethod)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}
