package model

type LoginRequest struct {
	Acct_Num string `json:"acct_num"`
	Password string `json:"password"`
}