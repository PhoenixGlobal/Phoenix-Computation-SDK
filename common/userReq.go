package common

// ReqGenToken user login request parameters
type ReqGenToken struct {
	Email  string `json:"email" example:"xxxxxxx@gmail.com"` //email address, not empty
	Passwd string `json:"passwd" example:"xxxxxxxxxxxx"`     //password, not empty
}
