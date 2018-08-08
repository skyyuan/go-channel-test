package models

type Message struct {
	SignalID string `json:"signal_id"`
	Data     string `json:"data"`
}
