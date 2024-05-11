package model

type DetailTransaction struct {
	Trx_id     string    `json:"trx_id"`
	Apv_code   string    `json:"apv_code"`
	Trx_typ    string    `json:"trx_typ"`
	Amt        float32   `json:"amt"`
	Status     string    `json:"status"`
	Desc       string    `json:"desc"`
	Loc_acct   string    `json:"loc_acct"`
}