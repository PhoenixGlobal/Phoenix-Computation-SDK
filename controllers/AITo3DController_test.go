package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
	"testing"
)

func GetToken() string {
	reqToken := common.ReqGenToken{
		Email:  "xxx@gmail.com",
		Passwd: "xxxxxx"}
	token, _ := GenToken(reqToken)
	tokenMap := make(map[string]interface{})
	err := json.Unmarshal(token, &tokenMap)
	if err != nil {
		fmt.Println("err=", err)
		return ""
	}
	tokenStr := tokenMap["token"].(string)
	return tokenStr
}

//create text2Model job
func TestCreateAITo3DJob(t *testing.T) {
	token := GetToken()
	reqBody := common.ReqAITo3DJob{
		JobName:     "test AITo3D test",
		Computation: "AI-to-3D (Tripo)",
		Prompt:      "a cute dog",
		Type:        "text2Model",
		ImageURL:    "",
	}
	result, err := CreateAITo3DJob(reqBody,token)
	fmt.Println(111111, string(result), err)
}

//create img2Model job
func TestCreateAITo3DJob2(t *testing.T) {
	token := GetToken()

	fileName := "absolute file path"
	result, err := UploadDeAIFile(fileName, token)
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(result, &resultMap)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	imgUrl := resultMap["msg"].(string)

	reqBody := common.ReqAITo3DJob{
		JobName:     "test AITo3D ",
		Computation: "AI-to-3D (Tripo)",
		Prompt:      "",
		Type:        "img2Model",
		ImageURL:    imgUrl,
	}
	result, err = CreateAITo3DJob(reqBody,token)
	fmt.Println(111111, string(result), err)
}

func TestQueryAITo3DJobDetail(t *testing.T) {
	token := GetToken()
	reqBody := common.ReqQueryAITo3DJobDetail{
		JobID:  "21559616267152752126804958740584316039613520634800696977564855889808589769897",
	}
	result, err := QueryAITo3DJobDetail(reqBody,token)
	fmt.Println(111111, string(result), err)
}

func TestQueryAITo3DTaskModelURL(t *testing.T) {
	token := GetToken()
	reqBody := common.ReqQueryAITo3DTaskModelURL{
		JobID:  "21559616267152752126804958740584316039613520634800696977564855889808589769897",
		TaskID: "b98f8c0d-5b02-4195-a750-de5fc7c1f7bf",
	}
	result, err := QueryAITo3DTaskModelURL(reqBody,token)
	fmt.Println(111111, string(result), err)
}

func TestAnimate(t *testing.T) {
	token := GetToken()
	reqBody := common.ReqAnimateModel{
		JobID:  "21559616267152752126804958740584316039613520634800696977564855889808589769897",
		TaskID: "b98f8c0d-5b02-4195-a750-de5fc7c1f7bf",
		Animation: "preset:walk",
	}
	result, err := Animate(reqBody,token)
	fmt.Println(111111, string(result), err)
}

func TestConvert(t *testing.T) {
	token := GetToken()
	reqBody := common.ReqConvertModel{
		JobID:  "21559616267152752126804958740584316039613520634800696977564855889808589769897",
		TaskID: "b98f8c0d-5b02-4195-a750-de5fc7c1f7bf",
		Format: "FBX",
	}
	result, err := Convert(reqBody,token)
	fmt.Println(111111, string(result), err)
}