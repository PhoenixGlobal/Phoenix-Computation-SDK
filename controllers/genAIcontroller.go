package controllers

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/util"
)

const (
	apiEndpoint      = "https://www.phoenix.global/sdk/computation/"
	genAIEndPoint    = "genAI/"
	queryTaskMethod  = "queryTask"
	genImgMethod     = "genImg"
	genImgSDXLMethod = "genImgSDXL"
	genImgFluxMethod = "genImgFlux"
)

// QueryTask query task by task id
func QueryTask(taskID string, token string) (json.RawMessage, error) {
	params := url.Values{}
	params.Add("task_id", taskID)
	rawURL := fmt.Sprint(apiEndpoint, genAIEndPoint, queryTaskMethod)
	result, err := util.SendHttpGet(rawURL, params, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GenImg gen image
func GenImg(reqJSON common.ReqGenImg) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, genAIEndPoint, genImgMethod)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GenImgSDXL  gen image by SDXL
func GenImgSDXL(reqJSON common.ReqGenImgSDXL) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, genAIEndPoint, genImgSDXLMethod)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GenImgFlux  gen image by Flux
func GenImgFlux(reqJSON common.ReqGenImgFlux) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, genAIEndPoint, genImgFluxMethod)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}
