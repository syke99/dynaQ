package models

type QueryValue struct {
	Type   string      `json:"type"`
	Column string      `json:"name"`
	Value  interface{} `json:"value"`
}
