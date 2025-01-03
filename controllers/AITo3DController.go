package controllers

import (
	"encoding/json"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/util"
)

const SkyNetControllerURL string ="https://www.phoenix.global/prd"

// CreateAITo3DJob Create a job of input type
func CreateAITo3DJob(reqBody common.ReqAITo3DJob, token string) (json.RawMessage, error) {
	url:=SkyNetControllerURL+"/deepLearn/createTripo3d"
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(url, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func QueryAITo3DJobDetail(reqBody common.ReqQueryAITo3DJobDetail, token string) (json.RawMessage, error) {
	url:=SkyNetControllerURL+"/deepLearn/queryTripoDetail"
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(url, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func QueryAITo3DTaskModelURL(reqBody common.ReqQueryAITo3DTaskModelURL, token string) (json.RawMessage, error) {
	url:=SkyNetControllerURL+"/deepLearn/queryTripo3d"
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(url, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

////The value for animation can only be preset:idle,preset:walk,preset:climb,preset:jump,preset:run,preset:slash,preset:shoot,preset:hurt,preset:fall,preset:turn.
func Animate(reqBody common.ReqAnimateModel, token string) (json.RawMessage, error) {
	url:=SkyNetControllerURL+"/deepLearn/animateTrip3d"
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(url, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//You can only convert GLB to USDZ, FBX, OBJ, or STL formats, and the conversion can only be done once.
func Convert(reqBody common.ReqConvertModel, token string) (json.RawMessage, error) {
	url:=SkyNetControllerURL+"/deepLearn/convert"
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(url, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}