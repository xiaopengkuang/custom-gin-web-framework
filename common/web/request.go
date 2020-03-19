package web

type Request struct {
	Module    string `json:"module"`
	Service   string `json:"service"`
	Operation string `json:"operation"`
	Para      string `json:"para"`
}
