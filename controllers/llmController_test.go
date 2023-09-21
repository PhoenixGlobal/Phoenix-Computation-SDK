package controllers

import (
	"fmt"
	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
	"testing"
)

func TestCallLLM(t *testing.T) {
	reqBody:=common.ReqLLM{
		Prompt:       "Funniest joke ever:",
		Temperature:  0.95,
		MaxNewTokens: 200,
		UserToken: "",
	}
	result,err:=CallLLM(reqBody)
	fmt.Println(111111,string(result),err)
}