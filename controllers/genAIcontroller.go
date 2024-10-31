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
	txt2MotionMethod = "txt2Motion"
	img2MotionMethod = "img2Motion"
	cogVideoMethod   = "cogVideo"
	pyramidMethod    = "pyramid"
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

// Txt2Motion  call text to motion api
func Txt2Motion(reqJSON common.ReqTxt2Motion) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, genAIEndPoint, txt2MotionMethod)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Img2Motion  call img to motion api
func Img2Motion(reqJSON common.ReqImg2Motion) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, genAIEndPoint, img2MotionMethod)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CogVideo Generating a gif by CogVideo mode
func CogVideo(reqJSON common.ReqCogVideo) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, genAIEndPoint, cogVideoMethod)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Pyramid Generating a gif by Pyramid mode
func Pyramid(reqJSON common.ReqPyramid) (json.RawMessage, error) {
	reqBody, err := json.Marshal(reqJSON)
	if err != nil {
		return nil, err
	}
	rawURL := fmt.Sprint(apiEndpoint, genAIEndPoint, pyramidMethod)
	result, err := util.SendHTTPPostForLLM(rawURL, reqBody, reqJSON.UserToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}
