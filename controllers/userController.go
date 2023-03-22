package controllers

import (
	"encoding/json"
	"net/url"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/util"

	"github.com/PhoenixGlobal/Phoenix-Computation-SDK/common"
)

const GenTokenURL string = "https://www.phoenix.global/sdk/computation/user/genToken"
const UserInfoURL string = "https://www.phoenix.global/sdk/computation/user/userInfo"

// GenToken Get access token
func GenToken(reqBody common.ReqGenToken) (json.RawMessage, error) {
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	result, err := util.SendHttpPost(GenTokenURL, reqJson, "")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserInfo Query user information
func UserInfo(token string) (json.RawMessage, error) {
	params := url.Values{}
	result, err := util.SendHttpGet(UserInfoURL, params, token)
	if err != nil {
		return nil, err
	}
	return result, nil
}
