package model

type WafHeader struct {
	Key     string `json:"key"`
	ReValue string `json:"value"`
}

type WafItems struct {
	Name      string      `json:"name"`
	ReHeaders []WafHeader `json:"headers"`
	ReContent []string    `json:"content"`
	ReCookies []string    `json:"cookie"`
}

type Waf struct {
	Items []WafItems `json:"items"`
}
