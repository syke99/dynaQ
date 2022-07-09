package models

type QueryValue struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
