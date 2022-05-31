package models

type Did struct {
	Result Result `json:"result"`
}

type Result struct {
	Did     string `json:"did"`
	Verkey  string `json:"verkey"`
	Posture string `json:"posture"`
	KeyType string `json:"key_type"`
	Method  string `json:"method"`
}
