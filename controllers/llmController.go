package controllers

import (
	"encoding/json"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/util"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
)

const LlmApiURL string = "https://www.phoenix.global/sdk/llm/v1/completions"

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
	if err!=nil{
		return nil,err
	}
	textMapArr:=resultMap["choices"].([]interface{})
	textMap:=textMapArr[0].(map[string]interface{})
	text:=textMap["text"].(string)
	retResult:=common.ResLLM{
		Code: 200,
		Msg:  "success",
		Text: text,
	}
	res, err := json.Marshal(retResult)
	if err != nil {
		return nil, err
	}
	return res, err
}