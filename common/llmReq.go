package common

// ReqLLM  LLM request parameters
type ReqLLM struct {
	Prompt       string  `json:"prompt" example:"Funniest joke ever:"` //prompt, not empty
	Temperature  float64 `json:"temperature" example:"0.95"`           //temperature, not empty
	MaxNewTokens int     `json:"max_new_tokens" example:"200"`         //tokens, not empty
	UserToken    string  `json:"user_token" example:"xxxxxxxxxxxx"`    //user_token, not empty
}

type ResLLM struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Text string `json:"text"`
}

type ReqCreateLlmJob struct {
	JobName string `json:"jobName" example:"Millionaires' Problem"` //job name
	Count   int64  `json:"count"`                                   //本次购买llm对话次数
}
