package service

import "github.com/hugdubois/svc-fizzbuzz/core"

// ErrorMessage is the message returned when an error has occurred.
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// FizzBuzzResponse is the message returned by FizzBuzzHandler.
type FizzBuzzResponse struct {
	FizzBuzz []string `json:"fizzbuzz"`
}

// FizzBuzzTopResponse is the message returned by FizzBuzzTopHandler
type FizzBuzzTopResponse struct {
	Data FizzBuzzTopResponseData `json:"data"`
}

// FizzBuzzTopResponseData is the core of message returned by FizzBuzzTopHandler.
type FizzBuzzTopResponseData struct {
	Params       core.FizzBuzzParams `json:"params"`
	CountRequest int64               `json:"count_request"`
}

// StatusResponse is the message returned by StatusHandler.
type StatusResponse struct {
	SvcAlive   bool `json:"svc-alive"`
	StoreAlive bool `json:"store-alive"`
}
