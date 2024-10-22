package common

// ReqLLM  LLM request parameters
type ReqLLM struct {
	Prompt       string  `json:"prompt" example:"Funniest joke ever:"` //prompt, not empty
	Temperature  float64 `json:"temperature" example:"0.95"`           //temperature, not empty
	MaxNewTokens int     `json:"max_new_tokens" example:"200"`         //tokens, not empty
	UserToken    string  `json:"user_token" example:"xxxxxxxxxxxx"`    //user_token, not empty
}

// ReqLLM3  LLM llama3 request parameters
type ReqLLM3 struct {
	Prompt       string `json:"prompt" example:"Funniest joke ever:"` //prompt, not empty
	MaxNewTokens int    `json:"max_new_tokens" example:"200"`         //tokens, not empty
	UserToken    string `json:"user_token" example:"xxxxxxxxxxxx"`    //user_token, not empty
}

type ReqCreateLlmJob struct {
	JobName   string `json:"jobName" example:"Millionaires' Problem"` //job name
	TokensNum int64  `json:"tokensNum"`                               //本次购买llm对话次数
}

// ReqTextToMotion TextToMotion request parameters
type ReqTextToMotion struct {
	Prompt        string `json:"prompt"`                            //prompt, not empty
	NegativePromt string `json:"negative_prompt"`                   //tokens, not empty
	UserToken     string `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
}

// ReqImgToMotion ImgToMotion request parameters
type ReqImgToMotion struct {
	ImagePath string `json:"image_path"`                        //image path, not empty
	UserToken string `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
}

// ReqGenSDXLParam gen SDXL by parameters
type ReqGenSDXLParam struct {
	UidName              string  `json:"uidname" example:"111@gmail.com"`
	Prompt               string  `json:"prompt" example:"a blue dog"`       //prompt, not empty
	NegativePromt        string  `json:"negative_prompt"`                   //tokens, not empty
	UserToken            string  `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
	NSteps               int     `json:"n_steps" example:"20"`
	GuidanceScale        float64 `json:"guidance_scale" example:"8"` // float, 0 to 10
	Seed                 int     `json:"seed" example:"1"`           // int, >0
	PipeName             string  `json:"pipe_name" example:"txt2img"`
	LoraPathList         string  `json:"lora_path_list" example:"https://xxxx"`
	LoraWeightList       string  `json:"lora_weight_list" example:"0.8"`
	TextualInversionPath string  `json:"textual_inversion_path" example:"without"`
	LoraScale            float64 `json:"lora_scale" example:"0.8"`
	Size                 string  `json:"size" example:"1024^*^1024"`
	ClipSkip             int     `json:"clip_skip" example:"1"`
	InitImage            string  `json:"init_image" example:"https://xxx.jpg"`
	MaskImage            string  `json:"mask_image" example:"https://xxx.jpg"`
}

// ReqImgToPrompt ImgToPrompt parameters
type ReqImgToPrompt struct {
	UserToken string `json:"user_token" example:"xxxxxxxxxxxx"`    //user_token, not empty
	ImagePath string `json:"image_path" example:"https://xxx.jpg"` //image path, not empty
	Prompt    string `json:"prompt" example:"Funniest joke ever:"` //prompt, not empty
}
