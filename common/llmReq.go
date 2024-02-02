package common

// ReqLLM  LLM request parameters
type ReqLLM struct {
	Prompt       string  `json:"prompt" example:"Funniest joke ever:"` //prompt, not empty
	Temperature  float64 `json:"temperature" example:"0.95"`           //temperature, not empty
	MaxNewTokens int     `json:"max_new_tokens" example:"200"`         //tokens, not empty
	UserToken    string  `json:"user_token" example:"xxxxxxxxxxxx"`    //user_token, not empty
}

type ReqCreateLlmJob struct {
	JobName   string `json:"jobName" example:"Millionaires' Problem"` //job name
	TokensNum int64  `json:"tokensNum"`                               //本次购买llm对话次数
}

// ReqGenImage genImage request parameters
type ReqGenImage struct {
	UidName       string `json:"uidname" example:"111@gmail.com"`
	Prompt        string `json:"prompt"`                            //prompt, not empty
	NegativePromt string `json:"negative_prompt"`                   //tokens, not empty
	UserToken     string `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
}

// ReqGenSDXLImage gen SDXL image request parameters
type ReqGenSDXLImage struct {
	UidName         string  `json:"uidname" example:"111@gmail.com"`
	Prompt          string  `json:"prompt" example:"a blue dog"`       //prompt, not empty
	NegativePromt   string  `json:"negative_prompt"`                   //tokens, not empty
	UserToken       string  `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
	NSteps          int     `json:"n_steps" example:"20"`
	GuidanceRescale float64 `json:"guidance_rescale" example:"8"`  // float, 0 to 10
	HighNoiseFrac   float64 `json:"high_noise_frac" example:"0.8"` // float, 0 to 1
	Seed            int     `json:"seed" example:"1"`              // int, >0
}

// ReqTextToMotion TextToMotion request parameters
type ReqTextToMotion struct {
	UidName       string `json:"uidname" example:"111@gmail.com"`
	Prompt        string `json:"prompt"`                            //prompt, not empty
	NegativePromt string `json:"negative_prompt"`                   //tokens, not empty
	UserToken     string `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
}
