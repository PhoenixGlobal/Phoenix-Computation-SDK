package common

// ReqBuyLlmToken the request parameters of buying LLM token
type ReqBuyLlmToken struct {
	UserToken string `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
	JobName   string `json:"jobName" example:"Buy LLM Token"`   //job name
	TokensNum int64  `json:"tokensNum"`                         //token num
}

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
