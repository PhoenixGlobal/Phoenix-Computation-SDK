package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/controllers"
)

func genImage111() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MDUyMzE0NTd9.B9NHqpb4jDK_U8z1bNm7B1OO9yYOvfcdKsWW_e4D2LE"
	reqBody := common.ReqGenSDXLImage{
		Prompt:          "a yellow dog, Golden Retriever",
		NegativePromt:   "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:       tokenStr,
		NSteps:          40,
		GuidanceRescale: 8,
		HighNoiseFrac:   0.8,
		Seed:            rand.Intn(10),
	}
	res, err := controllers.GenSDXLImage(reqBody)
	fmt.Print("111111------111111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genImage222() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjIyMkBnbWFpbC5jb20iLCJleHAiOjE3MDUyMzE4NTR9.5e0cL9OnIZaSRieNSKFc4ZbaydfLX8hpYjbMukL721g"
	reqBody := common.ReqGenSDXLImage{
		Prompt:          "iron man skiing on steep slope",
		NegativePromt:   "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:       tokenStr,
		NSteps:          40,
		GuidanceRescale: 8,
		HighNoiseFrac:   0.8,
		Seed:            rand.Intn(10),
	}
	res, err := controllers.GenSDXLImage(reqBody)
	fmt.Print("22222------22222")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genImage333() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjMzM0BnbWFpbC5jb20iLCJleHAiOjE3MDUyMzE5MzN9.OGqQii7wxrpQQREcS_kxNsx7Vz67MQFPi8l5IkmpT20"
	reqBody := common.ReqGenSDXLImage{
		Prompt:          "chun-li mixed with serena williams",
		NegativePromt:   "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:       tokenStr,
		NSteps:          40,
		GuidanceRescale: 8,
		HighNoiseFrac:   0.8,
		Seed:            rand.Intn(10),
	}
	res, err := controllers.GenSDXLImage(reqBody)
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genImageBase111() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MDUzMTEzMDN9.Tu1NWp34KSfoNOjIwIeZguQtA7hiiL2X_uk2F0oDM3Y"
	reqBody := common.ReqGenImage{
		Prompt:        "a yellow dog, Golden Retriever",
		NegativePromt: "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:     tokenStr,
	}
	res, err := controllers.GenImage(reqBody)
	fmt.Print("111111------111111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genImageBase222() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjIyMkBnbWFpbC5jb20iLCJleHAiOjE3MDUzMTEzNTF9.p_nqbo1kgEHdiMq_bf2e9XnvdiWhfMikoQwEMw2vJDs"
	reqBody := common.ReqGenImage{
		Prompt:        "iron man skiing on steep slope",
		NegativePromt: "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:     tokenStr,
	}
	res, err := controllers.GenImage(reqBody)
	fmt.Print("222222------2222222")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genImageBase333() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjMzM0BnbWFpbC5jb20iLCJleHAiOjE3MDUzMTEzODR9.5f7t7gv5dAJ7AGJEisz2WV19e4UovW1r63kxYVSyi80"
	reqBody := common.ReqGenImage{
		Prompt:        "chun-li mixed with serena williams",
		NegativePromt: "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:     tokenStr,
	}
	res, err := controllers.GenImage(reqBody)
	fmt.Print("33333------333333")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genGifBase111() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MDY4Njg4ODN9.sxwjmrovuoaTEXGfrF1JI4BdtaaRpq2GMEocRyDBE2A"
	reqBody := common.ReqTextToMotion{
		Prompt:        "iron man skiing on steep slope, hd, high quality",
		NegativePromt: "blurry",
		UserToken:     tokenStr,
	}
	res, err := controllers.TextToMotion(reqBody)
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genGifBase222() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjIyMkBnbWFpbC5jb20iLCJleHAiOjE3MDY4NDgzMjB9.-dSGgYvjhDyLcMIVsUM94TkC9FiWDgF08s3nBLU-PVo"
	reqBody := common.ReqTextToMotion{
		Prompt:        "A dog is eatting bones, hd, high quality",
		NegativePromt: "blurry",
		UserToken:     tokenStr,
	}
	res, err := controllers.TextToMotion(reqBody)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Print("22222------22222")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genGifBase333() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjMzM0BnbWFpbC5jb20iLCJleHAiOjE3MDY4NDgzNTF9.EFz4sUQyVXiCMKxm8MpF9_NzVhPHZ5T-Ns0TE6Wzqp8"
	reqBody := common.ReqTextToMotion{
		Prompt:        "michael jordan dunk",
		NegativePromt: "blurry",
		UserToken:     tokenStr,
	}
	res, err := controllers.TextToMotion(reqBody)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Print("33333------33333")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genSDXLParam111() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MDkxMTIyMDh9.ceXVifxq4MmYI3EwhbfvUZp5GpczkK1ZpP1jcT90lUs"
	reqBody := common.ReqGenSDXLParam{
		Prompt:        "glowneon, starbucks coffee mug emitting sparks, dark amethyst and cyan <lora:glowneon_xl_v1:1> <lora:add-detail-xl:0.8>",
		NegativePromt: "(low quality, worst quality:1.4), cgi,  text, signature, watermark, extra limbs",
		LoraPathList:  "https://civitai-delivery-worker-prod.5ac0637cfd0766c97916cefa3764fbdf.r2.cloudflarestorage.com/model/1269491/grogu.FMBG.safetensors?X-Amz-Expires=86400&response-content-disposition=attachment%3B%20filename%3D%22Grogu.safetensors%22&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=e01358d793ad6966166af8b3064953ad/20240227/us-east-1/s3/aws4_request&X-Amz-Date=20240227T101806Z&X-Amz-SignedHeaders=host&X-Amz-Signature=301cf854248e7fb1382c032a810240b15143c0ccd06b82a724d67c6cee0a8dd1",
		UserToken:     tokenStr,
	}
	res, err := controllers.GenSDXLParam(reqBody)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genSDXLParam222() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjIyMkBnbWFpbC5jb20iLCJleHAiOjE3MDkxMTIyMzd9.4IH-xW7Imbtbz1wy60_6yzRulZHuPqbQbXQSubkcA-s"
	reqBody := common.ReqGenSDXLParam{
		Prompt:        "glowneon, a robot girl, blue, yellow",
		NegativePromt: "low quality, worst quality",
		LoraPathList:  "https://civitai-delivery-worker-prod.5ac0637cfd0766c97916cefa3764fbdf.r2.cloudflarestorage.com/model/3553185/glowneonXlV1.j7UL.safetensors?X-Amz-Expires=86400&response-content-disposition=attachment%3B%20filename%3D%22glowneon_xl_v1.safetensors%22&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=e01358d793ad6966166af8b3064953ad/20240227/us-east-1/s3/aws4_request&X-Amz-Date=20240227T065502Z&X-Amz-SignedHeaders=host&X-Amz-Signature=24bfd761a442ff1a05aa7bab4f3c9b0c524379f4c2dd425558505efff0d08c90",
		UserToken:     tokenStr,
	}
	res, err := controllers.GenSDXLParam(reqBody)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Print("22222------22222")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genSDXLParam333() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjMzM0BnbWFpbC5jb20iLCJleHAiOjE3MDkxMTIyNzR9.LQub0KG74WuflprJQrChGgGbGVkyC2s92ctGJwfVfB0"
	reqBody := common.ReqGenSDXLParam{
		Prompt:        "glowneon, glowing samurai emitting sparks and electricity, dark red and orange, glowing eyes, cinematic film still <lora:glowneon_xl_v1:1> <lora:add-detail-xl:0.8>",
		NegativePromt: "(low quality, worst quality:1.4), cgi,  text, signature, watermark, extra limbs",
		UserToken:     tokenStr,
	}
	res, err := controllers.GenSDXLParam(reqBody)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Print("33333------33333")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genBaseParam111() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTA3NTY1NjN9.1sgxCiyGVgta8DNE9WJxu22fP3is87R3-B-Y3ecRUMo"
	reqBody := common.ReqGenImage{
		Prompt:         "a yellow dog, Golden Retriever",
		NegativePromt:  "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:      tokenStr,
		LoraPathList:   "Glowing_Runes.safetensors",
		LoraWeightList: "0.9",
	}
	res, err := controllers.GenImage(reqBody)
	fmt.Print("111111------111111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genBaseParam222() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjIyMkBnbWFpbC5jb20iLCJleHAiOjE3MTA3NTY1ODZ9.VY0YzYSIHJEkoiPZ6HJra0Cvxi3wEScvySPUUJY5dnU"
	reqBody := common.ReqGenImage{
		Prompt:         "iron man skiing on steep slope",
		NegativePromt:  "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:      tokenStr,
		LoraPathList:   "Glowing_Runes.safetensors",
		LoraWeightList: "0.9",
	}
	res, err := controllers.GenImage(reqBody)
	fmt.Print("222222------2222222")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func genBaseParam333() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjMzM0BnbWFpbC5jb20iLCJleHAiOjE3MTA3NTY2MDd9.kZOBS7ijlxVZFbCz22VYvuFoLCBO7EyUxKb43_tFyNw"
	reqBody := common.ReqGenImage{
		Prompt:         "chun-li mixed with serena williams",
		NegativePromt:  "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:      tokenStr,
		LoraPathList:   "Glowing_Runes.safetensors",
		LoraWeightList: "0.9",
	}
	res, err := controllers.GenImage(reqBody)
	fmt.Print("33333------333333")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToImgSDXl111() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTEwOTkxNjN9.6XYCXuWoA_aqwnhOo_Q7aF-jZxp8C95tqVfoA5hsdN4"
	reqBody := common.ReqGenSDXLParam{
		Prompt:               "Passenger plane, flying upward",
		NegativePromt:        "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:            tokenStr,
		NSteps:               40,
		GuidanceScale:        9,
		Seed:                 3493035682,
		PipeName:             "img2img",
		LoraPathList:         "400GB-LoRA-XL-Repository.safetensors",
		LoraWeightList:       "0.9",
		TextualInversionPath: "without",
		LoraScale:            1,
		Size:                 "1024^*^1024",
		ClipSkip:             1,
		InitImage:            "https://api.telegram.org/file/bot6951423417:AAFUgGCPJow64o2NiWNqDy2L8U84ggsPsvs/photos/file_15.jpg",
	}
	res, err := controllers.GenSDXLParam(reqBody)
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToImgSDXl222() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjIyMkBnbWFpbC5jb20iLCJleHAiOjE3MTEwOTk4NzN9.7z-XfCHMCFDsi4ZqGIL_MzxcMvUZdUf_vdBEohsUNaw"
	reqBody := common.ReqGenSDXLParam{
		Prompt:               "Passenger plane, flying upward",
		NegativePromt:        "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:            tokenStr,
		NSteps:               40,
		GuidanceScale:        9,
		Seed:                 3493035682,
		PipeName:             "img2img",
		LoraPathList:         "400GB-LoRA-XL-Repository.safetensors",
		LoraWeightList:       "0.9",
		TextualInversionPath: "without",
		LoraScale:            1,
		Size:                 "1024^*^1024",
		ClipSkip:             1,
		InitImage:            "https://api.telegram.org/file/bot6951423417:AAFUgGCPJow64o2NiWNqDy2L8U84ggsPsvs/photos/file_15.jpg",
	}
	res, err := controllers.GenSDXLParam(reqBody)
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToImgSDXl333() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjMzM0BnbWFpbC5jb20iLCJleHAiOjE3MTEwOTk5MTF9.8I9cWN4ulc3ki_wMqENlHTgh_PaSjVdvwD3YsxrYClU"
	reqBody := common.ReqGenSDXLParam{
		Prompt:               "Passenger plane, flying upward",
		NegativePromt:        "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:            tokenStr,
		NSteps:               40,
		GuidanceScale:        9,
		Seed:                 3493035682,
		PipeName:             "img2img",
		LoraPathList:         "400GB-LoRA-XL-Repository.safetensors",
		LoraWeightList:       "0.9",
		TextualInversionPath: "without",
		LoraScale:            1,
		Size:                 "1024^*^1024",
		ClipSkip:             1,
		InitImage:            "https://api.telegram.org/file/bot6951423417:AAFUgGCPJow64o2NiWNqDy2L8U84ggsPsvs/photos/file_15.jpg",
	}
	res, err := controllers.GenSDXLParam(reqBody)
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToMotion111() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTExMTM4MzB9.u9d-wQLJAZt5KXpDS6XD7Is19P6HTeOMZfq9w5lWblc"
	reqBody := common.ReqImgToMotion{
		ImagePath: "https://image.civitai.com/xG1nkqKTMzGDvpLrqFT7WA/fac5e435-3fc0-4638-bdb8-15b350b8dcef/original=true/tmp8mhp14g3.jpeg",
		UserToken: tokenStr,
	}
	res, err := controllers.ImgToMotion(reqBody)
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToMotion222() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTExMTM4MzB9.u9d-wQLJAZt5KXpDS6XD7Is19P6HTeOMZfq9w5lWblc"
	reqBody := common.ReqImgToMotion{
		ImagePath: "https://image.civitai.com/xG1nkqKTMzGDvpLrqFT7WA/fac5e435-3fc0-4638-bdb8-15b350b8dcef/original=true/tmp8mhp14g3.jpeg",
		UserToken: tokenStr,
	}
	res, err := controllers.ImgToMotion(reqBody)
	fmt.Print("22222------22222")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToMotion333() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTExMTM4MzB9.u9d-wQLJAZt5KXpDS6XD7Is19P6HTeOMZfq9w5lWblc"
	reqBody := common.ReqImgToMotion{
		ImagePath: "https://image.civitai.com/xG1nkqKTMzGDvpLrqFT7WA/fac5e435-3fc0-4638-bdb8-15b350b8dcef/original=true/tmp8mhp14g3.jpeg",
		UserToken: tokenStr,
	}
	res, err := controllers.ImgToMotion(reqBody)
	fmt.Print("33333------33333")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToImgBase111() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTEwOTkxNjN9.6XYCXuWoA_aqwnhOo_Q7aF-jZxp8C95tqVfoA5hsdN4"
	reqBody := common.ReqGenImage{
		Prompt:               "airliner, clear weather",
		NegativePromt:        "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:            tokenStr,
		NSteps:               40,
		GuidanceScale:        9,
		Seed:                 3493035682,
		PipeName:             "img2img",
		LoraPathList:         "^*^",
		LoraWeightList:       "0.8",
		TextualInversionPath: "without",
		LoraScale:            1,
		Size:                 "512^*^512",
		ClipSkip:             1,
		InitImage:            "https://api.telegram.org/file/bot6951423417:AAFUgGCPJow64o2NiWNqDy2L8U84ggsPsvs/photos/file_41.jpg",
	}
	res, err := controllers.GenImage(reqBody)
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToImgBase222() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTEwOTkxNjN9.6XYCXuWoA_aqwnhOo_Q7aF-jZxp8C95tqVfoA5hsdN4"
	reqBody := common.ReqGenImage{
		Prompt:               "airliner, clear weather",
		NegativePromt:        "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:            tokenStr,
		NSteps:               40,
		GuidanceScale:        9,
		Seed:                 3493035682,
		PipeName:             "img2img",
		LoraPathList:         "^*^",
		LoraWeightList:       "0.8",
		TextualInversionPath: "without",
		LoraScale:            1,
		Size:                 "512^*^512",
		ClipSkip:             1,
		InitImage:            "https://api.telegram.org/file/bot6951423417:AAFUgGCPJow64o2NiWNqDy2L8U84ggsPsvs/photos/file_41.jpg",
	}
	res, err := controllers.GenImage(reqBody)
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToImgBase333() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjMzM0BnbWFpbC5jb20iLCJleHAiOjE3MTEwOTk5MTF9.8I9cWN4ulc3ki_wMqENlHTgh_PaSjVdvwD3YsxrYClU"
	reqBody := common.ReqGenImage{
		Prompt:               "airliner, clear weather",
		NegativePromt:        "FastNegativeV2, (low quality:1.3),(worst quality:1.3),(monochrome:0.8),(deformed:1.3),(malformed hands:1.4),(poorly drawn hands:1.4),(mutated fingers:1.4),(bad anatomy:1.3),(extra limbs:1.35),(poorly drawn face:1.4),(watermark:1.3),ugly,watermark,jpeg artifacts,error,text,username",
		UserToken:            tokenStr,
		NSteps:               40,
		GuidanceScale:        9,
		Seed:                 3493035682,
		PipeName:             "img2img",
		LoraPathList:         "^*^",
		LoraWeightList:       "0.8",
		TextualInversionPath: "without",
		LoraScale:            1,
		Size:                 "512^*^512",
		ClipSkip:             1,
		InitImage:            "https://api.telegram.org/file/bot6951423417:AAFUgGCPJow64o2NiWNqDy2L8U84ggsPsvs/photos/file_41.jpg",
	}
	res, err := controllers.GenImage(reqBody)
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToPrompt111() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTM4NDcyNzR9.q7AWMXHEz2HyxCd7MByTCPrX5-9NKKHMFsVp9STOBhA"
	reqBody := common.ReqImgToPrompt{
		Prompt:    "",
		UserToken: tokenStr,
		ImagePath: "https://picx.zhimg.com/70/v2-3b4fc7e3a1195a081d0259246c38debc_1440w.avis?source=172ae18b&biz_tag=Post",
	}
	res, err := controllers.ImgToPrompt(reqBody)
	fmt.Print("11111------11111")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToPrompt222() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTM4NDcyNzR9.q7AWMXHEz2HyxCd7MByTCPrX5-9NKKHMFsVp9STOBhA"
	reqBody := common.ReqImgToPrompt{
		Prompt:    "",
		UserToken: tokenStr,
		ImagePath: "https://inews.gtimg.com/news_bt/OIPr9G8LrCpP4K3cdWrmvalN2p2YWWB3URkibvS38awOEAA/641",
	}
	res, err := controllers.ImgToPrompt(reqBody)
	fmt.Print("22222------22222")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func imgToPrompt333() {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjExMUBnbWFpbC5jb20iLCJleHAiOjE3MTM4NTUxODl9.JxxYw84bVPU7RaLRkGSOW3-7TT1Mv-7VfpFqP-iYapI"
	reqBody := common.ReqImgToPrompt{
		Prompt:    "",
		UserToken: tokenStr,
		ImagePath: "https://inews.gtimg.com/news_bt/OIPr9G8LrCpP4K3cdWrmvalN2p2YWWB3URkibvS38awOEAA/641",
	}
	res, err := controllers.ImgToPrompt(reqBody)
	fmt.Print("33333------33333")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(res, &resultMap)
	fmt.Println(err, resultMap)
}

func main() {
	// go genImage111()
	// go genImage222()
	// go genImage333()
	// go genGifBase111()
	// go genGifBase222()
	// go genGifBase333()
	// go genSDXLParam111()
	// go genSDXLParam222()
	// go genSDXLParam333()
	// go genBaseParam111()
	// go genBaseParam222()
	// go genBaseParam333()
	// go imgToMotion111()
	// go imgToImgSDXl111()
	// go imgToImgSDXl222()
	// go imgToImgSDXl333()
	// go imgToImgBase111()
	// go imgToImgBase222()
	// go imgToImgBase333()
	// go imgToMotion111()
	// go imgToMotion222()
	// imgToMotion333()
	// go imgToPrompt111()
	// go imgToPrompt222()
	// go imgToPrompt333()
	// time.Sleep(1180 * time.Second)
	fileName := "./winner.csv"
	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
	// 创建CSV读取器
	reader := csv.NewReader(file)
	// var resultAddressList []string
	resultAddress := "["
	// 循环读取每行
	i := 1
	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println("读取文件失败：", err)
			break
		}
		// fmt.Println(record)
		// if record[9] == "connectedWallet" {
		// 	continue
		// }
		resultAddress = resultAddress + "\"" + record[0] + "\"" + ","
		if i%100 == 0 || i == 229 {
			resultAddress = resultAddress[:len(resultAddress)-1] + "]"
			fmt.Println(resultAddress)
			fmt.Println()
			fmt.Println()
			resultAddress = "["
		}
		i = i + 1
	}
	defer file.Close()
}
