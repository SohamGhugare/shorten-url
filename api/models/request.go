package models

import "time"

/*
REQUEST MODEL
Contains all the necessary fields for binding an incoming request
*/
type Request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}
