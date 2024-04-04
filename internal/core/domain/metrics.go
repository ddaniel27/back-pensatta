package domain

type Spider map[string]interface{}

type Appropiation map[string]interface{}

type Metrics struct {
	SpiderValues       Spider       `json:"spider_values"`
	AppropiationValues Appropiation `json:"appropiation_values"`
}
