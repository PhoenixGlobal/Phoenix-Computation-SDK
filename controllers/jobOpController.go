package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/util"
)

const CreateInputJobURL string = "https://www.phoenix.global/sdk/computation/panel/createJobByInput"
const CreateFileJobURL string = "https://www.phoenix.global/sdk/computation/panel/createJobByFile"
const FillDataURL string = "https://www.phoenix.global/sdk/computation/panel/fillData"
const DeleteJobURL string = "https://www.phoenix.global/sdk/computation/panel/deleteJobByUser"
const UploadFileURL string = "https://www.phoenix.global/sdk/computation/deAI/uploadFile"
const CreateAIJobURL string = "https://www.phoenix.global/sdk/computation/deAI/createJob"
const InferenceJobURL string = "https://www.phoenix.global/sdk/computation/deAI/inference"

// CreateJobByInput Create a job of input type
func CreateJobByInput(reqBody common.ReqCreateJobByInput, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(CreateInputJobURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateJobByFile Create a job of upload type
func CreateJobByFile(params common.ReqCreateJobByFile, featureFileName string, targetFileName string, testingFileName string, token string) (json.RawMessage, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	boundary := body_writer.Boundary()
	body_writer.SetBoundary(boundary)

	// set params
	body_writer.WriteField("jobName", params.JobName)
	body_writer.WriteField("computation", params.Computation)
	// set feature file
	_, err := body_writer.CreateFormFile("featureData", featureFileName)
	if err != nil {
		return nil, err
	}
	featureData, err := ioutil.ReadFile(featureFileName)
	if err != nil {
		return nil, err
	}
	body_buf.Write(featureData)
	// set target file
	if targetFileName != "" {
		_, err = body_writer.CreateFormFile("targetData", targetFileName)
		if err != nil {
			return nil, err
		}
		targetData, err := ioutil.ReadFile(targetFileName)
		if err != nil {
			return nil, err
		}
		body_buf.Write(targetData)
	}
	// set testing file
	if testingFileName != "" {
		_, err = body_writer.CreateFormFile("testingData", testingFileName)
		if err != nil {
			return nil, err
		}
		testingData, err := ioutil.ReadFile(testingFileName)
		if err != nil {
			return nil, err
		}
		body_buf.Write(testingData)
	}

	body_writer.Close()
	result, err := util.SendPostMultiPart(CreateFileJobURL, body_buf, body_writer, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FillData(reqBody common.ReqFillData, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(FillDataURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteJobByUser(reqBody common.ReqDeleteJobByUser, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(DeleteJobURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UploadDeAIFile Upload file
func UploadDeAIFile(fileName string, token string) (json.RawMessage, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	boundary := body_writer.Boundary()
	body_writer.SetBoundary(boundary)
	_, err := body_writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	body_buf.Write(fileData)
	body_writer.Close()
	result, err := util.SendPostMultiPart(UploadFileURL, body_buf, body_writer, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateAIJob(reqBody common.ReqCreateAIJob, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(CreateAIJobURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func InferenceJob(reqBody common.ReqInference, token string) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(InferenceJobURL, reqJson, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}
