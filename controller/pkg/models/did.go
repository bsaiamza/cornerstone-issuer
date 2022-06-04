package models

type GetPublicDidResponse struct {
	Result Did `json:"result"`
}

type Did struct {
	Did     string `json:"did"`
	Verkey  string `json:"verkey"`
	Posture string `json:"posture"`
	KeyType string `json:"key_type"`
	Method  string `json:"method"`
}
