package common

// ReqLLM  LLM request parameters
type ReqLLM struct {
	Prompt       string  `json:"prompt" example:"Funniest joke ever:"` //prompt, not empty
	Temperature  float64 `json:"temperature" example:"0.95"`           //temperature, not empty
	MaxNewTokens int     `json:"max_new_tokens" example:"200"`         //tokens, not empty
	UserToken    string  `json:"user_token" example:"xxxxxxxxxxxx"`    //user_token, not empty
}

// ReqGenImage genImage request parameters
type ReqGenImage struct {
	UidName       string `json:"uidname" example:"111@gmail.com"`
	Prompt        string `json:"prompt"`                            //prompt, not empty
	NegativePromt string `json:"negative_prompt"`                   //tokens, not empty
	UserToken     string `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
}

type ResLLM struct {
	Code          int     `json:"code"`
	Msg           string  `json:"msg"`
	Text          string  `json:"text"`
	TokensBalance float64 `json:"tokens_balance"`
}

type ReqCreateLlmJob struct {
	JobName   string `json:"jobName" example:"Millionaires' Problem"` //job name
	TokensNum int64  `json:"tokensNum"`                               //本次购买llm对话次数
}
